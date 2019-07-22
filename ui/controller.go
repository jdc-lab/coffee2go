package ui

import (
	"github.com/jdc-lab/coffee2go/ui/webview"
	"log"
)

// Controller is the connection between UI and app.
// it mostly just inherits a UI, but it also provides some convenience-methods, which make sense for all UI's. (e.g. Bind)
type Controller struct {
	Main  UI
	Login LoginUI
	Chat  ChatUI
}

func NewLorcaController(width int, height int) (*Controller, error) {
	g := &Controller{}

	var err error
	var ui *Lorca
	if ui, err = NewLorca(width, height); err != nil {
		return nil, err
	}

	g.Main = ui

	g.Login = LorcaLogin{
		ui,
	}

	g.Chat = LorcaChat{
		ui,
	}

	return g, nil
}

func NewWebviewController(width int, height int) (*Controller, error) {
	g := &Controller{}

	var err error
	var ui *webview.Master
	if ui, err = webview.NewWebview(width, height); err != nil {
		return nil, err
	}

	g.Main = ui
	g.Login = webview.NewLogin(ui)
	g.Chat = webview.NewChat(ui)

	return g, nil
}

func (c *Controller) ChatBind(name string, f interface{}) {
	c.bind(name, f, c.Chat)
}

func (c *Controller) LoginBind(name string, f interface{}) {
	c.bind(name, f, c.Login)
}

func (c *Controller) bind(name string, f interface{}, u View) {
	// TODO: check if f is a func
	// TODO: wrap function  with another one which checks if we are in the correct module currently and prevents running the binding, if not
	if err := u.Bind("go"+name, f); err != nil {
		log.Fatal(err)
	}
}
