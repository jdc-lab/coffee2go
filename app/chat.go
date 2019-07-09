package app

import (
	"fmt"
	"log"
	"strings"

	"github.com/jdc-lab/coffee2go/xmpp"
)

type chat struct {
	*app
	client               *xmpp.Client
	roster               []xmpp.Item
	selectedJID          string
	conversations        map[string]xmpp.Conversation
	servername, username string
}

func newChat(a *app, client *xmpp.Client, servername, username string) *chat {
	c := chat{app: a, client: client, servername: servername, username: username}

	c.conversations = make(map[string]xmpp.Conversation)

	// setup needed bindings (note: "go" is appended to each name)
	c.ui.Bind("Send", c.send)
	c.ui.Bind("OnChatLoaded", c.onChatLoaded)
	c.ui.Bind("LoadConversation", c.loadConversation)

	return &c
}

func (c *chat) open() {
	c.ui.LoadChat(c.servername, c.username)
}

func (c *chat) close() {
	log.Println("Closing chat")
	if err := c.client.Close(); err != nil {
		log.Println(err)
	}
}

func (c *chat) send(text string) {
	c.client.Send()

	msg := xmpp.Message{
		FromRemote: false,
		Text:       text,
	}

	// If the conversation (identified by JID) exists,
	// append the new message text to the conversation's history.
	if con, ok := c.conversations[c.selectedJID]; ok {
		con.History = append(con.History, msg)
	} else {
		// Otherwise, create a new conversation with a new history.
		c.conversations[c.selectedJID] = xmpp.Conversation{
			History: []xmpp.Message{msg},
		}
	}

	c.ui.AppendHistory(false, text)
	// TODO: send message via xmpp
}

func (c *chat) onChatLoaded() {
	log.Printf("\nStarting app UI\n")

	c.client.Listen(c.onMsgRecv)

	c.roster = c.client.RefreshRoster()
	c.ui.BuildRoster(c.roster)

	// set first one as current selected
	if len(c.roster) > 0 {
		c.selectedJID = c.roster[0].Jid
	}
	fmt.Printf("\n\n JID!!!! %s\n\n", c.roster[0].Jid)
	c.ui.Select(c.selectedJID)
}

func (c *chat) onMsgRecv(chat xmpp.Chat) {

	msg := xmpp.Message{
		FromRemote: true,
		Text:       chat.Text,
	}

	remoteJID := strings.Split(chat.Remote, "/")[0]

	// If the conversation (identified by remote name) exists,
	// append the new message text to the conversation's history.
	if con, ok := c.conversations[remoteJID]; ok {
		con.History = append(con.History, msg)
	} else {
		// Otherwise, create a new conversation with a new history.
		c.conversations[remoteJID] = xmpp.Conversation{
			History: []xmpp.Message{msg},
		}
	}

	if remoteJID == c.selectedJID {
		c.ui.AppendHistory(true, msg.Text)
	}
}

func (c *chat) loadConversation(jid string) *xmpp.Conversation {
	c.selectedJID = jid

	if con, ok := c.conversations[jid]; ok {
		return &con
	}
	return &xmpp.Conversation{
		History: []xmpp.Message{},
	}
}
