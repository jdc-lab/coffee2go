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
	history          *fyne.Container
	roster           *fyne.Container
	sendBar          *fyne.Container
}

func NewChat(m *Master) *Chat {
	c := Chat{
		Master: m,
		send:   widget.NewButton("Send", nil),
		input:  widget.NewEntry(),

		history: fyne.NewContainerWithLayout(layout.NewVBoxLayout()),
		roster:  fyne.NewContainerWithLayout(layout.NewVBoxLayout()),
	}

	c.sendBar = fyne.NewContainerWithLayout(layout.NewHBoxLayout(), c.input, c.send)

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
				f(c.input.Text)
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

	c.history.AddObject(widget.NewLabel(history))
}

func (c *Chat) BuildRoster(contacts []xmpp.Item) {

	for _, contact := range contacts {
		// cracy...
		contact := contact
		// ...
		c.roster.AddObject(widget.NewButton(contact.Name, func() {
			c.Select(contact.Jid)
		}))
	}
}

func (c *Chat) Select(jid string) {
	conversation := c.loadConversation(jid)
	//c.history = fyne.NewContainerWithLayout(layout.NewVBoxLayout())

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
