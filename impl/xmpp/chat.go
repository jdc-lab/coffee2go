package xmpp

import (
	"github.com/jdc-lab/coffee2go/domain"
	"github.com/mattn/go-xmpp"
	"log"
)

type chatConnection struct {
	client *xmpp.Client
}

func (c *chatConnection) Run() {
	go func() {
		for {
			_, err := c.client.Recv()
			if err != nil {
				// close on error (e.g. connection lost, if message closed)
				log.Println(err)
				return
			}
		}
	}()
}

func (c *chatConnection) Send(message domain.Message) (err error) {
	panic("implement me")
}

func (c *chatConnection) GetContacts() (err error) {
	panic("implement me")
}

func (c *chatConnection) SwitchChat() (err error) {
	panic("implement me")
}
