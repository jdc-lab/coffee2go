package session

import (
	"errors"
	"github.com/google/uuid"
	"github.com/jdc-lab/coffee2go/uc"
)

type SessionHandler struct {
	sessions map[string]*uc.Chat
}

func New() *SessionHandler {
	return &SessionHandler{
		make(map[string]*uc.Chat, 0),
	}
}

func (s *SessionHandler) Add(session *uc.Chat) (sessionID string, err error) {
	newId, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}

	sessionID = newId.String()
	s.sessions[sessionID] = session

	return sessionID, nil
}

func (s *SessionHandler) Get(sessionID string) (session *uc.Chat, err error) {
	if session, ok := s.sessions[sessionID]; ok {
		return session, nil
	}

	return nil, errors.New("Invalid sessionID")
}
