package userrepo

import "github.com/art-es/chat/api/internal/core/entity"

type Repository struct {
	items map[string]*entity.User
}

func New() *Repository {
	return &Repository{items: make(map[string]*entity.User)}
}

func (r *Repository) Get(name string) *entity.User {
	return r.items[name]
}

func (r *Repository) Add(u *entity.User) {
	r.items[u.Name] = u
}

func (r *Repository) GetOrCreate(name string) *entity.User {
	user, ok := r.items[name]
	if !ok {
		user = entity.NewUser(name)
		r.items[name] = user
	}
	return user
}
