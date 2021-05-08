package websocketsvc

import (
	"github.com/art-es/chat/api/internal/core/entity"
	"github.com/art-es/chat/api/internal/core/port"

	"github.com/gorilla/websocket"
)

type Handler interface {
	Handle(*entity.Session, []byte)
}

type Service struct {
	Handler
	Sessions port.SessionRepository
}

func New(h Handler, sessions port.SessionRepository) *Service {
	return &Service{h, sessions}
}

func (s Service) ListenConn(sess *entity.Session) error {
	defer s.Sessions.Delete(sess)

	for {
		msgtype, rawmsg, err := sess.Conn.ReadMessage()
		if err != nil {
			return err
		}

		switch msgtype {
		case websocket.CloseMessage:
			return nil

		case websocket.TextMessage, websocket.BinaryMessage:
			s.Handle(sess, rawmsg)
		}
	}
}
