package app

import (
	"flag"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/jdc-lab/coffee2go/xmpp"
	"net/http"
	"net/url"
)

type Login struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type session struct {
	token uuid.UUID
	// todo use interface between xmpp and server to make the chat more generic
	chat *xmpp.Client
}

type Server struct {
	sessions    []session
	loginPreset Login
}

func NewServer() *Server {
	return &Server{}
}

// starts webserver
func (s *Server) Run() {
	router := chi.NewRouter()

	devURL, host, username, password := parseFlags()
	s.loginPreset = Login{
		Host:     *host,
		Username: *username,
		Password: *password,
	}

	s.setupAPI(router)

	if *devURL == "" {
		s.setupWeb(router, "client/dist")
	} else {
		client, err := url.Parse(*devURL)
		if err == nil {
			s.setupWebDev(router, client)
		}
	}

	http.ListenAndServe(":8080", router)
}

func parseFlags() (devURL, hostname, username, password *string) {
	devURL = flag.String("dev-url", "", "Enables using another url for the client. This is for starting the client-server externally in dev mode.")
	hostname = flag.String("host", "", `The Server to connect with.`)
	username = flag.String("username", "", `The XMPP username.`)
	password = flag.String("password", "", `The corresponding password.`)

	flag.Parse()

	return
}
