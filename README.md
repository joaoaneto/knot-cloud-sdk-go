# KNoT Cloud SDK for Go

knot-cloud-sdk-go is a SDK abstraction to the KNoT Cloud for Go programming language.

# Install the SDK

Run the command `glide install` to install all dependencies.

# Configuring the SDK

### Arguments
- `hostname` **String** KNoT Cloud host name.
- `port` **Number** KNoT Cloud port.
- `protocol` **String** Either `'ws'` or `'wss'`.

### Example

```go
package main

import (
	"fmt"

	"github.com/joaoaneto/knot-cloud-websocket-go/config"
	"github.com/joaoaneto/knot-cloud-websocket-go/pkg/service/cloud"
)

func main() {
  // Create a custom configuration with cloud address and protocol
  config := config.New("ws.knot.cloud", 443, "wss")
  // Using the config created above, we can create a cloud instance
  cloud := cloud.New(config)
}
```

# Methods

## connect(id, token): &lt;Void&gt;

Connects to the cloud instance by using user/app/gateway credentials (id/token). After connection is established, you will be able to enjoy the cloud capabilities.

### Example

```go
package main

import (
	"fmt"

	"github.com/joaoaneto/knot-cloud-websocket-go/config"
	"github.com/joaoaneto/knot-cloud-websocket-go/pkg/service/cloud"
)

func main() {
  config := config.New("ws.knot.cloud", 443, "wss")
  cloud := cloud.New(config)

  // Connects to the cloud
  cloud.Connect("d31e353f-03b0-4e6c-aa88-05e54868133c", "d671cd352906d1d2f1c71e11086d5494a8095a4f")
}
```

## listDevices(): &lt;Void&gt;

Lists the devices registered on cloud.

### Example

```go
package main

import (
	"fmt"

	"github.com/joaoaneto/knot-cloud-websocket-go/config"
	"github.com/joaoaneto/knot-cloud-websocket-go/pkg/service/cloud"
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

  // expected result:
  // {"type":"devices","data":[{"type":"knot:app","metadata":{"name":"sdadasdad"},"knot":{"id":"aaa4d757-d4bc-4b10-abca-bc308bfeef35","isThingManager":false}}]}
}
```

## SubscribeData(): &lt;Channel&gt;

Subscribe to data events sent by devices and receive it through a channel (FIFO).

### Example

```go
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
```
