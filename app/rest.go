package app

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
)

type Login struct {
	Hostname, Username, Password string
}

type rest struct {
	presetLogin Login
}

func (rs *rest) setup(router *chi.Mux) {
	router.Get("/api/login/getPrefill", func(w http.ResponseWriter, r *http.Request) {
		js, err := json.Marshal(rs.presetLogin)
		if err != nil {
			// todo: propper logging
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(js)
	})
}
