package app

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"net/http"
)

type Login struct {
	Hostname string `json:"hostname"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type rest struct{}

func (rs *rest) setup(router *chi.Mux) {
	router.Post("/api/login", func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)

		var login Login
		err := decoder.Decode(&login)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		// TODO: connect to xmpp
		token, err := uuid.NewRandom()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		tokenJson, err := json.Marshal(struct {
			Token string `json:"token"`
		}{
			Token: token.String(),
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Write(tokenJson)
	})
}
