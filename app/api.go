package app

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func (s *Server) setupAPI(router *chi.Mux) {
	router.Route("/api/login", func(router chi.Router) {
		s.setupLogin(router)
	})

	// TODO: auth by session token possible?
	router.Handle("/push/*", s.sse)

	// everything after logged in
	router.Route("/api/auth/{token}", func(router chi.Router) {
		router.Use(s.loggedIn)
		s.setupPushRegister(router)
		router.Get("/messages/new", s.fetchNewMessages)
	})
}

func (s *Server) fetchNewMessages(w http.ResponseWriter, r *http.Request) {
	/*if session, ok := s.session(w, r); ok {
		// TODO


	} else {
		http.Error(w, "", http.StatusForbidden)
		return
	}*/
}

// middleware to check if the given token is from a valid, already logged in user
func (s *Server) loggedIn(next http.Handler) http.Handler {
	// TODO: do not use session token in URL instead as param??
	// TODO: maybe also check username??
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := uuid.Parse(chi.URLParam(r, "token"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if _, ok := s.sessions[token]; !ok {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		ctx := context.WithValue(r.Context(), "token", token)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func token(w http.ResponseWriter, r *http.Request) (uuid.UUID, bool) {
	token, ok := r.Context().Value("token").(uuid.UUID)
	if !ok {
		log.Println("context value 'token' is not a  uuid")
		http.Error(w, "", http.StatusInternalServerError)
		return token, ok
	}
	return token, ok
}

func (s *Server) session(w http.ResponseWriter, r *http.Request) (*session, bool) {
	token, ok := token(w, r)
	if !ok {
		return nil, ok
	}

	sess, ok := s.sessions[token]
	return sess, ok
}
