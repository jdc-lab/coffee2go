package app

import (
	"flag"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
)

type Server struct {
	r rest
}

func NewServer() *Server {
	return &Server{}
}

// starts webserver
func (s *Server) Run() {
	router := chi.NewRouter()

	s.r.setup(router)

	hostname, username, password := parseFlags()
	tmpl := template.Must(template.ParseFiles("client/index.html"))

	login := Login{
		*hostname,
		*username,
		*password,
	}
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// prefill login if login data is provided. (for development)
		tmpl.Execute(w, login)
	})

	fs := http.FileServer(http.Dir("client/static"))
	router.Handle("/static*", fs)

	http.ListenAndServe(":8080", router)
}

func parseFlags() (hostname, username, password *string) {
	hostname = flag.String("host", "", `The Server to connect with.`)
	username = flag.String("username", "", `The XMPP username.`)
	password = flag.String("password", "", `The corresponding password.`)

	flag.Parse()

	return
}
