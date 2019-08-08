package chat

import (
	"fmt"
	"strings"
	"time"
)

type UserID string

type User struct {
	Name string
	ID   UserID
}

type History struct {
	From      UserID
	To        UserID
	Subject   string
	Message   string
	Timestamp time.Time
}

type Client interface {
	// Connect to a chat server by using the provided host, username and password.
	// Should return AlreadyLoggedIn error if called a second time.
	Login(host, username, password string) error
	Send(to UserID, message string) error
	GetContacts() []User
	GetConversation(UserID) []History

	// starts listening to the server
	Run(cbRecvMessage func(history History))
}

// errors
func AlreadyLoggedIn(messages ...string) error {
	return &chatError{fmt.Sprint("already logged in", strings.Join(messages, " "))}
}

type chatError struct {
	s string
}

func (e *chatError) Error() string {
	return e.s
}
