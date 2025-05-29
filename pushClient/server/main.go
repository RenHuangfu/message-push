// server/main.go
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	pushclient "push-client/pushClient"
	"sync"
)

var (
	templates *template.Template
	// 为简单起见，使用全局客户端实例。对于多用户，你需要一个 map 或会话管理。
	currentClient *pushclient.Client
	clientMutex   sync.Mutex // 保护 currentClient 的并发访问
	logger        *log.Logger
)

// PageData 用于 HTML 模板的数据
type PageData struct {
	AppID    string
	ClientID string
	Status   string
	Messages []string // 消息将通过 SSE 流式传输
}

const thirdPartyAPI = "http://8.147.117.133:8008/search" // 第三方 API 地址

func init() {
	logger = log.New(os.Stdout, "[Server] ", log.LstdFlags)

	// 解析模板
	// 假设模板在相对于 main.go 的 "templates" 子目录中
	wd, _ := os.Getwd()
	// 如果从 `your-project/` 目录运行: templatePath := "server/templates/*.html"
	templatePath := filepath.Join(wd, "server", "templates", "*.html")
	logger.Printf("正在加载模板: %s", templatePath)
	templates = template.Must(template.ParseGlob(templatePath))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	clientMutex.Lock()
	defer clientMutex.Unlock()

	status := "未连接"
	var appID, clientID string
	if currentClient != nil {
		appID = currentClient.AppID
		clientID = currentClient.ClientID
		if currentClient.IsConnected() {
			status = "已连接"
		} else {
			// 检查 context 是否已取消，以区分主动断开和连接失败
			if currentClient.Ctx().Err() != nil {
				status = "已断开"
			} else if currentClient.WebSocketURL != "" { // 如果尝试过连接
				status = "连接失败或已断开"
			}
		}
	}

	data := PageData{
		AppID:    appID,
		ClientID: clientID,
		Status:   status,
		Messages: []string{}, // 初始消息为空，将通过 SSE 更新
	}
	err := templates.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		logger.Printf("执行模板错误: %v", err)
		http.Error(w, "内部服务器错误", http.StatusInternalServerError)
	}
}

type ConnectRequest struct {
	AppID    string `json:"appId"`
	ClientID string `json:"clientId"`
}

func connectHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "仅允许 POST 方法", http.StatusMethodNotAllowed)
		return
	}

	var req ConnectRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("无效的请求体: %v", err), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if req.AppID == "" || req.ClientID == "" {
		http.Error(w, "AppID 和 ClientID 是必需的", http.StatusBadRequest)
		return
	}

	clientMutex.Lock()
	// 如果存在现有客户端，先断开它
	if currentClient != nil {
		logger.Println("在新的连接之前断开旧的客户端。")
		currentClient.Disconnect() // Disconnect 是同步的，会等待goroutine结束
		// 等待一小段时间确保资源完全释放，或在Disconnect中确保这一点
		// time.Sleep(200 * time.Millisecond)
	}

	logger.Printf("为 AppID: %s, ClientID: %s 创建新的推送客户端", req.AppID, req.ClientID)
	currentClient = pushclient.NewClient(req.AppID, req.ClientID, thirdPartyAPI, logger)
	// 捕获要传递给 goroutine 的 client 实例
	clientToConnect := currentClient
	clientMutex.Unlock() // 在调用 Connect 之前解锁，因为它可能是阻塞的

	go func(client *pushclient.Client) {
		// 这个 goroutine 负责连接，并将结果通过 client 的通道传递
		err := client.Connect() // Connect 方法会阻塞直到连接成功或失败
		if err != nil {
			logger.Printf("客户端连接失败: %v", err)
			// 错误已经通过 client.ErrorChan 发送，并由 SSE 处理
		} else {
			logger.Println("客户端通过 goroutine 成功连接。")
		}
	}(clientToConnect)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "连接过程已启动。"})
}

func disconnectHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "仅允许 POST 方法", http.StatusMethodNotAllowed)
		return
	}

	clientMutex.Lock()
	defer clientMutex.Unlock()

	if currentClient == nil || !currentClient.IsConnected() {
		logger.Println("没有活动的客户端可断开或已经断开。")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "没有活动的客户端或已经断开。"})
		return
	}

	logger.Println("正在断开客户端...")
	currentClient.Disconnect() // Disconnect 是同步的

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "断开过程已启动。"})
}

// sseHandler 将事件 (状态、消息、错误) 流式传输到客户端浏览器
func sseHandler(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "不支持流式传输!", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*") // 生产环境中应配置具体的源

	// 为此 SSE 连接创建本地通道，用于从 client 的通道接收数据
	// 这些通道是临时的，仅用于当前的 SSE 连接
	clientMsgChan := make(chan string)
	clientErrChan := make(chan error)
	clientStatusChan := make(chan string)

	clientMutex.Lock()
	activeClient := currentClient // 捕获当前客户端的引用
	clientMutex.Unlock()

	var wg sync.WaitGroup // 用于等待 goroutine 结束

	if activeClient != nil {
		// Goroutine: 监听 activeClient 的 MessageChan
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case msg, ok := <-activeClient.MessageChan:
					if !ok { // SDK 的 MessageChan 关闭了
						logger.Println("SSE: SDK MessageChan 已关闭。")
						return
					}
					select { // 非阻塞发送到此 SSE 会话的本地通道
					case clientMsgChan <- msg:
					case <-r.Context().Done(): // SSE 客户端断开连接
						return
					case <-activeClient.Ctx().Done(): // Push client 本身断开了
						return
					}
				case <-r.Context().Done():
					return
				case <-activeClient.Ctx().Done():
					return
				}
			}
		}()

		// Goroutine: 监听 activeClient 的 ErrorChan
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case err, ok := <-activeClient.ErrorChan:
					if !ok {
						logger.Println("SSE: SDK ErrorChan 已关闭。")
						return
					}
					select {
					case clientErrChan <- err:
					case <-r.Context().Done():
						return
					case <-activeClient.Ctx().Done():
						return
					}
				case <-r.Context().Done():
					return
				case <-activeClient.Ctx().Done():
					return
				}
			}
		}()

		// Goroutine: 监听 activeClient 的 StatusChan
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case status, ok := <-activeClient.StatusChan:
					if !ok {
						logger.Println("SSE: SDK StatusChan 已关闭。")
						return
					}
					select {
					case clientStatusChan <- status:
					case <-r.Context().Done():
						return
					case <-activeClient.Ctx().Done():
						return
					}
				case <-r.Context().Done():
					return
				case <-activeClient.Ctx().Done():
					return
				}
			}
		}()
	} else {
		// 如果没有客户端实例，发送一个初始的 "未连接" 状态
		fmt.Fprintf(w, "event: status\ndata: %s\n\n", "\"未连接 (无客户端实例)\"")
		flusher.Flush()
	}

	// SSE 主循环：将从 client 通道收到的数据发送给浏览器
	for {
		select {
		case msg := <-clientMsgChan:
			escapedMsg, _ := json.Marshal(msg) // 确保消息是有效的 JSON 字符串
			fmt.Fprintf(w, "event: message\ndata: %s\n\n", escapedMsg)
			flusher.Flush()
		case err := <-clientErrChan:
			escapedErr, _ := json.Marshal(err.Error())
			fmt.Fprintf(w, "event: error\ndata: %s\n\n", escapedErr)
			flusher.Flush()
		case status := <-clientStatusChan:
			escapedStatus, _ := json.Marshal(status)
			fmt.Fprintf(w, "event: status\ndata: %s\n\n", escapedStatus)
			flusher.Flush()
		case <-r.Context().Done(): // 浏览器关闭了 SSE 连接
			logger.Println("SSE: 浏览器客户端断开连接, 关闭 SSE 流。")
			// activeClient 的 goroutine 应该也会因为 activeClient.Ctx().Done() 或 r.Context().Done() 而退出
			wg.Wait() // 等待所有监听 activeClient 通道的 goroutine 结束
			return
		}
	}
}

func main() {
	port := flag.String("port", "7788", "服务器监听端口")
	wd, _ := os.Getwd()
	defaultStaticDir := filepath.Join(wd, "server", "static")
	dir := flag.String("dir", defaultStaticDir, "静态文件目录 (相对于项目根目录的 server/static)")
	flag.Parse()

	// 静态文件服务
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(*dir))))
	// 如果 dir 是绝对路径或相对于可执行文件的路径，则不需要 StripPrefix 如果路径匹配
	// 但为了清晰，使用 StripPrefix
	staticFS := http.FileServer(http.Dir(*dir))
	http.Handle("/static/", http.StripPrefix("/static/", staticFS))

	// 应用处理器
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/api/connect", connectHandler)
	http.HandleFunc("/api/disconnect", disconnectHandler)
	http.HandleFunc("/api/events", sseHandler) // Server-Sent Events 端点

	logger.Printf("服务器启动，监听端口 %s，静态文件目录 %s", *port, *dir)
	if err := http.ListenAndServe(":"+*port, nil); err != nil {
		logger.Fatalf("启动服务器失败: %v", err)
	}
}
