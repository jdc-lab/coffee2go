package app

import (
	"github.com/go-chi/chi"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func (s *Server) setupWeb(router *chi.Mux, folder string) {
	// Handle static content, we have to explicitly put our top level dirs in here
	// - otherwise the NotFoundHandler will catch them
	FileServer(router, "/js", http.Dir(folder+"/js"))
	FileServer(router, "/css", http.Dir(folder+"/css"))
	FileServer(router, "/img", http.Dir(folder+"/img"))

	// EVERYTHING else redirect to index.html
	router.NotFound(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, folder+"/index.html")
	}))
}

func (s *Server) setupWebDev(router *chi.Mux, client *url.URL) {
	// EVERYTHING redirect to client
	router.NotFound(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		httputil.NewSingleHostReverseProxy(client).ServeHTTP(res, req)
	}))
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
