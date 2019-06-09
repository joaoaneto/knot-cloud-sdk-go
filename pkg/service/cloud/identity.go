package cloud

import (
	"encoding/json"
	"fmt"

	"github.com/joaoaneto/knot-cloud-sdk-go/knot"
)

// IdentityRequestMessage represents the identity request message
type IdentityRequestMessage struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}

// Identity send a authentication request to cloud
func (c *Client) Identity(id string, token string) {
	var identityMessage *IdentityRequestMessage

	if len(id) <= 0 || len(token) <= 0 {
		identityMessage = &IdentityRequestMessage{}
	}

	identityMessage = &IdentityRequestMessage{id, token}
	input := knot.NewFrame("identity", identityMessage)

	req := IdentityRequest{Request: c.NewRequest(input), Frame: input}
	req.Send()
}

// IdentityRequest represents the identity request
type IdentityRequest struct {
	Request knot.Request
	Frame   *knot.Frame
}

// Send is responsible to serialize the identity message and send it to cloud
func (r *IdentityRequest) Send() {
	data, _ := json.Marshal(r.Frame)
	r.Request.Send(data)
	message := r.Request.Read()
	fmt.Println(string(message))
}
