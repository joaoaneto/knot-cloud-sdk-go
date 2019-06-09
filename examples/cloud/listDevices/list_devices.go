package main

import (
	"fmt"

	"github.com/joaoaneto/knot-cloud-sdk-go/config"
	"github.com/joaoaneto/knot-cloud-sdk-go/pkg/service/cloud"
)

func main() {
	// cloud config
	config := config.New("ws.cloud", 443, "ws")
	cloud := cloud.New(config)

	// connect to cloud by using your credentials
	cloud.Connect("d31e353f-03b0-4e6c-aa88-05e54868133c", "d671cd352906d1d2f1c71e11086d5494a8095a4f")

	// list devices from cloud and print them
	devices := cloud.ListDevices()
	fmt.Println(devices)
}
