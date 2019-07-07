package app

import (
	"github.com/jdc-lab/coffee2go/conf"
	"github.com/jdc-lab/coffee2go/ui"
	"github.com/mattn/go-xmpp"
	"sync"
)

type chatText struct {
	sync.Mutex
	text string
}

func (t *chatText) Set(text string) {
	t.Lock()
	defer t.Unlock()
	t.text = text
}

func (t *chatText) Get() string {
	t.Lock()
	defer t.Unlock()
	return t.text
}

type App struct {
	ui     *ui.Controller
	client xmpp.Client
	text   chatText
}

func New() (*App, error) {
	a := &App{}

	var uc *ui.Controller
	var err error

	bindings := ui.Bindings{
		Send:    a.text.Set,
		GetText: a.text.Get,
	}

	if uc, err = ui.NewController(conf.Width, conf.Height, bindings); err != nil {
		return nil, err
	}

	a.ui = uc
	return a, nil
}

func (a *App) Run() {
	a.ui.Run()
}
