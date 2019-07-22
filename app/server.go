package app

import (
	"flag"
	"github.com/go-chi/chi"
	"net/http"
)

type Server struct {
	r rest
}

func NewServer() *Server {
	hostname, username, password := parseFlags()
	return &Server{
		rest{
			Login{
				*hostname,
				*username,
				*password,
			},
		},
	}
}

// starts webserver
func (s *Server) Run() {
	router := chi.NewRouter()

	s.r.setup(router)

	fs := http.FileServer(http.Dir("client"))
	router.Handle("/*", fs)

	http.ListenAndServe(":8080", router)
}

func parseFlags() (hostname, username, password *string) {
	hostname = flag.String("host", "", `The Server to connect with.`)
	username = flag.String("username", "", `The XMPP username.`)
	password = flag.String("password", "", `The corresponding password.`)

	flag.Parse()

	return
}
