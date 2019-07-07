//+build generate

package main

import "github.com/zserge/lorca"

func main() {
	lorca.Embed("ui", "ui/assets.go", "www")
}
