package ui

import (
	"github.com/jdc-lab/coffee2go/conf"
	"github.com/zserge/lorca"
	"os"
	"os/signal"
	"runtime"
)

type ui interface {
	Bind(name string, f interface{}) error
	Load(url string) error
	Wait()
	Close()
}

type Lorca struct {
	inner lorca.UI
}

func NewLorca(width, height int, args ...string) (*Lorca, error) {
	if runtime.GOOS == "linux" {
		args = append(args, conf.LinuxAppendArgs)
	}

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
