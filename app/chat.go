package app

import (
	"github.com/jdc-lab/coffee2go/xmpp"
	"log"
)

type chat struct {
	*app
	client               *xmpp.Client
	roster               []xmpp.Item
	currentJid           string
	histories            map[string][]string
	servername, username string
}

func newChat(a *app, client *xmpp.Client, servername, username string) *chat {
	c := chat{app: a, client: client, servername: servername, username: username}

	c.histories = make(map[string][]string)

	// setup needed bindings (note: "go" is appended to each name)
	c.ui.Bind("Send", c.send)
	c.ui.Bind("OnAppLoaded", c.afterAppUiLoaded)

	return &c
}

func (c *chat) open() {
	c.ui.LoadChat(c.servername, c.username)
}

func (c *chat) close() {
	c.client.Close()
}

func (c *chat) send(text string) {
	if c.client == nil {
		panic("This function should never be called if client is not logged in.")
	}
	c.client.Send()

	// If the chat history (identified by JID) exists,
	// append the new message text to the history.
	if h, ok := c.histories[c.currentJid]; ok {
		h = append(h, text)
	} else {
		// Otherwise, create a new history.
		c.histories[c.currentJid] = []string{
			text,
		}
	}

	c.ui.AppendHistory(false, text)
	// TODO: send message via xmpp
}

func (c *chat) afterAppUiLoaded() {
	log.Printf("Starting app UI")

	c.client.Listen(c.onMsgRecv)

	c.roster = c.client.RefreshRoster()
	c.ui.BuildRoster(c.roster)

	// set first one as current selected
	if len(c.roster) > 0 {
		c.currentJid = c.roster[0].Jid
	}
	c.ui.Select(c.currentJid)
}

func (c *chat) onMsgRecv(msg xmpp.Chat) {

	// If the chat history (identified by Remote name) exist,
	// append the new message text to the history.
	if h, ok := c.histories[msg.Remote]; ok {
		h = append(h, msg.Text)
	} else {
		// Otherwise, create a new history.
		c.histories[msg.Remote] = []string{
			msg.Text,
		}
	}

	c.ui.AppendHistory(true, msg.Text)
}
