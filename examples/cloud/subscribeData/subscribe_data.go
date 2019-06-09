package main

import (
	"fmt"

	"github.com/joaoaneto/knot-cloud-sdk-go/config"
	"github.com/joaoaneto/knot-cloud-sdk-go/pkg/service/cloud"
)

func main() {
	// cloud config
	config := config.New("ws.cloud", 443, "ws")
	client := cloud.New(config)

	// connect to cloud by using your credentials
	client.Connect("d31e353f-03b0-4e6c-aa88-05e54868133c", "d671cd352906d1d2f1c71e11086d5494a8095a4f")

	// subscribe to data and receive them through a channel
	channel := client.SubscribeData()
	for {
		message := <-channel
		fmt.Println(message.Payload)
	}
}
