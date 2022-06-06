# GCP Pub/Sub

The [Google Cloud Platform](https://cloud.google.com/) [Pub/Sub](https://cloud.google.com/pubsub/)
integration publishes all the events to a configurable GCP Pub/Sub topic.
Using the GCP console (or APIs) you are able to create one or multiple Pub/Sub
subscriptions, for integrating this with your application(s) or store the data
in one of the storage options provided by the Google Cloud Platform.

## Events

The GCP Pub/Sub integration exposes all events as documented by [Event types](events.md).

### Attributes

The following attributes are added to each Pub/Sub message:

* `event`: the event type
* `devEui`: the device EUI to which the event relates

## Example code

The following code example demonstrates how to consume integration events using
a [GCP Pub/Sub Subscription](https://cloud.google.com/pubsub/docs/overview).

### Go

#### `main.go`

```go
package main

import (
	"bytes"
	"context"
	"encoding/hex"
	"errors"
	"log"

	"cloud.google.com/go/pubsub"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"google.golang.org/api/option"

	"github.com/brocaar/chirpstack-api/go/v3/as/integration"
)

type handler struct {
	json         bool
	client       *pubsub.Client
	subscription *pubsub.Subscription
}

func (h *handler) receive() error {
	for {
		if err := h.subscription.Receive(context.TODO(), h.receiveFunc); err != nil {
			return err
		}
	}
}

func (h *handler) receiveFunc(ctx context.Context, msg *pubsub.Message) {
	msg.Ack()

	event, ok := msg.Attributes["event"]
	if !ok {
		log.Printf("event attribute is missing")
	}

	var err error

	switch event {
	case "up":
		err = h.up(msg.Data)
	case "join":
		err = h.join(msg.Data)
	default:
		log.Printf("handler for event %s is not implemented", event)
		return
	}

	if err != nil {
		log.Printf("handling event '%s' returned error: %s", event, err)
	}
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

func newHandler(json bool, projectID, subscriptionName, credentialsFile string) (*handler, error) {
	var opts []option.ClientOption

	if credentialsFile != "" {
		opts = append(opts, option.WithCredentialsFile(credentialsFile))
	}

	client, err := pubsub.NewClient(context.TODO(), projectID, opts...)
	if err != nil {
		return nil, err
	}

	subscription := client.Subscription(subscriptionName)
	exists, err := subscription.Exists(context.TODO())
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("subscription does not exist")
	}

	return &handler{
		json:         json,
		subscription: subscription,
		client:       client,
	}, nil
}

func main() {
	h, err := newHandler(
		false,               // set to true when using JSON encoding
		"project-id",        // replace with GCP project ID
		"subscription-name", // replace with GCP Pub/Sub Subscription name
		"",                  // path to GCP credentials file
	)
	if err != nil {
		panic(err)
	}

	if err := h.receive(); err != nil {
		panic(err)
	}
}
```

#### `go.mod`

```text
module example

go 1.14

require (
	cloud.google.com/go/pubsub v1.6.0
	github.com/brocaar/chirpstack-api/go/v3 v3.7.2
	github.com/golang/protobuf v1.4.2
	google.golang.org/api v0.29.0
)
```

### Python

Please refer to the [Setting up authentication](https://cloud.google.com/pubsub/docs/reference/libraries#client-libraries-install-python)
section for creating a service-account and setting up the credentials.

#### `main.py`

```python
from chirpstack_api.as_pb import integration
from google.cloud import pubsub_v1
from google.protobuf.json_format import Parse


class Handler:
    def __init__(self, json, project_id, subscription_name):
        self.json = json
        self.project_id = project_id
        self.subscription_name = subscription_name

    def receive(self):
        subscriber = pubsub_v1.SubscriberClient()
        sub_path = subscriber.subscription_path(self.project_id, self.subscription_name)

        while True:
            resp = subscriber.pull(sub_path, max_messages=10)
            for msg in resp.received_messages:
                event = msg.message.attributes['event']
                subscriber.acknowledge(sub_path, [msg.ack_id])

                if event == 'up':
                    self.up(msg.message.data)
                elif event == 'join':
                    self.join(msg.message.data)
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

    # GCP Project ID
    "project-id",

    # GCP Pub/Sub Subsciption name
    "subscription-name",
)
h.receive()
```

#### `requirements.txt`

```text
chirpstack-api==3.7.2
google-cloud-pubsub==1.7.0
```
