//go:generate go run -tags generate gen.go

package main

import (
	"flag"
	"github.com/jdc-lab/coffee2go/app"
)

func main() {
	server := flag.String("server", "", `The server to connect with.`)
	username := flag.String("username", "", `The XMPP username.`)
	password := flag.String("password", "", `The corresponding password.`)

	flag.Parse()

	a, err := app.New()

	if err != nil {
		panic(err)
	}
	a.Run(*server, *username, *password)
}
