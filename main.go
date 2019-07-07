//go:generate go run -tags generate gen.go

package main

import (
	"github.com/jdc-lab/coffee2go/app"
)

func main() {
	a, err := app.New()
	if err != nil {
		panic(err)
	}

	a.Run()
}
