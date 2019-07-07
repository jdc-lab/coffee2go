package app

import (
	"github.com/jdc-lab/coffee2go/conf"
	"github.com/jdc-lab/coffee2go/ui"
	"github.com/jdc-lab/coffee2go/xmpp"
)

type App struct {
	ui     ui.Controller
	client *xmpp.Client
}

func New() (*App, error) {
	a := &App{}

	var uc *ui.Controller
	var err error

	if uc, err = ui.NewController(conf.Width, conf.Height); err != nil {
		return nil, err
	}

	a.ui = *uc

	// seup needed bindings (note: "go" is appended to each name)
	a.ui.Bind("Send", a.send)
	a.ui.Bind("Login", a.login)

	if a.client, err = xmpp.NewClient("127.0.0.1:5223", "jh@localhost.localdomain", "jh", true); err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() {
	a.client.Listen(a.ui.AppendHistory)
	a.ui.Run()
}

func (a *App) send(text string) {
	a.ui.AppendHistory("Me: " + text)
	// TODO: send message via xmpp
}

func (a *App) login(server, username, password string) {
	a.ui.Login(server, username, password)
}
