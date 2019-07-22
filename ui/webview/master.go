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
	"reflect"
)

type Master struct {
	window webview.WebView
	width  int
	height int
	url    string
	exit   bool
}

type Binding struct {
	sync      func()
	Returning []interface{} `json:"returning"`
	callback  interface{}
}

func (b *Binding) GoCall(params []interface{}) {
	args := make([]reflect.Value, 0)
	for _, param := range params {
		args = append(args, reflect.ValueOf(param))
	}

	rCallback := reflect.ValueOf(b.callback)
	rResult := rCallback.Call(args)

	result := make([]interface{}, 0)
	for _, res := range rResult {
		result = append(result, res.Interface())
	}

	b.Returning = result
	if b.sync != nil {
		b.sync()
	}
}

func NewWebview(width int, height int) (*Master, error) {
	m := &Master{
		width:  width,
		height: height,
	}

	return m, nil
}

func (m *Master) Run(ready func()) {
	m.url = startServer()
	ready()
	for !m.exit {
	}
	/*
		defer m.window.Exit()

			ready()
			m.window.Run()*/
}

func (m *Master) bind(name string, f interface{}) {
	b := &Binding{
		callback:  f,
		Returning: make([]interface{}, 0),
	}
	//m.window.Dispatch(func() {
	sync, err := m.window.Bind(name+"Webview", b)
	if err == nil {
		b.sync = sync
	} else {
		println("Error while binding " + name)
	}
	//})

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
