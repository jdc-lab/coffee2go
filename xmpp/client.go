package xmpp

import (
	"crypto/tls"
	"strings"

	"github.com/mattn/go-xmpp"
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
