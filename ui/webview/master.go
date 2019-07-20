package webview

import (
	"bytes"
	"github.com/zserge/webview"
	"io"
	"log"
	"mime"
	"net"
	"net/http"
	"path/filepath"
)

type Master struct {
	window webview.WebView
	url    string
}

func NewWebview(width int, height int) (*Master, error) {
	w := webview.New(webview.Settings{
		Title:     "Coffee2Go",
		Width:     width,
		Height:    height,
		Resizable: true,
		Debug:     true,
		URL:       `data:text/html,<html><script type="text/javascript"></script></html>`,
	})

	m := &Master{
		window: w,
	}

	return m, nil
}

func (m *Master) Run(ready func()) {
	m.url = startServer()
	defer m.window.Exit()

	ready()
	m.window.Run()
}

func startServer() string {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		defer ln.Close()
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			// TODO: maybe include something like a session key to authenticate only this app and not another external browser
			path := r.URL.Path
			if len(path) > 0 && path[0] == '/' {
				path = path[1:]
			}
			if path == "" {
				path = "index.html"
			}
			if bs, err := Asset(path); err != nil {
				w.WriteHeader(http.StatusNotFound)
			} else {
				w.Header().Add("Content-Type", mime.TypeByExtension(filepath.Ext(path)))
				io.Copy(w, bytes.NewBuffer(bs))
			}
		})

		log.Fatal(http.Serve(ln, nil))
	}()

	return "http://" + ln.Addr().String()
}
