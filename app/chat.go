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
	selected             xmpp.Item
	conversations        map[string]xmpp.Conversation
	servername, username string
}

func newChat(a *app, client *xmpp.Client, servername, username string) *chat {
	c := chat{app: a, client: client, servername: servername, username: username}

	c.conversations = make(map[string]xmpp.Conversation)

	// setup needed bindings (note: "go" is appended to each name)
	c.ui.ChatBind("Send", c.send)
	c.ui.ChatBind("OnChatLoaded", c.onChatLoaded)
	c.ui.ChatBind("LoadConversation", c.loadConversation)

	return &c
}

func (c *chat) open() {
	c.ui.Chat.LoadChat(c.servername, c.username)
}

func (c *chat) close() {
	log.Println("Closing chat")
	if err := c.client.Close(); err != nil {
		log.Println(err)
	}
}

func (c *chat) send(text string) {
	if text == "" {
		return
	}

	c.client.SendMessage(c.selected.Jid, text)

	msg := xmpp.Message{
		FromRemote: false,
		Text:       text,
	}

	// If the conversation (identified by JID) exists,
	// append the new message text to the conversation's history.
	if con, ok := c.conversations[c.selected.Jid]; ok {
		con.History = append(con.History, msg)
		c.conversations[c.selected.Jid] = con
	} else {
		// Otherwise, create a new conversation with a new history.
		c.conversations[c.selected.Jid] = xmpp.Conversation{
			History: []xmpp.Message{msg},
		}
	}

	c.ui.Chat.AppendHistory(false, text)
}

func (c *chat) onChatLoaded() {
	log.Printf("\nStarting app UI\n")

	c.client.Listen(c.onMsgRecv)

	c.roster = c.client.RefreshRoster()
	c.ui.Chat.BuildRoster(c.roster)

	// set first one as current selected
	if len(c.roster) > 0 {
		c.selected = c.roster[0]
	}
	fmt.Printf("\n\n JID!!!! %s\n\n", c.roster[0].Jid)
	c.ui.Chat.Select(c.selected.Jid, c.selected.Name)
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
		c.conversations[remoteJID] = con
	} else {
		// Otherwise, create a new conversation with a new history.
		c.conversations[remoteJID] = xmpp.Conversation{
			History: []xmpp.Message{msg},
		}
	}

	if remoteJID == c.selected.Jid {
		c.ui.Chat.AppendHistory(true, msg.Text)
	}
}

func (c *chat) loadConversation(jid string) *xmpp.Conversation {
	c.selected.Jid = jid

	if con, ok := c.conversations[jid]; ok {
		return &con
	}
	return &xmpp.Conversation{
		History: []xmpp.Message{},
	}
}
