package ui

import (
	"github.com/zserge/lorca"
)

type Desktop interface{}

type Lorca struct {
	inner lorca.UI
}

func New(width, height int, args ...string) (*Lorca, error) {
	ui, err := lorca.New("", "", width, height, args...)

	if err != nil {
		return nil, err
	}

	lorca := &Lorca{
		inner: ui,
	}
	return lorca, nil
}
