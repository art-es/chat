package websocket

import (
	"github.com/art-es/chat/api/internal/core/entity"
	"github.com/art-es/chat/api/internal/core/entity/messagetype"

	log "github.com/sirupsen/logrus"
)

type Handler struct {
	Resolver
}

func (h *Handler) Handle(sess *entity.Session, msg []byte) {
	if resolve, ok := h.Resolver[messagetype.ParseInput(msg)]; ok {
		if err := resolve(sess, msg); err != nil {
			log.WithError(err).Error("websocket: resolve")
		}
	} else {
		log.WithField("rawmsg", msg).Error("websocket: unknown message type")
	}
}

func NewHandler(r Resolver) *Handler {
	return &Handler{r}
}
