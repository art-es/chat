package sessionrepo

import (
	"github.com/art-es/chat/api/internal/core/entity"
	"github.com/gorilla/websocket"
)

type Repository struct {
	items     map[*websocket.Conn]*entity.Session
	usernames []string
}

func New() *Repository {
	return &Repository{items: make(map[*websocket.Conn]*entity.Session)}
}

func (r *Repository) GetUsernames() []string {
	return r.usernames
}

func (r *Repository) Add(s *entity.Session) {
	r.items[s.Conn] = s
	r.addUsername(s.User.Name)
}

func (r *Repository) addUsername(username string) {
	for _, un := range r.usernames {
		if un == username {
			return
		}
	}
	r.usernames = append(r.usernames, username)
}

func (r *Repository) Delete(s *entity.Session) {
	r.deleteUsername(s.User.Name)
	delete(r.items, s.Conn)
}

func (r *Repository) deleteUsername(username string) {
	for i, un := range r.usernames {
		if un == username {
			r.usernames = append(r.usernames[:i], r.usernames[i+1:]...)
			return
		}
	}
}
