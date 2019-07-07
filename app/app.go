package app

import (
	"github.com/jdc-lab/coffee2go/conf"
	"github.com/jdc-lab/coffee2go/ui"
	"github.com/mattn/go-xmpp"
)

type App struct {
	ui     ui.Controller
	client xmpp.Client
}

func New() (*App, error) {
	a := &App{}

	var uc *ui.Controller
	var err error

	bindings := ui.Bindings{
		Send: a.Send,
	}

	if uc, err = ui.NewController(conf.Width, conf.Height, bindings); err != nil {
		return nil, err
	}

	a.ui = *uc
	return a, nil
}

func (a *App) Run() {
	a.ui.Run()
}

func (a *App) Send(text string) {
	println("SEND " + text)
}
