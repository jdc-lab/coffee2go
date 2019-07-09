package app

import (
	"github.com/jdc-lab/coffee2go/conf"
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

func New() (*app, error) {
	a := &app{}

	var uc *ui.Controller
	var err error

	if uc, err = ui.NewLorcaController(conf.Width, conf.Height); err != nil {
		return nil, err
	}

	a.ui = *uc

	return a, nil
}

func (a *app) changeModule(m module) {
	a.active.close()
	a.active = m
}

func (a *app) Run(server, username, password string) {
	a.ui.Run(func() {
		a.active = newLogin(a, server, username, password)
		a.active.open()
	})
}

func (a *App) conversationData(jid string)
