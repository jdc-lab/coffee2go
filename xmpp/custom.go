package xmpp

type Conversation struct {
	History []Message
}

type Message struct {
	FromRemote bool
	Text       string
}
