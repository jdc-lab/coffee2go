package sse

import (
	"errors"
	"github.com/alexandrevicenzi/go-sse"
	"github.com/google/uuid"
	"log"
	"os"
	"time"
)

type PushHandler struct {
	pushTokens map[string]string
	sse        *sse.Server
}

func New() *PushHandler {
	return &PushHandler{
		make(map[string]string, 0),
		sse.NewServer(&sse.Options{
			Logger: log.New(os.Stdout, "go-sse: ", log.Ldate|log.Ltime|log.Lshortfile),
		}),
	}
}

func channelName(token string) string {
	return "/push/" + token
}

func (p PushHandler) Register(sessionID string) (pushToken string, err error) {
	newToken, err := uuid.NewRandom()
	if err != nil {
		return "", errors.New("could not generate pushToken")
	}

	p.pushTokens[sessionID] = newToken.String()
	return pushToken, nil
}

func (p PushHandler) Send(sessionID string, data string) (err error) {
	channel := channelName(p.pushTokens[sessionID])

	if !p.sse.HasChannel(channel) {
		// no need to send message if client is not connected
		return nil
	}

	msgID, err := uuid.NewRandom()
	if err != nil {
		log.Println("could not generate msgID")
		return nil
	}
	p.sse.SendMessage(channel, sse.NewMessage(msgID.String(), time.Now().String(), data))
}
