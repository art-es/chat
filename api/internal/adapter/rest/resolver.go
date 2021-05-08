package rest

import (
	"net/http"
	"time"

	"github.com/art-es/chat/api/internal/core/port"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Resolver struct {
	WSUpgrader       *websocket.Upgrader
	Sessions         port.SessionRepository
	SessionService   port.SessionService
	WebsocketService port.WebsocketService
}

func NewResolver(s port.SessionService, ws port.WebsocketService) *Resolver {
	return &Resolver{
		WSUpgrader: &websocket.Upgrader{
			HandshakeTimeout: 10 * time.Second,
			Error:            func(w http.ResponseWriter, r *http.Request, status int, reason error) {},
			CheckOrigin:      func(r *http.Request) bool { return true },
		},
		SessionService:   s,
		WebsocketService: ws,
	}
}

func (r Resolver) NewSession(c *gin.Context) error {
	username := c.Query("username")
	if username == "" {
		return newErrorString("username is required query parameter", http.StatusBadRequest)
	}

	conn, err := r.WSUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return err
	}

	sess := r.SessionService.Register(username, conn)

	return r.WebsocketService.ListenConn(sess)
}

func (r Resolver) GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"usernames": r.Sessions.GetUsernames()})
}
