//go:generate go run -tags generate gen.go

package main

import (
	"github.com/jdc-lab/coffee2go/app"
)

func main() {
	args := []string{}
	app.New(args...).Run()
}
