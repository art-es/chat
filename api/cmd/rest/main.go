package main

import (
	"net/http"
	"time"

	"github.com/art-es/chat/api/internal/ioc"

	log "github.com/sirupsen/logrus"
)

func main() {
	c, err := ioc.NewContainer()
	if err != nil {
		log.WithError(err).Fatal("main: ioc.NewContainer")
	}

	srv := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	if err := c.Invoke(func(h http.Handler) { srv.Handler = h }); err != nil {
		log.WithError(err).Fatal("main: container.Invoke http.Handler")
	}

	if err := srv.ListenAndServe(); err != nil {
		log.WithError(err).WithField("srv", srv).Fatal("main: httpServer.ListenAndServe")
	}
}
