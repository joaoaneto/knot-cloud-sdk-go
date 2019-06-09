package cloud

import (
	"encoding/json"

	"github.com/joaoaneto/knot-cloud-sdk-go/knot"
)

// DevicesRequestMessage represents the devices request message
type DevicesRequestMessage struct {
}

// DevicesResponseMessage represents the devices response message
type DevicesResponseMessage struct {
}

// ListDevices request devices from cloud
func (c *Client) ListDevices() string {
	listDevicesMessage := &DevicesRequestMessage{}
	input := knot.NewFrame("devices", listDevicesMessage)

	req := DevicesRequest{Request: c.NewRequest(input), Frame: input}
	return req.Send()
}

// DevicesRequest represents the devices request
type DevicesRequest struct {
	Request knot.Request
	Frame   *knot.Frame
}

// Send is responsible to serialize the devices message and send it to cloud
func (r *DevicesRequest) Send() string {
	data, _ := json.Marshal(r.Frame)
	r.Request.Send(data)
	message := r.Request.Read()
	return string(message)
}
