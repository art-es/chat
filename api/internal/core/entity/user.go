package entity

type User struct {
	Name     string
	Sessions []*Session
}

func (u *User) AddSession(s *Session) {
	u.Sessions = append(u.Sessions, s)
}

func NewUser(name string) *User {
	return &User{Name: name}
}
