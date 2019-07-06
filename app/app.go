package app

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime"

	"github.com/mattn/go-xmpp"

	"github.com/jdc-lab/coffee2go/conf"
	"github.com/jdc-lab/coffee2go/ui"
)

type app struct {
	ui     ui.Desktop
	client xmpp.Client
}

func New(args ...string) *app {
	a := &app{}

	if runtime.GOOS == "linux" {
		args = append(args, conf.LinuxAppendArgs)
	}
	var err error

	if a.ui, err = ui.New(conf.Width, conf.Height, args...); err != nil {
		log.Fatal(err)
	}
	return a
}

func (a *app) Run() {
	a.ui.Bind("run", func() {
		log.Printf("Starting UI")
	})
	listener, err := net.Listen("tcp", conf.NetAddr)

	if err != nil {
		log.Fatal(err)
	}
	go http.Serve(listener, http.FileServer(FS))
	a.ui.Load(fmt.Sprintf("http://%s", listener.Addr()))

	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)

	select {
	case <-sigc:
	case <-a.ui.Done():
	}

	defer a.ui.Close()
	defer listener.Close()
}
