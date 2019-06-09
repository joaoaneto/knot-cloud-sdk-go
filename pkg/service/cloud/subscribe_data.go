package cloud

import (
	"encoding/json"

	"github.com/joaoaneto/knot-cloud-sdk-go/knot"
)

// DataMessagePayload represents the devices request message payload
type DataMessagePayload struct {
	SensorID int         `json:"sensorId"`
	Value    interface{} `json:"value"`
}

// DataMessage represents the devices request message
type DataMessage struct {
	From    string             `json:"from"`
	Payload DataMessagePayload `json:"payload"`
}

// FrameResponse represents the devices request message
type FrameResponse struct {
	Type string      `json:"type"`
	Data DataMessage `json:"data"`
}

// SubscribeData request devices from cloud
func (c *Client) SubscribeData() chan *DataMessage {
	channel := make(chan *DataMessage, 1)
	subscribeRequest := &SubscribeRequest{}
	input := knot.NewFrame("data", subscribeRequest)

	req := SubscribeRequest{Request: c.NewRequest(input)}
	go req.Subscribe(channel)
	return channel
}

// SubscribeRequest represents the devices request
type SubscribeRequest struct {
	Request knot.Request
}

// Subscribe is responsible to serialize the devices message and send it to cloud
func (r *SubscribeRequest) Subscribe(ch chan<- *DataMessage) {
	for {
		parsedMessage := &FrameResponse{}
		message := r.Request.Read()

		if err := json.Unmarshal(message, &parsedMessage); err != nil {
			panic(err)
		}

		if parsedMessage.Type == "data" {
			ch <- &parsedMessage.Data
		}
	}
}
