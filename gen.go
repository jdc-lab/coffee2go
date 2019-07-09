//+build generate

package main

import "github.com/aligator/lorca"

func main() {
	lorca.Embed("ui", "ui/assets.go", "www")
}
