// pushclient/client.go
package pushclient

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket" // 你需要 `go get github.com/gorilla/websocket`
)

const (
	defaultHeartbeatInterval = 30 * time.Second // 默认心跳间隔
	writeWait                = 10 * time.Second //写入超时
	pongWait                 = 60 * time.Second // Pong 消息等待超时 (应大于心跳间隔)
	maxMessageSize           = 512              // 最大消息大小
)

// ThirdPartyAPIResponse 是第三方 API 返回 WebSocket URL 的预期结构
type ThirdPartyAPIResponse struct {
	URL string `json:"url"`
}

// AuthMessage 是通过 WebSocket 发送的认证消息结构
type AuthMessage struct {
	AppID    string `json:"app_id"`
	ClientID string `json:"client_id"`
}

// Client 管理与第三方推送服务的连接
type Client struct {
	AppID            string
	ClientID         string
	ThirdPartyAPIURL string // 例如："http://39.107.58.199:8008/search"
	WebSocketURL     string

	conn        *websocket.Conn
	mu          sync.Mutex // 互斥锁，保护并发访问 conn 和 isConnected
	isConnected bool

	// MessageChan 用于接收来自 WebSocket 的消息 (不包括心跳)
	MessageChan chan string
	// ErrorChan 用于接收 WebSocket 操作期间发生的错误
	ErrorChan chan error
	// StatusChan 用于连接状态更新 ("connecting", "connected", "disconnected", "error: <reason>")
	StatusChan chan string

	ctx    context.Context    // 用于控制 goroutine 的生命周期
	cancel context.CancelFunc // 用于取消 context

	logger *log.Logger // 日志记录器
}

// NewClient 创建一个新的推送客户端实例
func NewClient(appID, clientID, apiURL string, logger *log.Logger) *Client {
	if logger == nil {
		logger = log.New(log.Writer(), "[PushClient] ", log.LstdFlags)
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &Client{
		AppID:            appID,
		ClientID:         clientID,
		ThirdPartyAPIURL: apiURL,
		MessageChan:      make(chan string, 10), // 缓冲通道
		ErrorChan:        make(chan error, 1),   // 缓冲通道
		StatusChan:       make(chan string, 5),  // 缓冲通道
		ctx:              ctx,
		cancel:           cancel,
		logger:           logger,
	}
}

// updateStatus 更新状态并通过 StatusChan 发送
func (c *Client) updateStatus(status string) {
	c.logger.Println(status)
	select {
	case c.StatusChan <- status:
	default: // 如果 StatusChan 满了或没有监听者，则不阻塞
		c.logger.Println("StatusChan 已满或无监听者, 状态被丢弃:", status)
	}
}

// getWebSocketURL 从第三方 API 获取 WebSocket URL
func (c *Client) getWebSocketURL() (string, error) {
	c.updateStatus(fmt.Sprintf("正在从 %s 获取 WebSocket URL", c.ThirdPartyAPIURL))

	payload := map[string]string{
		"app_id":    c.AppID,
		"client_id": c.ClientID,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("序列化 API 请求失败: %w", err)
	}

	req, err := http.NewRequestWithContext(c.ctx, "POST", c.ThirdPartyAPIURL, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return "", fmt.Errorf("创建 API 请求失败: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{Timeout: 10 * time.Second}
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("API 请求失败: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var bodyBytes []byte
		if resp.Body != nil {
			// 在 Go 1.16+ 中，可以使用 io.ReadAll
			// 这里我们用一个简单的循环读取
			buf := new(bytes.Buffer)
			buf.ReadFrom(resp.Body)
			bodyBytes = buf.Bytes()
		}
		return "", fmt.Errorf("API 请求返回非成功状态: %d %s. 响应体: %s", resp.StatusCode, resp.Status, string(bodyBytes))
	}

	var apiResp ThirdPartyAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return "", fmt.Errorf("解码 API 响应失败: %w", err)
	}

	if apiResp.URL == "" {
		return "", errors.New("API 响应未包含 WebSocket URL")
	}
	c.logger.Printf("成功获取 WebSocket URL: %s", apiResp.URL)
	return apiResp.URL, nil
}

// Connect 建立 WebSocket 连接并开始监听
func (c *Client) Connect() error {
	c.mu.Lock()
	if c.isConnected {
		c.mu.Unlock()
		return errors.New("已经连接")
	}
	c.mu.Unlock()

	// 为新的连接尝试重置 context
	c.ctx, c.cancel = context.WithCancel(context.Background())

	wsURL, err := c.getWebSocketURL()
	if err != nil {
		errMsg := fmt.Sprintf("获取 WebSocket URL 失败: %v", err)
		c.updateStatus(errMsg)
		c.ErrorChan <- err
		return err
	}
	c.WebSocketURL = wsURL
	c.updateStatus(fmt.Sprintf("正在连接到 WebSocket: %s", c.WebSocketURL))

	conn, resp, err := websocket.DefaultDialer.DialContext(c.ctx, c.WebSocketURL, nil)
	if err != nil {
		errMsg := fmt.Sprintf("WebSocket 连接错误: %v", err)
		if resp != nil {
			errMsg = fmt.Sprintf("WebSocket 连接错误: %v (状态: %s)", err, resp.Status)
		}
		c.updateStatus(errMsg)
		c.ErrorChan <- err
		return err
	}
	c.conn = conn

	c.mu.Lock()
	c.isConnected = true
	c.mu.Unlock()
	c.updateStatus("WebSocket 连接已建立")

	// 发送认证消息
	authMsg := AuthMessage{AppID: c.AppID, ClientID: c.ClientID}
	authPayload, _ := json.Marshal(authMsg)
	if err := c.sendMessage(websocket.TextMessage, authPayload); err != nil {
		errMsg := fmt.Sprintf("发送认证消息失败: %v", err)
		c.updateStatus(errMsg)
		c.ErrorChan <- err
		c.Disconnect() // 如果认证失败则断开连接
		return err
	}
	c.logger.Println("已发送认证消息到 WebSocket")

	// 启动读取和心跳 goroutine
	go c.readPump()
	go c.heartbeatPump()

	return nil
}

// sendMessage 是一个带写入超时的发送消息的辅助函数
func (c *Client) sendMessage(messageType int, data []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.conn == nil || !c.isConnected {
		return errors.New("未连接")
	}
	c.conn.SetWriteDeadline(time.Now().Add(writeWait))
	return c.conn.WriteMessage(messageType, data)
}

// SendCustomMessage 允许通过 WebSocket 发送任意字符串消息
func (c *Client) SendCustomMessage(message string) error {
	if !c.IsConnected() {
		return errors.New("未连接，无法发送消息")
	}
	c.logger.Printf("发送自定义消息: %s", message)
	return c.sendMessage(websocket.TextMessage, []byte(message))
}

func (c *Client) readPump() {
	defer func() {
		c.logger.Println("读取 goroutine (readPump) 已停止。")
		c.Disconnect() // 确保 readPump 退出时执行清理
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait)) // 设置初始读取超时
	c.conn.SetPongHandler(func(string) error {       // 处理服务器的 Pong 消息
		c.logger.Println("收到服务器的 PONG")
		c.conn.SetReadDeadline(time.Now().Add(pongWait)) // 重置读取超时
		return nil
	})

	for {
		select {
		case <-c.ctx.Done(): // 被 Disconnect() 或其他方式取消
			c.logger.Println("Context 已取消, 停止 readPump。")
			return
		default:
			messageType, message, err := c.conn.ReadMessage()
			if err != nil {
				// IsUnexpectedCloseError 检查是否是 "意外" 的关闭错误
				// CloseGoingAway 通常是浏览器tab关闭, CloseAbnormalClosure 是网络问题等
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure, websocket.CloseNormalClosure) {
					errMsg := fmt.Sprintf("读取错误: %v", err)
					c.updateStatus(errMsg)
					c.ErrorChan <- err
				} else if err == websocket.ErrCloseSent { // 自己发送了关闭帧
					c.logger.Println("readPump: WebSocket 由本地关闭。")
				} else if e, ok := err.(*websocket.CloseError); ok && e.Code == websocket.CloseNormalClosure {
					c.logger.Println("readPump: WebSocket 由远端正常关闭。")
				} else {
					c.logger.Printf("读取错误 (未处理类型): %v", err) // 例如 net.OpError on read, context canceled
				}
				return // 任何读取错误都退出 goroutine
			}

			// 收到任何消息都重置读取超时
			c.conn.SetReadDeadline(time.Now().Add(pongWait))

			msgStr := string(message)
			c.logger.Printf("收到消息 (类型 %d): %s", messageType, msgStr)

			switch msgStr {
			case "heartbeat": // 服务器要求心跳确认
				c.logger.Println("收到服务器的 'heartbeat', 发送 'heartbeat_ack'")
				if err := c.sendMessage(websocket.TextMessage, []byte("heartbeat_ack")); err != nil {
					errMsg := fmt.Sprintf("发送 heartbeat_ack 失败: %v", err)
					c.updateStatus(errMsg)
					c.ErrorChan <- err
				}
			case "heartbeat_ack": // 服务器确认了我们的 heartbeat_check
				c.logger.Println("收到服务器的 'heartbeat_ack' (对我们 check 的响应)")
			default:
				// 将其他消息转发到 MessageChan
				select {
				case c.MessageChan <- msgStr:
				case <-c.ctx.Done(): // 检查 context 是否已取消
					return
				default: // 如果 MessageChan 满了，防止阻塞
					c.logger.Println("MessageChan 已满或无监听者, 消息被丢弃:", msgStr)
				}
			}
		}
	}
}

func (c *Client) heartbeatPump() {
	ticker := time.NewTicker(defaultHeartbeatInterval)
	defer func() {
		ticker.Stop()
		c.logger.Println("心跳 goroutine (heartbeatPump) 已停止。")
	}()

	for {
		select {
		case <-ticker.C:
			c.mu.Lock() // 保护 c.conn
			if c.conn == nil || !c.isConnected {
				c.mu.Unlock()
				c.logger.Println("心跳: 未连接, 跳过发送心跳。")
				return // 如果未连接，则停止心跳
			}
			c.mu.Unlock()

			c.logger.Println("发送 'heartbeat_check' (客户端 PING)")
			// Gorilla WebSocket 库的 WriteControl 会处理 PING/PONG 控制帧。
			// 但你的 JS 发送的是 'heartbeat_check' 文本消息，我们保持一致。
			if err := c.sendMessage(websocket.TextMessage, []byte("heartbeat_check")); err != nil {
				errMsg := fmt.Sprintf("发送 heartbeat_check 失败: %v", err)
				c.updateStatus(errMsg)
				c.ErrorChan <- err
				// 如果发送心跳失败，也考虑是否需要断开连接
				return // 发送错误时停止心跳
			}
		case <-c.ctx.Done(): // 被 Disconnect() 取消
			return
		}
	}
}

// Disconnect 关闭 WebSocket 连接
func (c *Client) Disconnect() {
	c.mu.Lock()
	if !c.isConnected {
		c.mu.Unlock()
		c.logger.Println("已经断开或未连接。")
		return
	}

	c.logger.Println("正在断开 WebSocket 连接...")
	c.isConnected = false // 立即设置，防止新的操作

	if c.cancel != nil { // 通知所有 goroutine 停止
		c.cancel()
	}

	if c.conn != nil {
		// 优雅关闭：发送一个关闭帧并等待服务器响应
		err := c.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		if err != nil && err != websocket.ErrCloseSent && !errors.Is(err, errors.New("use of closed network connection")) { // 检查更具体的网络关闭错误
			c.logger.Printf("写入关闭消息错误: %v", err)
		}
		// time.Sleep(100 * time.Millisecond) // 可选：给关闭消息一点发送时间
		c.conn.Close() // 然后强制关闭
		c.conn = nil
	}
	c.mu.Unlock()

	c.updateStatus("WebSocket 已断开")
}

// IsConnected 返回当前连接状态
func (c *Client) IsConnected() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.isConnected && c.conn != nil
}

// Ctx 返回客户端的 context，允许外部监听其取消事件
func (c *Client) Ctx() context.Context {
	return c.ctx
}
