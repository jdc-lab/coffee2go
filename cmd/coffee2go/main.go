package main

import (
	"fmt"
	"github.com/jdc-lab/coffee2go/impl/session"
	"github.com/jdc-lab/coffee2go/impl/sse"
	"github.com/jdc-lab/coffee2go/impl/xmpp"
	"github.com/jdc-lab/coffee2go/uc"
	"time"
)

func main() {
	hc := uc.HandlerConstructor{
		Connection: &xmpp.AuthHandler{},
		Push:       &sse.PushHandler{},
		Conf:       &dummyConf{},
		Session:    session.New(),
	}

	handler := hc.New()
	sessionID, err := handler.ConnectServer("127.0.0.1:5223", "jh@localhost", "jh")

	if err != nil {
		panic(err)
	}

	fmt.Println(sessionID)

	for {
		time.Sleep(1)
	}
}
