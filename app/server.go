package app

import (
	"flag"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"html/template"
	"net/http"
	"strings"
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

	FileServer(router, "/static", http.Dir("client/static"))
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	})
}

func parseFlags() (hostname, username, password *string) {
	hostname = flag.String("host", "", `The Server to connect with.`)
	username = flag.String("username", "", `The XMPP username.`)
	password = flag.String("password", "", `The corresponding password.`)

	flag.Parse()

	return
}
