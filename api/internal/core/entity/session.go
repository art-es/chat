package entity

import "github.com/gorilla/websocket"

type Session struct {
	Conn *websocket.Conn
	User *User
}

func NewSession(conn *websocket.Conn, user *User) *Session {
	return &Session{Conn: conn, User: user}
}
