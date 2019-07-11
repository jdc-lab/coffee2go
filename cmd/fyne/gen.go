//+build generate

package main

import "github.com/aligator/lorca"

func main() {
	lorca.Embed("assets", "../lorca/assets/assets.go", "fakewww")
}
