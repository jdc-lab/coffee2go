package uc

// interactor implements Handler -> can handle all usecases (implementations are in extra files)
// It also contains all interfaces needed by the usecase-implementations which are implemented by impl package
type interactor struct {
	connection Connection // Handles connections to a chat server
	chat       Chat       // Handles everything needed by chats
	push       Push       // Handles push message connections
	conf       Conf       // Handles configuration
}

type Connection interface {
	/*UserLogin(host, username, password string) (user *domain.ChatConnection, token string, err error)
	GenUserToken(username string) (token string, err error)
	GetUser(token string) (userName *domain.ChatConnection, err error)*/
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
