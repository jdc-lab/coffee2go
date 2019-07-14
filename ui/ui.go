package ui

import "github.com/jdc-lab/coffee2go/xmpp"

type UI interface {
	Run(ready func())
}

type View interface {
	Close()
	Bind(name string, f interface{}) error
}

type LoginUI interface {
	View
	PrefillForm(server, username, password string)
	LoadLogin()
}

type ChatUI interface {
	View
	AppendHistory(bool, string)
	BuildRoster([]xmpp.Item)
	Select(jid string, name string)
	LoadChat(servername, username string)
}
