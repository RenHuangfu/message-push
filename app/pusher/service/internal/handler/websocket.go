package handler

import (
	"encoding/json"
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
	var cliReq entity.ClientRequest
	err := json.NewDecoder(req.Body).Decode(&cliReq)
	if err != nil {
		h.log.Errorf("marshal failed: %v", err)
		return
	}

	conn, err := h.upgrader.Upgrade(w, req, nil)
	if err != nil {
		h.log.Errorf("upgrade failed: %v", err)
		return
	}

	h.log.Info("new websocket connection")

	h.sendService.RegisterClient(&cliReq, conn)
}
