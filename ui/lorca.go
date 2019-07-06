package ui

import (
	"os"
	"os/signal"

	"github.com/zserge/lorca"
)

type Desktop interface {
	Bind(name string, f interface{}) error
	Load(url string) error
	Wait()
	Close()
}

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

func (l *Lorca) Bind(name string, f interface{}) error {
	err := l.inner.Bind(name, f)
	return err
}

func (l *Lorca) Load(url string) error {
	err := l.inner.Load(url)
	return err
}

func (l *Lorca) Wait() {
	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)

	select {
	case <-sigc:
	case <-l.inner.Done():
	}
}

func (l *Lorca) Close() {
	l.inner.Close()
}
