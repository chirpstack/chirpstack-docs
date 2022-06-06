# AWS SNS

The [Simple Notification Service (SNS)](https://aws.amazon.com/sns/) integration
publishes all the events to a SNS Topic to which other applications or AWS
services can subscribe for further processing.

## Events

The AWS Simple Notification Service integration exposes all events as
documented by [Event types](events.md).

### Message attributes

The following message attributes are added to each published message:

* `event` - the event type
* `dev_eui` - the device EUI
* `application_id` - the ChirpStack Application Server application ID

## Example code

The following code example demonstrates how to consume integration events using
an [AWS SQS](https://aws.amazon.com/sqs/) subscription.

!!! important
	Make sure the _Enable raw message delivery_ option is enabled on the subscription.
	If not enabled, the SQS messages will not have the expected attributes.

### Go

#### `main.go`

```go
package main

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"

	"github.com/brocaar/chirpstack-api/go/v3/as/integration"
)

type handler struct {
	json bool

	sqs      *sqs.SQS
	queueURL string
}

func (h *handler) receive() error {
	for {
		result, err := h.sqs.ReceiveMessage(&sqs.ReceiveMessageInput{
			MessageAttributeNames: []*string{
				aws.String(sqs.QueueAttributeNameAll),
			},
			QueueUrl:            &h.queueURL,
			MaxNumberOfMessages: aws.Int64(1),
		})
		if err != nil {
			return err
		}

		for _, msg := range result.Messages {
			_, err := h.sqs.DeleteMessage(&sqs.DeleteMessageInput{
				QueueUrl:      &h.queueURL,
				ReceiptHandle: msg.ReceiptHandle,
			})
			if err != nil {
				log.Printf("delete message error: %s", err)
			}

			event, ok := msg.MessageAttributes["event"]
			if !ok || event.StringValue == nil {
				log.Printf("event attribute is missing")
				continue
			}

			switch *event.StringValue {
			case "up":
				err = h.up(*msg.Body)
			case "join":
				err = h.join(*msg.Body)
			default:
				log.Printf("handler for event %s is not implemented", *event.StringValue)
				err = nil
			}

			if err != nil {
				log.Printf("handling event '%s' returned error: %s", *event.StringValue, err)
			}

		}
	}

	return nil
}

func (h *handler) up(body string) error {
	var up integration.UplinkEvent
	if err := h.unmarshal(body, &up); err != nil {
		return err
	}
	log.Printf("Uplink received from %s with payload: %s", hex.EncodeToString(up.DevEui), hex.EncodeToString(up.Data))
	return nil
}

func (h *handler) join(body string) error {
	var join integration.JoinEvent
	if err := h.unmarshal(body, &join); err != nil {
		return err
	}
	log.Printf("Device %s joined with DevAddr %s", hex.EncodeToString(join.DevEui), hex.EncodeToString(join.DevAddr))
	return nil
}

func (h *handler) unmarshal(body string, v proto.Message) error {
	if h.json {
		unmarshaler := &jsonpb.Unmarshaler{
			AllowUnknownFields: true, // we don't want to fail on unknown fields
		}
		return unmarshaler.Unmarshal(bytes.NewReader([]byte(body)), v)
	}

	b, err := base64.StdEncoding.DecodeString(body)
	if err != nil {
		return err
	}

	return proto.Unmarshal(b, v)
}

func newHandler(json bool, accessKeyID, secretAccessKey, region, queueURL string) (*handler, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKeyID, secretAccessKey, ""),
	})
	if err != nil {
		return nil, err
	}

	return &handler{
		json:     json,
		sqs:      sqs.New(sess),
		queueURL: queueURL,
	}, nil
}

func main() {
	h, err := newHandler(
		// set true when using JSON encoding
		false,

		// AWS AccessKeyID
		"...",

		// AWS SecretAccessKey
		"...",

		// AWS region
		"eu-west-1",

		// SQS queue url
		"https://sqs...",
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
	github.com/aws/aws-sdk-go v1.33.12
	github.com/brocaar/chirpstack-api/go/v3 v3.7.2
	github.com/golang/protobuf v1.4.2
)
```

### Python

Please refer to the [Boto3 Configuration](https://boto3.amazonaws.com/v1/documentation/api/latest/guide/quickstart.html)
for setting up the API credentials.

#### `main.py`

```python
import boto3
from chirpstack_api.as_pb import integration
from google.protobuf.json_format import Parse


class Handler:
    def __init__(self, json, queue_url):
        self.json = json
        self.queue_url = queue_url

    def receive(self):
        sqs = boto3.client('sqs')

        while True:
            resp = sqs.receive_message(
                QueueUrl=self.queue_url,
                MessageAttributeNames=[
                    'All',
                ],
                MaxNumberOfMessages=1,
                WaitTimeSeconds=10,
            )

            if not 'Messages' in resp:
                continue

            msg = resp['Messages'][0]
            receipt_handle = msg['ReceiptHandle']

            sqs.delete_message(
                QueueUrl=self.queue_url,
                ReceiptHandle=receipt_handle,
            )

            event = msg['MessageAttributes']['event']['StringValue']

            if event == "up":
                self.up(msg['Body'])
            elif event == "join":
                self.join(msg['Body'])
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

    # SQS queue url
    'https://sqs....',
)
h.receive()
```

#### `requirements.txt`

```text
boto3==1.14.29
chirpstack-api==3.7.2
```
