//go:generate go run -tags generate gen.go

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jdc-lab/coffee2go/app"
)

func main() {
	a, err := app.New()

	if err != nil {
		panic(err)
	}

	server := flag.String("server", "", `The server to connect with.`)
	username := flag.String("username", "", `The XMPP username.`)
	password := flag.String("password", "", `The corresponding password.`)

	fmt.Printf("Server Addr: %s", *server)
	fmt.Println(os.Args)

	a.Run(*server, *username, *password)
}
