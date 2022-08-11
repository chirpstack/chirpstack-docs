package main

import (
	"context"
	"encoding/hex"
	"errors"
	"log"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	"github.com/chirpstack/chirpstack/api/go/v4/integration"
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
	log.Printf("Uplink received from %s with payload: %s", up.GetDeviceInfo().DevEui, hex.EncodeToString(up.Data))
	return nil
}

func (h *handler) join(b []byte) error {
	var join integration.JoinEvent
	if err := h.unmarshal(b, &join); err != nil {
		return err
	}
	log.Printf("Device %s joined with DevAddr %s", join.GetDeviceInfo().DevEui, join.DevAddr)
	return nil
}

func (h *handler) unmarshal(b []byte, v proto.Message) error {
	if h.json {
		return protojson.UnmarshalOptions{
			DiscardUnknown: true,
			AllowPartial:   true,
		}.Unmarshal(b, v)
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
