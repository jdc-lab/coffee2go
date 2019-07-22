package webview

import (
	"fmt"
	"github.com/zserge/webview"
)

type binding struct {
	name string
	f    interface{}
}

type Login struct {
	*Master
	bindCache []binding
}

func NewLogin(m *Master) *Login {
	l := Login{
		Master: m,
	}

	return &l
}

func (l *Login) Close() {
	fmt.Println("Close not implemented")
}

func (l *Login) Bind(name string, f interface{}) error {
	l.bindCache = append(l.bindCache, binding{
		name, f,
	})

	return nil
}

func (l *Login) PrefillForm(server, username, password string) {
	fmt.Println("pre not implemented")
}

func (l *Login) LoadLogin() {
	if l.window != nil {
		l.window.Terminate()
	}

	webview.Open("Coffee2Go", l.url, l.width, l.height, true)

	l.window.Dispatch(func() {
		println("bind 2")
		l.Bind("goLogin2", func(server, username, password string) {
			println(server + " " + username + " " + password)
		})

		// bind all
		for _, b := range l.bindCache {
			l.bind(b.name, b.f)
		}
	})

	//go func() {
	defer l.window.Exit()
	l.window.Run()
	//}()
}
