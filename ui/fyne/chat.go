package fyne

import "github.com/jdc-lab/coffee2go/xmpp"

type chat struct {
}

func (c *chat) Close() {
	panic("implement me")
}

func (c *chat) Bind(name string, f interface{}) error {
	panic("implement me")
}

func (c *chat) AppendHistory(bool, string) {
	panic("implement me")
}

func (c *chat) BuildRoster([]xmpp.Item) {
	panic("implement me")
}

func (c *chat) Select(jid string) {
	panic("implement me")
}

func (c *chat) LoadChat(servername, username string) {
	panic("implement me")
}
