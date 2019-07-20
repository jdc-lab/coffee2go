package webview

import (
	"fmt"
)

type Login struct {
	*Master
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
	fmt.Println("bind not implemented")
	return nil
}

func (l *Login) PrefillForm(server, username, password string) {
	fmt.Println("pre not implemented")
}

func (l *Login) LoadLogin() {
	l.window.Dispatch(func() {
		l.window.Eval(`window.location.href = "` + l.url + `";`)
	})
}
