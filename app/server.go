package app

import (
	"flag"
	"net/http"
)

type Server struct {
}

// starts webserver
func (s *Server) Run() {
	//hostname, username, password := parseFlags()

	fs := http.FileServer(http.Dir("client"))
	http.Handle("/", fs)

	http.ListenAndServe(":8080", nil)
}

func parseFlags() (hostname, username, password *string) {
	hostname = flag.String("host", "", `The Server to connect with.`)
	username = flag.String("username", "", `The XMPP username.`)
	password = flag.String("password", "", `The corresponding password.`)

	flag.Parse()

	return
}
