package xmpp_authHandler

import (
	"github.com/jdc-lab/coffee2go/domain"
	"time"
)

//  Todo: Make configurable
var tokenTimeToLive = time.Hour * 2

// tokenHandler handles xmpp authentication request, implementing uc.AuthHandler interface
type tokenHandler struct {
}

func (t tokenHandler) UserLogin(host, username, password string) (user *domain.ChatConnection, err error) {
	panic("implement me")
}

func (t tokenHandler) Preset() (host, username, password string, err error) {
	panic("implement me")
}

func (t tokenHandler) GenUserToken(username string) (token string, err error) {
	panic("implement me")
}

func (t tokenHandler) GetUser(token string) (userName *domain.ChatConnection, err error) {
	panic("implement me")
}
