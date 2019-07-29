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
	From      User
	To        User
	Message   string
	timestamp time.Time
}

type Client interface {
	// Connect to a chat server by using the provided host, username and password.
	// Should return AlreadyLoggedIn error if called a second time.
	Login(host, username, password string) error
	Send(to UserID, message string) error
	OnRecv(func(from User)) // TODO: Recv as callback or with channels?
	GetContacts() []User
	GetConversation(UserID) []History
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
