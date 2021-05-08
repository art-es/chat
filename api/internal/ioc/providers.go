package ioc

import (
	"net/http"

	"github.com/art-es/chat/api/internal/adapter/memrepo/sessionrepo"
	"github.com/art-es/chat/api/internal/adapter/memrepo/userrepo"
	"github.com/art-es/chat/api/internal/adapter/rest"
	"github.com/art-es/chat/api/internal/adapter/websocket"
	"github.com/art-es/chat/api/internal/core/port"
	"github.com/art-es/chat/api/internal/core/service/sessionsvc"
	"github.com/art-es/chat/api/internal/core/service/websocketsvc"
)

var (
	repositoriesProvider = Provider{
		func() port.SessionRepository { return sessionrepo.New() },
		func() port.UserRepository { return userrepo.New() },
	}

	sessionProvider = Provider{
		func(s port.SessionRepository, u port.UserRepository) port.SessionService {
			return sessionsvc.New(s, u)
		},
	}

	websocketProvider = Provider{
		websocket.NewResolver,

		func(r websocket.Resolver) websocketsvc.Handler {
			return websocket.NewHandler(r)
		},

		func(h websocketsvc.Handler, sessions port.SessionRepository) port.WebsocketService {
			return websocketsvc.New(h, sessions)
		},
	}

	restProvider = Provider{
		rest.NewResolver,

		func(r *rest.Resolver) http.Handler {
			return rest.NewHandler(r)
		},
	}
)
