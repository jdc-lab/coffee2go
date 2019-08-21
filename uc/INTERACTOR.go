package uc

import "github.com/jdc-lab/coffee2go/domain"

// interactor implements Handler -> can handle all usecases (implementations are in extra files)
// It also contains all interfaces needed by the usecase-implementations which are implemented by impl package
type interactor struct {
	connection Connection // Handles connections to a chat server
	chat       Chat       // Handles everything needed by chats
	push       Push       // Handles push message connections
	conf       Conf       // Handles configuration
	session    Session    // Handles sessions
}

type Connection interface {
	Connect(host, username, password string) (serverConnection *Chat, err error)
}

type Chat interface {
	/*Send(message domain.Message, recipient domain.ChatConnection) (err error)
	GetContacts() (err error)
	SwitchChat() (err error)*/
}

type Push interface {
	Register() (err error)
	Send() (err error)
}

type Conf interface {
	GetConnectionPreset() (host, username, password string)
}

type Session interface {
	Add(session *Chat) (sessionID string, err error)
	Get(sessionId string) (session *domain.Session, err error)
}
