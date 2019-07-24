package app

import (
	"flag"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"html/template"
	"net/http"
)


type Login struct {
	Hostname string `json:"hostname"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type session struct {
	token uuid.UUID
}

type Server struct {
	sessions []session
}

func NewServer() *Server {
	return &Server{}
}

// starts webserver
func (s *Server) Run() {
	router := chi.NewRouter()

	s.setupAPI(router)
	s.setupWeb(router)

	http.ListenAndServe(":8080", router)
}

func (s *Server) setupWeb(router *chi.Mux) {
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
}

func parseFlags() (hostname, username, password *string) {
	hostname = flag.String("host", "", `The Server to connect with.`)
	username = flag.String("username", "", `The XMPP username.`)
	password = flag.String("password", "", `The corresponding password.`)

	flag.Parse()

	return
}
