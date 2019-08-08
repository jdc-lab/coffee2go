package app

import (
	"flag"
	"github.com/alexandrevicenzi/go-sse"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/jdc-lab/coffee2go/chat"
	"log"
	"net/http"
	"net/url"
	"os"
)

type Login struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type session struct {
	server    *Server
	client    chat.Client
	pushToken uuid.UUID
}

// used as callback if message received
func (sess *session) recvMessage(history chat.History) {
	// TODO: save history for later use (if client fetches it by api)
	sess.pushMessage("recv")
}

type Server struct {
	sessions    map[uuid.UUID]*session
	loginPreset Login
	sse         *sse.Server
}

func NewServer() *Server {
	return &Server{
		sessions: map[uuid.UUID]*session{},
		sse: sse.NewServer(&sse.Options{
			Logger: log.New(os.Stdout, "go-sse: ", log.Ldate|log.Ltime|log.Lshortfile),
		}),
	}
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
			s.setupWebClientProxy(router, client)
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
