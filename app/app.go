package app

import (
	"fmt"
	"log"
	"sync"

	"github.com/jdc-lab/coffee2go/conf"
	"github.com/jdc-lab/coffee2go/ui"
	"github.com/jdc-lab/coffee2go/xmpp"
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
	client *xmpp.Client
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

	if a.client, err = xmpp.NewClient("127.0.0.1:5223", "braun@desktop-8dbsccu", "03110110", true); err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() {
	// TODO: this is only for testing xmpp
	go func() {
		for {
			chat, err := a.client.Recv()
			if err != nil {
				fmt.Println("lol")
				log.Fatal(err)
			}
			switch v := chat.(type) {
			case xmpp.Chat:
				fmt.Println(v.Remote, v.Text)
			case xmpp.Presence:
				fmt.Println(v.From, v.Show)
			}
		}
	}()
	a.ui.Run()
}
