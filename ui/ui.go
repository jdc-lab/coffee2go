package ui

import (
	"fmt"
	"github.com/jdc-lab/coffee2go/conf"
	"log"
	"net"
	"net/http"
)

type Bindings struct {
	Send func(text string)
}

type Controller struct {
	ui       ui
	bindings Bindings
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
	c.ui.Bind("run", func() {
		log.Printf("Starting UI")
	})
	defer c.ui.Close()

	c.setupBindings()

	listener, err := net.Listen("tcp", conf.NetAddr)
	defer listener.Close()

	if err != nil {
		log.Fatal(err)
	}
	go http.Serve(listener, http.FileServer(FS))
	c.ui.Load(fmt.Sprintf("http://%s", listener.Addr()))

	c.ui.Wait()
}

func (c *Controller) setupBindings() {
	c.ui.Bind("send", c.bindings.Send)
}
