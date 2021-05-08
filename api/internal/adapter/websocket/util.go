package websocket

import (
	"encoding/json"

	"github.com/art-es/chat/api/internal/core/entity"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

func write(s *entity.Session, obj interface{}) {
	if err := s.Conn.WriteJSON(obj); err != nil {
		log.WithError(err).Error("websocket: write", log.Fields{"user": s.User.Name})
	}
}

func broadcast(ss []*entity.Session, obj interface{}) {
	b, err := json.Marshal(obj)
	if err != nil {
		log.WithError(err).Error("websocket: broadcast: json.Marshal")
		return
	}

	for _, s := range ss {
		if err = s.Conn.WriteMessage(websocket.TextMessage, b); err != nil {
			log.WithError(err).Error("websocket: broadcast: s.Conn.WriteMessage", log.Fields{"user": s.User.Name})
		}
	}
}
