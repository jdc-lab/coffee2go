package app

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime"

	"github.com/zserge/lorca"
)

type app struct {
	ui lorca.UI
}

func New(args ...string) *app {
	a := &app{}

	if runtime.GOOS == "linux" {
		args = append(args, "--class=Lorca")
	}
	var err error

	if a.ui, err = lorca.New("", "", 960, 540, args...); err != nil {
		log.Fatal(err)
	}
	return a
}

func (a *app) Run() {
	a.ui.Bind("run", func() {
		log.Printf("Starting UI")
	})
	listener, err := net.Listen("tcp", "127.0.0.1:0")

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
