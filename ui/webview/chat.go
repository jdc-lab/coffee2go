package webview

import (
	"fmt"
	"github.com/jdc-lab/coffee2go/xmpp"
)

type Chat struct {
	*Master
}

func NewChat(m *Master) *Chat {
	c := Chat{
		Master: m,
	}

	return &c
}

func (c *Chat) Close() {
	fmt.Println("Close not implemented")
}

func (c *Chat) Bind(name string, f interface{}) error {
	fmt.Println("not implemented")
	return nil
}

func (c *Chat) AppendHistory(fromRemote bool, history string) {
	fmt.Println("not implemented")
}

func (c *Chat) BuildRoster(contacts []xmpp.Item) {
	fmt.Println("not implemented")
}

func (c *Chat) Select(jid string, name string) {
	fmt.Println("not implemented")
}

func (c *Chat) LoadChat(servername, username string) {
	fmt.Println("not implemented")
}
