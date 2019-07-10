package app

import (
	"flag"
	"github.com/jdc-lab/coffee2go/ui"
)

type module interface {
	open()
	close()
}

type app struct {
	ui     ui.Controller
	active module
}

func New(uc ui.Controller) *app {
	a := &app{
		ui: uc,
	}

	return a
}

func (a *app) changeModule(m module) {
	if a.active != nil {
		a.active.close()
	}

	a.active = m
	a.active.open()
}

func (a *app) parseFlags() (server, username, password *string) {
	server = flag.String("server", "", `The server to connect with.`)
	username = flag.String("username", "", `The XMPP username.`)
	password = flag.String("password", "", `The corresponding password.`)

	flag.Parse()

	return
}

func (a *app) Run() {
	a.ui.Main.Run(func() {
		server, username, password := a.parseFlags()
		a.active = newLogin(a, *server, *username, *password)
		a.active.open()
	})
}
