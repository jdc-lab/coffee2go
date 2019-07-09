package app

import (
	"log"

	"github.com/jdc-lab/coffee2go/xmpp"
)

type chat struct {
	*app
	client               *xmpp.Client
	roster               []xmpp.Item
	currentJid           string
	conversations        map[string]xmpp.Conversation
	servername, username string
}

func newChat(a *app, client *xmpp.Client, servername, username string) *chat {
	c := chat{app: a, client: client, servername: servername, username: username}

	c.conversations = make(map[string]xmpp.Conversation)

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

	msg := xmpp.Message{
		FromRemote: false,
		Text:       text,
	}

	// If the conversation (identified by JID) exists,
	// append the new message text to the conversation's history.
	if con, ok := c.conversations[c.currentJid]; ok {
		con.History = append(con.History, msg)
	} else {
		// Otherwise, create a new conversation with a new history.
		c.conversations[c.currentJid] = xmpp.Conversation{
			History: []xmpp.Message{msg},
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

func (c *chat) onMsgRecv(chat xmpp.Chat) {

	msg := xmpp.Message{
		FromRemote: true,
		Text:       chat.Text,
	}

	// If the conversation (identified by remote name) exists,
	// append the new message text to the conversation's history.
	if con, ok := c.conversations[c.currentJid]; ok {
		con.History = append(con.History, msg)
	} else {
		// Otherwise, create a new conversation with a new history.
		c.conversations[c.currentJid] = xmpp.Conversation{
			History: []xmpp.Message{msg},
		}
	}

	c.ui.AppendHistory(true, msg.Text)
}

// func (c *chat) sendConversationData(jid string) *xmpp.Conversation {

// }
