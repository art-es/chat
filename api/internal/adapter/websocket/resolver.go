package websocket

import (
	"encoding/json"
	"fmt"

	"github.com/art-es/chat/api/internal/core/entity"
	mt "github.com/art-es/chat/api/internal/core/entity/messagetype"
	"github.com/art-es/chat/api/internal/core/port"
)

type (
	ResolveFunc func(*entity.Session, []byte) error
	Resolver    map[mt.InputType]ResolveFunc
)

func NewResolver(users port.UserRepository) Resolver {
	return map[mt.InputType]ResolveFunc{
		mt.SendMessage: sendMessageResolver(users),
	}
}

func sendMessageResolver(users port.UserRepository) ResolveFunc {
	type input struct {
		Recipient string `json:"recipient"`
		Content   string `json:"content"`
	}
	type output struct {
		mt.Type
		Sender  string `json:"sender"`
		Content string `json:"content"`
	}

	return func(sess *entity.Session, b []byte) error {
		var in input
		if err := json.Unmarshal(b, &in); err != nil {
			return err
		}

		recipient := users.Get(in.Recipient)
		if recipient == nil {
			return fmt.Errorf("Recipient %s doesn't exists", in.Recipient)
		}

		broadcast(
			append(sess.User.Sessions, recipient.Sessions...),
			output{mt.SentMessage, sess.User.Name, in.Content},
		)
		return nil
	}
}
