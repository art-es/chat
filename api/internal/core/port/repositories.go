package port

import "github.com/art-es/chat/api/internal/core/entity"

type SessionRepository interface {
	GetUsernames() []string
	Add(*entity.Session)
	Delete(*entity.Session)
}

type UserRepository interface {
	Get(string) *entity.User
	Add(*entity.User)
}
