package ui

import "github.com/jdc-lab/coffee2go/xmpp"

type ui interface {
	Run(ready func())
}

type view interface {
	Close()
	Bind(name string, f interface{}) error
}

type loginUI interface {
	view
	PrefillForm(server, username, password string)
	LoadLogin()
}

type chatUI interface {
	view
	AppendHistory(bool, string)
	BuildRoster([]xmpp.Item)
	Select(jid string)
	LoadChat(servername, username string)
}
