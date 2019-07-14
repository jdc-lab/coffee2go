package fyne

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/jdc-lab/coffee2go/xmpp"
	"log"
)

type Chat struct {
	*Master
	onChatLoaded     func()
	loadConversation func(jid string) *xmpp.Conversation
	send             *widget.Button
	input            *widget.Entry
	history          *widget.Entry
	roster           *fyne.Container
	sendBar          *fyne.Container
}

func NewChat(m *Master) *Chat {
	c := Chat{
		Master:  m,
		send:    widget.NewButton("Send", nil),
		input:   widget.NewEntry(),
		history: widget.NewMultiLineEntry(),
		roster:  fyne.NewContainerWithLayout(layout.NewVBoxLayout()),
	}

	c.history.SetReadOnly(true)
	c.sendBar = fyne.NewContainerWithLayout(layout.NewHBoxLayout(), c.input, c.send)

	// Doesn't exist: :-( c.input.SetOnTypedKey(c.typedKey)
	c.window.Canvas().SetOnTypedKey(c.typedKey)
	return &c
}

func (c *Chat) Close() {
	fmt.Println("Close not implemented")
}

func (c *Chat) Bind(name string, f interface{}) error {
	switch name {
	case "goSend":
		if f, ok := f.(func(text string)); ok {
			c.send.OnTapped = func() {
				text := c.input.Text
				f(text)
				c.input.SetText("")
			}
		} else {
			log.Fatal("Binding is not of correct function-type ", name)
		}
	case "goOnChatLoaded":
		if f, ok := f.(func()); ok {
			c.onChatLoaded = f
		} else {
			log.Fatal("Binding is not of correct function-type ", name)
		}
	case "goLoadConversation":
		if f, ok := f.(func(jid string) *xmpp.Conversation); ok {
			c.loadConversation = f
		} else {
			log.Fatal("Binding is not of correct function-typ ", name)
		}
	default:
		log.Fatal("Binding not implemented in fyne ", name)
	}
	return nil
}

func (c *Chat) AppendHistory(fromRemote bool, history string) {
	if fromRemote {
		history = "You: " + history
	} else {
		history = "Me: " + history
	}

	text := c.history.Text

	if text != "" {
		text += "\n"
	}

	text += history

	c.history.SetText(text)
}

func (c *Chat) BuildRoster(contacts []xmpp.Item) {

	for _, contact := range contacts {
		// cracy...
		contact := contact
		// ...
		c.roster.AddObject(widget.NewButton(contact.Name, func() {
			c.Select(contact.Jid, contact.Name)
		}))
	}

	// needed, otherwise sometimes the GUI is not fully loaded
	c.window.Resize(c.appSize)
}

func (c *Chat) Select(jid string, name string) {
	conversation := c.loadConversation(jid)
	c.history.SetText("")

	for _, msg := range conversation.History {
		c.AppendHistory(msg.FromRemote, msg.Text)
	}
}

func (c *Chat) LoadChat(servername, username string) {
	c.window.Resize(c.appSize)
	c.window.SetFixedSize(false)

	c.window.SetContent(fyne.NewContainerWithLayout(layout.NewBorderLayout(
		nil,
		c.sendBar,
		c.roster,
		nil),
		c.sendBar, c.roster, c.history))

	c.onChatLoaded()
}

func (c *Chat) typedKey(ev *fyne.KeyEvent) {
	if ev.Name == fyne.KeyReturn || ev.Name == fyne.KeyEnter {
		c.send.OnTapped()
	}
}
