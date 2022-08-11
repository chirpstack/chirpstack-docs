package main

import (
	"encoding/hex"
	"io/ioutil"
	"log"
	"net/http"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	"github.com/chirpstack/chirpstack/api/go/v4/integration"
)

type handler struct {
	json bool
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	event := r.URL.Query().Get("event")

	switch event {
	case "up":
		err = h.up(b)
	case "join":
		err = h.join(b)
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

func main() {
	// json: false   - to handle Protobuf payloads (binary)
	// json: true    - to handle JSON payloads (Protobuf JSON mapping)
	http.Handle("/", &handler{json: false})
	log.Fatal(http.ListenAndServe(":8090", nil))
}
