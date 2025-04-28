package handler

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gorilla/websocket"
	"message-push/app/pusher/common/model/entity"
	"message-push/app/pusher/service/internal/service"
	"net/http"
)

// WebSocketHandler 定义 WebSocket 处理器
type WebSocketHandler struct {
	log         *log.Helper
	upgrader    websocket.Upgrader // WebSocket 握手升级器
	sendService *service.SendService
}

func NewWebSocketHandler(logger log.Logger, service *service.SendService) *WebSocketHandler {
	return &WebSocketHandler{
		log: log.NewHelper(logger),
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			// 安全校验（根据需求开启）
			CheckOrigin: func(r *http.Request) bool {
				return true // 允许所有来源，生产环境需严格校验
			},
		},
		sendService: service,
	}
}

// ServeHTTP 处理 WebSocket 连接
func (h *WebSocketHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// 先完成WebSocket握手
	conn, err := h.upgrader.Upgrade(w, req, nil)
	if err != nil {
		h.log.Errorf("WebSocket upgrade failed: %v", err)
		return
	}

	// 通过WebSocket连接读取认证信息
	var cliReq entity.ClientRequest
	if err := conn.ReadJSON(&cliReq); err != nil {
		h.log.Errorf("Read auth failed: %v", err)
		return
	}

	// 注册客户端
	h.sendService.RegisterClient(&cliReq, conn)
}
