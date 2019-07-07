package app

import (
	"sync"

	"github.com/jdc-lab/coffee2go/conf"
	"github.com/jdc-lab/coffee2go/ui"
	"github.com/jdc-lab/coffee2go/xmpp"
)

type chatText struct {
	sync.Mutex
	text string
}

type App struct {
	ui     ui.Controller
	client *xmpp.Client
	text   chatText
}

func New() (*App, error) {
	a := &App{}

	var uc *ui.Controller
	var err error

	bindings := ui.Bindings{
		Send: a.send,
	}

	if uc, err = ui.NewController(conf.Width, conf.Height, bindings); err != nil {
		return nil, err
	}

	a.ui = *uc

	if a.client, err = xmpp.NewClient("127.0.0.1:5223", "braun@desktop-8dbsccu", "03110110", true); err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() {
	a.client.Listen()
	a.ui.Run()
}

func (a *App) send(text string) {
	a.ui.AppendHistory(text)
	// TODO: send message via xmpp
}
