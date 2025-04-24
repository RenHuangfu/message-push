package handler

import (
	"net/http"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/gorilla/websocket"
)

// WebSocketHandler 定义 WebSocket 处理器
type WebSocketHandler struct {
	log      *log.Helper
	upgrader websocket.Upgrader // WebSocket 握手升级器
}

func NewWebSocketHandler(logger log.Logger) http.Handler {
	wsHandler := &WebSocketHandler{
		log: log.NewHelper(logger),
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			// 安全校验（根据需求开启）
			CheckOrigin: func(r *http.Request) bool {
				return true // 允许所有来源，生产环境需严格校验
			},
		},
	}
	http.Handle("/ws", wsHandler)
	return wsHandler
}

// ServeHTTP 处理 WebSocket 连接
func (h *WebSocketHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// 升级 HTTP 连接到 WebSocket
	conn, err := h.upgrader.Upgrade(w, req, nil)
	if err != nil {
		h.log.Errorf("upgrade failed: %v", err)
		return
	}
	defer conn.Close()

	h.log.Info("new websocket connection")

	// 启动读写协程
	go h.readLoop(conn)
	go h.writeLoop(conn)
}

// 读取消息循环
func (h *WebSocketHandler) readLoop(conn *websocket.Conn) {
	defer conn.Close()
	for {
		// 读取消息类型和数据
		msgType, data, err := conn.ReadMessage()
		if err != nil {
			h.log.Errorf("read message failed: %v", err)
			return
		}
		h.log.Infof("received message: type=%d, data=%s", msgType, data)

		// 处理消息（示例：回显）
		if err := conn.WriteMessage(msgType, data); err != nil {
			h.log.Errorf("write message failed: %v", err)
			return
		}
	}
}

// 写入消息循环（示例：定时发送心跳）
func (h *WebSocketHandler) writeLoop(conn *websocket.Conn) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			// 发送心跳消息
			if err := conn.WriteMessage(websocket.TextMessage, []byte("heartbeat")); err != nil {
				h.log.Errorf("send heartbeat failed: %v", err)
				return
			}
		}
	}
}
