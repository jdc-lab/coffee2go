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
	ChatRegisterForPush(sessionID string) (pushToken string, err error)
	ChatPushNewMessage() (err error)
	ChatGetContacts() (err error)
	ChatSwitch() (err error)
}

type HandlerConstructor struct {
	Connection Connection
	Push       Push
	Conf       Conf
	Session    Session
}

func (c HandlerConstructor) New() Handler {
	if c.Connection == nil {
		log.Fatal("missing Connection")
	}
	if c.Push == nil {
		log.Fatal("missing Push")
	}
	if c.Conf == nil {
		log.Fatal("missing Conf")
	}
	if c.Session == nil {
		log.Fatal("missing Session")
	}

	return interactor{
		connection: c.Connection,
		push:       c.Push,
		conf:       c.Conf,
		session:    c.Session,
	}
}
