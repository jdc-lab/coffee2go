package fyne

import (
	"log"
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

type Login struct {
	*Master
	onLoginLoaded func()
	name          *widget.Label
	server        *widget.Entry
	username      *widget.Entry
	password      *widget.Entry
	login         *widget.Button
}

func NewLogin(m *Master) *Login {
	l := Login{
		Master:   m,
		name:     widget.NewLabel("Coffee2Go"),
		server:   widget.NewEntry(),
		username: widget.NewEntry(),
		password: widget.NewPasswordEntry(),
		login:    widget.NewButton("Drink Coffee2Go", nil),
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
	l.window.Resize(fyne.NewSize(300, 0))
	l.window.SetFixedSize(true)
	
	l.window.SetContent(fyne.NewContainerWithLayout(layout.NewGridLayout(1),
		fyne.NewContainerWithLayout(layout.NewFormLayout(),	
			widget.NewLabel("Server:"),
			l.server,
			widget.NewLabel("Username:"),
			l.username,
			widget.NewLabel("Passwort:"),
			l.password),
		fyne.NewContainerWithLayout(layout.NewCenterLayout(), l.login)))
	l.onLoginLoaded()
}
