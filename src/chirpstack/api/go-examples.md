# Go examples

* [Go gRPC reference](https://pkg.go.dev/google.golang.org/grpc)
* [ChirpStack Go SDK reference](https://pkg.go.dev/github.com/chirpstack/chirpstack/api/go/v4)

## Enqueue downlink

The example below demonstrates:

* Configuration of gRPC _dial options_ including API token
* Connect to a gRPC API
* Define service client (in this case for the `DeviceService`)
* Perform an API call for a service (in this case `Enqueue`)

### `main.go`

```go
package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"github.com/chirpstack/chirpstack/api/go/v4/api"
)

// configuration
var (
	// This must point to the API interface
	server = "localhost:8080"

	// The DevEUI for which we want to enqueue the downlink
	devEUI = "0101010101010101"

	// The API token (retrieved using the web-interface)
	apiToken = "..."
)

type APIToken string

func (a APIToken) GetRequestMetadata(ctx context.Context, url ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": fmt.Sprintf("Bearer %s", a),
	}, nil
}

func (a APIToken) RequireTransportSecurity() bool {
	return false
}

func main() {
	// define gRPC dial options
	dialOpts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithPerRPCCredentials(APIToken(apiToken)),
		grpc.WithInsecure(), // remove this when using TLS
	}

	// connect to the gRPC server
	conn, err := grpc.Dial(server, dialOpts...)
	if err != nil {
		panic(err)
	}

	// define the DeviceService client
	deviceClient := api.NewDeviceServiceClient(conn)

	// make an Enqueue api call
	resp, err := deviceClient.Enqueue(context.Background(), &api.EnqueueDeviceQueueItemRequest{
		QueueItem: &api.DeviceQueueItem{
			DevEui:    devEUI,
			FPort:     10,
			Confirmed: false,
			Data:      []byte{0x01, 0x02, 0x03},
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("The downlink has been enqueued with id: %s\n", resp.Id)
}
```
