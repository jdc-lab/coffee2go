//go:generate go run -tags generate gen.go
//go:generate go get -u github.com/veily/go-bindata/...
//go:generate go-bindata -pkg webview -o ../../ui/webview/assets.go -prefix ../www/ ../www/...

package main

import (
	"github.com/jdc-lab/coffee2go/app"
	"github.com/jdc-lab/coffee2go/conf"
	"github.com/jdc-lab/coffee2go/ui"
)

func main() {
	var uc *ui.Controller
	var err error

	if uc, err = ui.NewWebviewController(conf.Width, conf.Height); err != nil {
		panic(err)
	}

	a := app.New(*uc)
	a.Run()
}
