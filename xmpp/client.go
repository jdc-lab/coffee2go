package xmpp

import (
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"log"
	"strings"

	"github.com/mattn/go-xmpp"
)

type Item struct {
	Jid          string   `xml:"jid,attr"`
	Name         string   `xml:"name,attr"`
	Subscription string   `xml:"subscription,attr"`
	Group        []string `xml:"group"`
}

type query struct {
	Xmlns string `xml:"xmlns,attr"`
	Ver   string `xml:"ver,attr"`
	Items []Item `xml:"item"`
}

type Client struct {
	xmpp.Client
	roster chan []Item
}

func serverName(host string) string {
	return strings.Split(host, ":")[0]
}

func NewClient(host string, username string, password string, insecureTLS bool) (*Client, error) {
	xmpp.DefaultConfig = tls.Config{
		ServerName:         serverName(host),
		InsecureSkipVerify: insecureTLS,
	}

	var c *xmpp.Client
	var err error

	options := xmpp.Options{
		Host:          host,
		User:          username,
		Password:      password,
		NoTLS:         false,
		Debug:         true,
		Session:       false,
		Status:        "xa",
		StatusMessage: "Hello",
	}

	if c, err = options.NewClient(); err != nil {
		return nil, err
	}

	return &Client{
		*c,
		make(chan []Item),
	}, nil
}

func (c *Client) Listen(msgRecvFunc func(message string)) {
	go func() {
		for {
			chat, err := c.Recv()
			if err != nil {
				log.Fatal(err)
			}

			switch v := chat.(type) {
			case xmpp.Chat:

				if len(v.Text) > 0 {
					msgRecvFunc(v.Remote + ": " + v.Text)
				}
			case xmpp.Presence:
				//fmt.Println(v.From, v.Show)
				fmt.Println("Not supported yet")
			case xmpp.Roster:
				fmt.Println("Roster: ", v)
			case xmpp.IQ:
				if v.Type == "result" {
					// parse query xml
					var q query
					err := xml.Unmarshal(v.Query, &q)
					if err != nil {
						fmt.Printf("error: %v", err)
						return
					}

					switch q.Xmlns {
					case "jabber:iq:roster":
						c.roster <- q.Items
					default:
						fmt.Println("Not supported yet", q)
					}
				}

			default: //
				fmt.Println("Not supported yet")
			}
		}
	}()
}

func (c *Client) RefreshRoster() []Item {
	if err := c.Roster(); err != nil {
		fmt.Println(err)
	}
	roster := <-c.roster

	return roster
}

func (c *Client) Send() {
}
