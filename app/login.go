package app

import (
	"github.com/jdc-lab/coffee2go/xmpp"
	"log"
)

type login struct {
	*app

	// presets
	server, username, password string
}

func newLogin(a *app, server, username, password string) *login {
	l := login{
		app:      a,
		server:   server,
		username: username,
		password: password,
	}

	// setup needed bindings (note: "go" is appended to each name)
	l.ui.LoginBind("Login", l.login)
	l.ui.LoginBind("OnLoginLoaded", l.afterLoginUiLoaded)

	return &l
}

func (l *login) open() {
	l.ui.Login.LoadLogin()
}

func (l *login) close() {
}

func (l *login) afterLoginUiLoaded() {
	log.Printf("Starting login UI")
	if l.server != "" || l.username != "" || l.password != "" {
		l.ui.Login.PrefillForm(l.server, l.username, l.password)
	}
}

func (l *login) login(server, username, password string) {
	// todo: flag insecureTLS should be false in production (maybe offer flag for client in login screen)
	if client, err := xmpp.NewClient(server, username, password, true); err != nil {
		log.Println("login failed: {}", err)
		// TODO: pass message to GUI
	} else {
		c := newChat(l.app, client, server, username)
		l.app.changeModule(c)
	}
}
