package app

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/jdc-lab/coffee2go/xmpp"
	"net/http"
)

func (s *Server) setupAPI(router *chi.Mux) {
	// logging in
	// starts connection to xmpp
	router.Post("/api/login", func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)

		var login Login
		err := decoder.Decode(&login)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// generate a new session token for this user
		token, err := uuid.NewRandom()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// connect to xmpp
		fmt.Println(login.Host, login.Username, login.Password)
		client, err := xmpp.NewClient(login.Host, login.Username, login.Password, true)
		if err != nil {
			// TODO: better handling of failed login
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		s.sessions = append(s.sessions, session{
			token,
			client,
		})

		// send new session token
		tokenJson, err := json.Marshal(struct {
			Token string `json:"token"`
		}{
			Token: token.String(),
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(tokenJson)
	})

	router.Get("/api/login/preset", func(w http.ResponseWriter, r *http.Request) {
		loginJson, err := json.Marshal(s.loginPreset)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(loginJson)
	})
}
