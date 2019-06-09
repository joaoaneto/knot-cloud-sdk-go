package cloud

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/joaoaneto/knot-cloud-sdk-go/config"
	"github.com/joaoaneto/knot-cloud-sdk-go/knot"
)

// Client represents the WS client state
type Client struct {
	address url.URL
	conn    *websocket.Conn
}

// New creates a new Client instance
func New(config config.KnotWSConfig) Client {
	host := config.Hostname + ":" + strconv.Itoa(config.Port)
	addr := flag.String("addr", host, "ws service address")
	u := url.URL{Scheme: config.Protocol, Host: *addr, Path: "/"}
	return Client{address: u}
}

// Connect establish a connection with WS server
func (c *Client) Connect(id string, token string) {
	wsHeaders := http.Header{
		"Origin":                   {c.address.Host},
		"Sec-WebSocket-Extensions": {"permessage-deflate; client_max_window_bits, x-webkit-deflate-frame"},
	}

	config := tls.Config{RootCAs: nil, InsecureSkipVerify: true}
	client, _ := tls.Dial("tcp", c.address.Host, &config)

	ws, _, err := websocket.NewClient(client, &c.address, wsHeaders, 1024, 1024)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	c.conn = ws
	c.Identity(id, token)
}

// NewRequest creates a new KNoT request instance
func (c *Client) NewRequest(input interface{}) knot.Request {
	return knot.Request{Input: input, Conn: c.conn}
}
