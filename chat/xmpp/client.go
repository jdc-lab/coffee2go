package xmpp

import (
	"crypto/tls"
	"fmt"
	"github.com/jdc-lab/coffee2go/chat"
	"github.com/mattn/go-xmpp"
	"log"
	"strings"
)

type Client struct {
	InsecureTLS bool

	client *xmpp.Client
	onRecv func(history chat.History)
}

// splits the hostname on : to get only the host without port
// TODO: ipv6 supported by go-xmpp??
func serverName(host string) string {
	return strings.Split(host, ":")[0]
}

// Connects to a xmpp server using the provided host, username and password.
func (c *Client) Login(host, username, password string) error {
	if c.client != nil {
		return chat.AlreadyLoggedIn()
	}

	xmpp.DefaultConfig = tls.Config{
		ServerName:         serverName(host),
		InsecureSkipVerify: c.InsecureTLS,
	}

	var xc *xmpp.Client
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

	if xc, err = options.NewClient(); err != nil {
		return err
	}

	c.client = xc

	return nil
}

func (c Client) Send(to chat.UserID, message string) error {
	panic("implement me")
}

func (c Client) GetContacts() []chat.User {
	panic("implement me")
}

func (c Client) GetConversation(chat.UserID) []chat.History {
	panic("implement me")
}

func (c *Client) Run(cbRecvMessage func(history chat.History)) {
	go func() {
		for {
			message, err := c.client.Recv()
			if err != nil {
				// close on error (e.g. connection lost, if message closed)
				log.Println(err)
				return
			}

			switch v := message.(type) {
			case xmpp.Chat:

				if len(v.Text) > 0 {
					cbRecvMessage(chat.History{
						From:      chat.UserID(v.Remote),
						To:        chat.UserID(c.client.JID()),
						Message:   v.Text,
						Subject:   v.Subject,
						Timestamp: v.Stamp,
					})
				}

			case xmpp.Presence:
				fmt.Println("Not supported yet", v.From, v.Show)

			case xmpp.IQ:
				fmt.Println("Not supported yet")
				/*if v.Type == "result" {
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
				}*/

			default:
				fmt.Println("Not supported yet")
			}
		}
	}()
}

/*
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

type Chat struct {
	xmpp.Chat
}

type query struct {
	Xmlns string `xml:"xmlns,attr"`
	Ver   string `xml:"ver,attr"`
	Items []Item `xml:"item"`
}

type Client struct {
	xmpp.Client
	roster  chan []Item
	closing chan bool
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
		make(chan bool),
	}, nil
}

func (c *Client) Listen(msgRecvFunc func(Chat)) {
	go func() {
		for {
			chat, err := c.Recv()
			if err != nil {
				// close on error (e.g. connection lost, if chat closed)
				log.Println(err)
				return
			}

			switch v := chat.(type) {
			case xmpp.Chat:

				if len(v.Text) > 0 {
					msgRecvFunc(Chat{v})
				}

			case xmpp.Presence:
				fmt.Println("Not supported yet", v.From, v.Show)

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

			default:
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

func (c *Client) SendMessage(jid string, message string) {
	c.Send(xmpp.Chat{
		Remote: jid,
		Type:   "chat",
		Text:   message,
	})
}*/
