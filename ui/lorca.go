package ui

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime"

	"github.com/jdc-lab/coffee2go/conf"
	"github.com/jdc-lab/coffee2go/xmpp"
	"github.com/zserge/lorca"
)

type ui interface {
	Bind(name string, f interface{}) error
	Close()
	Run()

	// methods to execute coffee2go specific actions
	AppendHistory(history string)
	PrefillForm(server, username, password string)
	Login(server string, username string)
	BuildRoster([]xmpp.Item)
	Select(jid string)
}

type Lorca struct {
	inner    lorca.UI
	listener net.Listener
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

func (l *Lorca) Run() {
	defer l.Close()

	var err error

	l.listener, err = net.Listen("tcp", conf.NetAddr)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := l.listener.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		if err := http.Serve(l.listener, http.FileServer(FS)); err != nil {
			log.Fatal(err)
		}
	}()

	if err = l.load(fmt.Sprintf("http://%s", l.listener.Addr())); err != nil {
		log.Fatal(err)
	}

	l.wait()
}

func (l *Lorca) Bind(name string, f interface{}) error {
	err := l.inner.Bind(name, f)
	return err
}

func (l *Lorca) Close() {
	l.inner.Close()
}

func (l *Lorca) AppendHistory(history string) {
	l.inner.Eval(fmt.Sprintf(`appendHistory(%q)`, history))
}

func (l *Lorca) PrefillForm(server, username, password string) {
	fn := fmt.Sprintf(`prefillForm(%q, %q, %q)`, server, username, password)
	l.inner.Eval(fn)
}

func (l *Lorca) BuildRoster(contacts []xmpp.Item) {
	fmt.Println("Building roster")
	for _, c := range contacts {
		fmt.Println("add contact", c)
		fn := fmt.Sprintf(`addContact(%q, %q, %q)`, c.Jid, c.Name, c.Subscription)
		l.inner.Eval(fn)
		fmt.Printf("\nJID: %s\n", c.Jid)
	}
}

func (l *Lorca) Select(jid string) {
	l.inner.Eval(fmt.Sprintf(`select(%q)`, jid))
}

// Login just switches from Login screen to main screen
func (l *Lorca) Login(server string, username string) {
	// TODO: server and usename needed (e.g. display somewhere in gui?)
	url := fmt.Sprintf("http://%s/%s", l.listener.Addr(), conf.AppFile)
	l.load(url)
}

func (l *Lorca) load(url string) error {
	err := l.inner.Load(url)
	return err
}

func (l *Lorca) wait() {
	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)

	select {
	case <-sigc:
	case <-l.inner.Done():
	}
}
