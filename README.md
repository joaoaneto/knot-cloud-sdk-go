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