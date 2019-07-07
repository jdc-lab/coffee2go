package ui

import (
	"log"
)

// Controller is the connection between ui and app.
// it mostly just inherits a ui, but it also provides some convenience-methods, which make sense for all ui's. (e.g. Bind)
type Controller struct {
	ui
}

func NewLorcaController(width int, height int) (*Controller, error) {
	g := &Controller{}

	var err error

	if g.ui, err = NewLorca(width, height); err != nil {
		return nil, err
	}

	return g, nil
}

func (c *Controller) Bind(name string, f interface{}) {
	// TODO: check if f is a func
	if err := c.ui.Bind("go"+name, f); err != nil {
		log.Fatal(err)
	}
}
