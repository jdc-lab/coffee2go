package ui

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/jdc-lab/coffee2go/conf"
)

type Bindings struct {
	Send  func(text string)
	Login func(server, username, password string)
}

type Controller struct {
	ui       ui
	bindings Bindings
	listener net.Listener
}

func NewController(width int, height int, bindings Bindings) (*Controller, error) {
	c := &Controller{}

	var err error

	if c.ui, err = NewLorca(width, height); err != nil {
		return nil, err
	}

	c.bindings = bindings

	return c, nil
}

func (c *Controller) Run() {
	if err := c.ui.Bind("run", func() {
		log.Printf("Starting UI")
	}); err != nil {
		log.Fatal(err)
	}
	defer c.ui.Close()

	c.setupBindings()

	var err error

	c.listener, err = net.Listen("tcp", conf.NetAddr)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := c.listener.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		if err := http.Serve(c.listener, http.FileServer(FS)); err != nil {
			log.Fatal(err)
		}
	}()

	if err = c.ui.Load(fmt.Sprintf("http://%s", c.listener.Addr())); err != nil {
		log.Fatal(err)
	}

	c.ui.Wait()
}

func (c *Controller) bind(name string, f interface{}) {
	if err := c.ui.Bind(name, f); err != nil {
		log.Fatal(err)
	}
}

func (c *Controller) setupBindings() {
	c.bind("goSend", c.bindings.Send)
	c.bind("goLogin", c.bindings.Send)
}

func (c *Controller) AppendHistory(history string) {
	c.ui.AppendHistory(history)
}

func (c *Controller) Login(server, username, password string) {
	url := fmt.Sprintf("http://%s/app.html", c.listener.Addr())
	c.ui.Load(url)
}
