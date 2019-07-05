package main

import (
	"github.com/zserge/lorca"
	"log"
	"net/url"
)

func main() {
	// Create UI with basic HTML passed via data URI
	ui, err := lorca.New("data:text/html,"+url.PathEscape(`
	<html>
		<head><title>coffee2go</title></head>
		<body><h1>Drink coffee2go and chat with XMPP</h1></body>
		
	</html>
	`), "", 480, 320)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()
	// Wait until UI window is closed
	<-ui.Done()
}
