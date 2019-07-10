//go:generate go run -tags generate gen.go

package main

import (
	"github.com/jdc-lab/coffee2go/app"
	"github.com/jdc-lab/coffee2go/conf"
	"github.com/jdc-lab/coffee2go/ui"
)

func main() {
	var uc *ui.Controller
	var err error

	if uc, err = ui.NewLorcaController(conf.Width, conf.Height); err != nil {
		panic(err)
	}

	a := app.New(*uc)
	a.Run()
}
