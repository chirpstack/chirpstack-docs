# Azure Service-Bus

The [Azure Service Bus](https://azure.microsoft.com/en-us/services/service-bus/)
integration publishes all the events to a Service Bus [Topic or Queue](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-messaging-overview)
to which applications can subscribe.

## Events

The Azure Service Bus integration exposes all events as documented by [Event types](events.md).

### User properties

The following user properties are added to each published message:

* `event` - the event type
* `dev_eui` - the device EUI
* `application_id` - the ChirpStack Application Server application ID

## Example code

The following code example demonstrates how to consume integration events using
an [Azure Service-Bus Queue](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-queues-topics-subscriptions).

### Go

#### `main.go`

```go
package main

import (
	"bytes"
	"context"
	"encoding/hex"
	"log"

	servicebus "github.com/Azure/azure-service-bus-go"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"

	"github.com/brocaar/chirpstack-api/go/v3/as/integration"
)

type handler struct {
	json bool

	ns    *servicebus.Namespace
	queue *servicebus.Queue
}

func (h *handler) receive() error {
	for {
		if err := h.queue.Receive(context.TODO(), servicebus.HandlerFunc(h.receiveFunc)); err != nil {
			return err
		}
	}
}

func (h *handler) receiveFunc(ctx context.Context, msg *servicebus.Message) error {
	ev, ok := msg.UserProperties["Event"]
	if !ok {
		log.Println("event attribute is missing")
	}

	event, ok := ev.(string)
	if !ok {
		log.Println("event must be of type string")
	}

	var err error

	switch event {
	case "up":
		err = h.up(msg.Data)
	case "join":
		err = h.join(msg.Data)
	default:
		log.Printf("handler for event %s is not implemented", event)
	}

	if err != nil {
		log.Printf("handling event '%s' returned error: %s", event, err)
	}

	return msg.Complete(ctx)
}

func (h *handler) up(b []byte) error {
	var up integration.UplinkEvent
	if err := h.unmarshal(b, &up); err != nil {
		return err
	}
	log.Printf("Uplink received from %s with payload: %s", hex.EncodeToString(up.DevEui), hex.EncodeToString(up.Data))
	return nil
}

func (h *handler) join(b []byte) error {
	var join integration.JoinEvent
	if err := h.unmarshal(b, &join); err != nil {
		return err
	}
	log.Printf("Device %s joined with DevAddr %s", hex.EncodeToString(join.DevEui), hex.EncodeToString(join.DevAddr))
	return nil
}

func (h *handler) unmarshal(b []byte, v proto.Message) error {
	if h.json {
		unmarshaler := &jsonpb.Unmarshaler{
			AllowUnknownFields: true, // we don't want to fail on unknown fields
		}
		return unmarshaler.Unmarshal(bytes.NewReader(b), v)
	}
	return proto.Unmarshal(b, v)
}

func newHandler(json bool, connStr string, queueName string) (*handler, error) {
	ns, err := servicebus.NewNamespace(
		servicebus.NamespaceWithConnectionString(connStr),
	)
	if err != nil {
		panic(err)
	}

	queue, err := ns.NewQueue(queueName)
	if err != nil {
		panic(err)
	}

	return &handler{
		json:  json,
		ns:    ns,
		queue: queue,
	}, nil

}

func main() {
	h, err := newHandler(
		// set true when using JSON encoding
		false,

		// service-bus connection string
		"Endpoint=sb://example.servicebus.windows.net/;SharedAccessKeyName=example-policy;SharedAccessKey=...",

		// queue name
		"events",
	)
	if err != nil {
		panic(err)
	}
	panic(h.receive())
}
```

#### `go.mod`

```text
module example

go 1.14

require (
	github.com/Azure/azure-service-bus-go v0.10.3
	github.com/brocaar/chirpstack-api/go/v3 v3.7.2
	github.com/golang/protobuf v1.3.5
)
```

### Python

#### `main.py`

```python
from azure.servicebus import ServiceBusClient
from chirpstack_api.as_pb import integration
from google.protobuf.json_format import Parse


class Handler:
    def __init__(self, json, connection_string, queue_name):
        self.json = json
        self.connection_string = connection_string
        self.queue_name = queue_name

    def receive(self):
        client = ServiceBusClient.from_connection_string(self.connection_string)
        queue_client = client.get_queue(self.queue_name)
        messages = queue_client.get_receiver()

        for message in messages:
            message.complete()
            body = b''.join(message.body)

            event = message.user_properties[b'Event']

            if event == b'up':
                self.up(body)
            elif event == b'join':
                self.join(body)
            else:
                print('handler for event %s is not implemented' % event)

    def up(self, body):
        up = self.unmarshal(body, integration.UplinkEvent())
        print('Uplink received from: %s with payload: %s' % (up.dev_eui.hex(), up.data.hex()))

    def join(self, body):
        join = self.unmarshal(body, integration.JoinEvent())
        print('Device: %s joined with DevAddr: %s' % (join.dev_eui.hex(), join.dev_addr.hex()))

    def unmarshal(self, body, pl):
        if self.json:
            return Parse(body, pl)
        
        pl.ParseFromString(body)
        return pl


h = Handler(
    # True -  JSON marshaler
    # False - Protobuf marshaler (binary)
    False,

    # Service-Bus connection string
    'Endpoint=sb://example.servicebus.windows.net/;SharedAccessKeyName=example-policy;SharedAccessKey=...',

    # Service-Bus queue name
    'events',
)
h.receive()
```

#### `requirements.txt`

```text
chirpstack-api==3.7.2
azure-servicebus==0.50.3
```
