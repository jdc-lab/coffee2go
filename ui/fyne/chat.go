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
}

func NewChat(m *Master) *Chat {
	l := Chat{
		Master:  m,
		send:    widget.NewButton("Send", nil),
		input:   widget.NewEntry(),
		history: fyne.NewContainerWithLayout(layout.NewVBoxLayout()),
	}

	return &l
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

// TODO: remove fromRemote, as it's ui specific and should not be exposed
func (c *Chat) AppendHistory(fromRemote bool, history string) {
	c.history.AddObject(widget.NewLabel(history))
}

func (c *Chat) BuildRoster([]xmpp.Item) {
	fmt.Println("implement me")
}

func (c *Chat) Select(jid string) {
	fmt.Println("implement me")
}

func (c *Chat) LoadChat(servername, username string) {
	c.window.Resize(c.appSize)
	c.window.SetFixedSize(false)

	c.input.PlaceHolder = "Message"
	sendBar := fyne.NewContainerWithLayout(layout.NewHBoxLayout(), c.input, c.send)

	c.window.SetContent(fyne.NewContainerWithLayout(layout.NewBorderLayout(
		nil,
		sendBar,
		nil,
		nil),
		sendBar, c.history))

	c.onChatLoaded()
}
