package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func NewHandler(r *Resolver) *gin.Engine {
	h := gin.New()
	h.Use(gin.Recovery())
	h.GET("/ws", warnmdw(r.NewSession))
	h.GET("/users", r.GetUsers)
	return h
}

func warnmdw(resolve func(c *gin.Context) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := resolve(c); err != nil {
			log.WithError(err).Warn("GET /ws")
		}
	}
}

func errormdw(resolve func(c *gin.Context) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := resolve(c); err != nil {
			status := http.StatusInternalServerError
			if httperr, ok := err.(*httpError); ok {
				status = httperr.status
			}

			c.JSON(status, gin.H{
				"ok":    false,
				"error": err.Error(),
			})
		}
	}
}
