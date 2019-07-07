package app

import (
	"code.gitea.io/gitea/modules/log"
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

	if uc, err = ui.NewLorcaController(conf.Width, conf.Height); err != nil {
		return nil, err
	}

	a.ui = *uc

	// seup needed bindings (note: "go" is appended to each name)
	a.ui.Bind("Send", a.send)
	a.ui.Bind("Login", a.login)
	//"127.0.0.1:5223", "jh@localhost.localdomain", "jh"

	return a, nil
}

func (a *App) Run() {
	a.ui.Run()
}

func (a *App) send(text string) {
	if a.client == nil {
		panic("This function should never be called if client is not logged in.")
	}

	a.ui.AppendHistory("Me: " + text)
	// TODO: send message via xmpp
}

func (a *App) login(server, username, password string) {
	// todo: flag insecureTLS should be false in production (maybe offer flag for client in login screen)
	if client, err := xmpp.NewClient(server, username, password, true); err != nil {
		log.Error("Login failed: {}", err)
		// TODO: pass message to GUI
	} else {
		a.client = client
		a.client.Listen(a.ui.AppendHistory)
		a.ui.Login(server, username)
	}
}
