package knot

import (
	"github.com/gorilla/websocket"
)

// Request represents a WS request
type Request struct {
	Input interface{}
	Conn  *websocket.Conn
}

// NewRequest creates a new WS request
func NewRequest(input interface{}, conn *websocket.Conn) Request {
	return Request{input, conn}
}

// Send is responsible to send a message to WS server
func (r *Request) Send(data []byte) {
	r.Conn.WriteMessage(websocket.BinaryMessage, data)
}

// Read is responsible to read a message from WS server
func (r *Request) Read() []byte {
	_, message, err := r.Conn.ReadMessage()
	if err != nil {
		return nil
	}

	return message
}
