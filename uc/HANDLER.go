package uc

import (
	"github.com/jdc-lab/coffee2go/domain"
	"log"
)

// Handler containing all usecases (separated as several interfaces)
type Handler interface {
	ConnectLogic
	ChatLogic
}

type ConnectLogic interface {
	ConnectServer(host, username, password string) (sessionID string, err error)
	ConnectPreset() (host, username, password string)
}

type ChatLogic interface {
	ChatSend(message domain.Message) (err error)
	ChatRegisterForPush() (err error)
	ChatPushNewMessage() (err error)
	ChatGetContacts() (err error)
	ChatSwitch() (err error)
}

type HandlerConstructor struct {
	Connection Connection
	Chat       Chat
	Push       Push
}

func (c HandlerConstructor) New() Handler {
	if c.Connection == nil {
		log.Fatal("missing Connection")
	}
	if c.Chat == nil {
		log.Fatal("missing Chat")
	}
	if c.Push == nil {
		log.Fatal("missing Push")
	}

	return interactor{
		connection: c.Connection,
		chat:       c.Chat,
		push:       c.Push,
	}
}
