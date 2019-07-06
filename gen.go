//+build generate

package main

import "github.com/zserge/lorca"

func main() {
	lorca.Embed("app", "app/assets.go", "www")
}
