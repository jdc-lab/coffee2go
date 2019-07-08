package app

import (
	"log"

	"github.com/jdc-lab/coffee2go/conf"
	"github.com/jdc-lab/coffee2go/ui"
	"github.com/jdc-lab/coffee2go/xmpp"
)

type App struct {
	ui         ui.Controller
	client     *xmpp.Client
	roster     []xmpp.Item
	currentJid string
	histories  map[string][]string
}

func New(server, username, password string) (*App, error) {
	a := &App{}
	a.histories = make(map[string][]string)

	var uc *ui.Controller
	var err error

	if uc, err = ui.NewLorcaController(conf.Width, conf.Height); err != nil {
		return nil, err
	}

	a.ui = *uc

	// setup needed bindings (note: "go" is appended to each name)
	a.ui.Bind("Send", a.send)
	a.ui.Bind("Login", a.login)
	a.ui.Bind("OnLoginLoaded", func() {
		log.Printf("Starting Login UI")
		if server != "" || username != "" || password != "" {
			a.ui.PrefillForm(server, username, password)
		}
	})
	a.ui.Bind("OnAppLoaded", a.afterAppUiLoaded)

	return a, nil
}

func (a *App) Run() {
	a.ui.Run()
}

func (a *App) send(text string) {
	if a.client == nil {
		panic("This function should never be called if client is not logged in.")
	}
	a.client.Send()

	// If the chat history (identified by JID) exists,
	// append the new message text to the history.
	if h, ok := a.histories[a.currentJid]; ok {
		h = append(h, text)
	} else {
		// Otherwise, create a new history.
		a.histories[a.currentJid] = []string{
			text,
		}
	}

	a.ui.AppendHistory(false, text)
	// TODO: send message via xmpp
}

func (a *App) login(server, username, password string) {
	// todo: flag insecureTLS should be false in production (maybe offer flag for client in login screen)
	if client, err := xmpp.NewClient(server, username, password, true); err != nil {
		log.Println("Login failed: {}", err)
		// TODO: pass message to GUI
	} else {
		a.client = client
		a.ui.Login(server, username)
	}
}

func (a *App) afterAppUiLoaded() {
	log.Printf("Starting App UI")

	a.client.Listen(a.onMsgRecv)

	a.roster = a.client.RefreshRoster()
	a.ui.BuildRoster(a.roster)

	// set first one as current selected
	if len(a.roster) > 0 {
		a.currentJid = a.roster[0].Jid
	}
	a.ui.Select(a.currentJid)
}

func (a *App) onMsgRecv(msg xmpp.Chat) {

	// If the chat history (identified by Remote name) exist,
	// append the new message text to the history.
	if h, ok := a.histories[msg.Remote]; ok {
		h = append(h, msg.Text)
	} else {
		// Otherwise, create a new history.
		a.histories[msg.Remote] = []string{
			msg.Text,
		}
	}

	a.ui.AppendHistory(true, msg.Text)
}
