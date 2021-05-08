package port

import (
	"github.com/art-es/chat/api/internal/core/entity"

	"github.com/gorilla/websocket"
)

type SessionService interface {
	Register(username string, conn *websocket.Conn) *entity.Session
}

type WebsocketService interface {
	ListenConn(*entity.Session) error
}
