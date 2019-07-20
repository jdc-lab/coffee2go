//+build generate

package main

import "github.com/aligator/lorca"

func main() {
	lorca.Embed("assets", "../assets/assets.go", "../fakewww")
}
