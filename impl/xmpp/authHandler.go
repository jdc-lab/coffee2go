package xmpp

import (
	"crypto/tls"
	"github.com/jdc-lab/coffee2go/uc"
	"github.com/mattn/go-xmpp"
	"strings"
	"time"
)

//  Todo: Make configurable
var tokenTimeToLive = time.Hour * 2

// splits the hostname on : to get only the host without port
// TODO: ipv6 supported by go-xmpp??
func serverName(host string) string {
	return strings.Split(host, ":")[0]
}

type AuthHandler struct {
}

func (a *AuthHandler) Connect(host, username, password string) (serverConnection uc.Chat, err error) {
	xmpp.DefaultConfig = tls.Config{
		ServerName:         serverName(host),
		InsecureSkipVerify: true, // TODO make configurable
	}

	var xc *xmpp.Client

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

	if xc, err = options.NewClient(); err != nil {
		return nil, err
	}

	return &chatConnection{
		xc,
	}, nil
}
