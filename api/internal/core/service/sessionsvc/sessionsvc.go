package sessionsvc

import (
	"github.com/art-es/chat/api/internal/core/entity"
	"github.com/art-es/chat/api/internal/core/port"

	"github.com/gorilla/websocket"
)

type Service struct {
	Sessions port.SessionRepository
	Users    port.UserRepository
}

func New(sessions port.SessionRepository, users port.UserRepository) *Service {
	return &Service{Sessions: sessions, Users: users}
}

func (s Service) Register(username string, conn *websocket.Conn) *entity.Session {
	user := s.Users.Get(username)
	if user == nil {
		user = entity.NewUser(username)
		s.Users.Add(user)
	}

	sess := entity.NewSession(conn, user)
	s.Sessions.Add(sess)

	user.AddSession(sess)
	return sess
}
