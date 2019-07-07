package xmpp

import (
	"crypto/tls"
	"fmt"
	"github.com/mattn/go-xmpp"
	"log"
	"strings"
)

type Chat struct {
	xmpp.Chat
}

type Presence struct {
	xmpp.Presence
}

type Client struct {
	xmpp.Client
}

type Roster struct {
	xmpp.Client
}

func serverName(host string) string {
	return strings.Split(host, ":")[0]
}

func NewClient(host string, username string, password string, insecureTLS bool) (*Client, error) {
	xmpp.DefaultConfig = tls.Config{
		ServerName:         serverName(host),
		InsecureSkipVerify: insecureTLS,
	}

	var c *xmpp.Client
	var err error

	options := xmpp.Options{
		Host:          host,
		User:          username,
		Password:      password,
		NoTLS:         false,
		Debug:         true,
		Session:       false,
		Status:        "xa",
		StatusMessage: "Hello",
	}

	if c, err = options.NewClient(); err != nil {
		return nil, err
	}

	return &Client{*c}, nil
}

func (c *Client) Listen(msgRecvFunc func(message string)) {
	go func() {
		for {
			chat, err := c.Recv()
			if err != nil {
				log.Fatal(err)
			}

			switch v := chat.(type) {
			case xmpp.Chat:

				if len(v.Text) > 0 {
					msgRecvFunc(v.Remote + ": " + v.Text)
				}
			case xmpp.Presence:
				//fmt.Println(v.From, v.Show)
				fmt.Println("Not supported yet")
			case xmpp.Roster:
				fmt.Println(v)
			default: //
				fmt.Println("Not supported yet")
			}
		}
	}()
}

func (c *Client) Send() {
}
