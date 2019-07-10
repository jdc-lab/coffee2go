package fyne

import (
	"code.gitea.io/gitea/modules/log"
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

type Login struct {
	*Master
	onLoginLoaded func()
	server        *widget.Entry
	username      *widget.Entry
	password      *widget.Entry
	login         *widget.Button
}

func NewLogin(m *Master) *Login {
	l := Login{
		Master:   m,
		server:   widget.NewEntry(),
		username: widget.NewEntry(),
		password: widget.NewPasswordEntry(),
		login:    widget.NewButton("go", nil),
	}

	return &l
}

func (l *Login) Close() {
	fmt.Println("Close not implemented")
}

func (l *Login) Bind(name string, f interface{}) error {
	switch name {
	case "goOnLoginLoaded":
		if f, ok := f.(func()); ok {
			l.onLoginLoaded = f
		} else {
			log.Fatal("Binding is not of correct function-type", name)
		}
	case "goLogin":
		if f, ok := f.(func(server, username, password string)); ok {
			l.login.OnTapped = func() {
				f(l.server.Text, l.username.Text, l.password.Text)
			}
		} else {
			log.Fatal("Binding is not of correct function-type", name)
		}
	default:
		log.Fatal("Binding not implemented in fyne", name)
	}
	return nil
}

func (l *Login) PrefillForm(server, username, password string) {
	l.server.SetText(server)
	l.username.SetText(username)
	l.password.SetText(password)
}

func (l *Login) LoadLogin() {
	l.window.SetContent(fyne.NewContainerWithLayout(layout.NewGridLayout(1),
		l.server, l.username, l.password,
		l.login))

	l.onLoginLoaded()
}
