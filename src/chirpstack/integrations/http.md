# HTTP

If configured, the HTTP integration will make `POST` requests to the
configured event endpoint or endpoints (multiple URLs can be configured, comma
separated). The `event` URL query parameter indicates the type of the event.

## Events

The HTTP integration exposes all events as documented by [Event types](events.md).

## Example code

The following code examples are for demonstration purposes only to
demonstrate how integration events can be decoded it the most simple way
without taking performance or security in mind. Additional code might be
required for production usage.

### Go

The following code example demonstrates how to implement an HTTP endpoint using
[Go](https://golang.org/) which decodes either a Protobuf or JSON payload. If
you run this example on the same host as ChirpStack, then
the endpoint for the HTTP integration is `http://localhost:8090`.

#### `main.go`

```go
package main

import (
	"bytes"
	"encoding/hex"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"

	"github.com/brocaar/chirpstack-api/go/v3/as/integration"
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

func main() {
	// json: false   - to handle Protobuf payloads (binary)
	// json: true    - to handle JSON payloads (Protobuf JSON mapping)
	http.Handle("/", &handler{json: false})
	log.Fatal(http.ListenAndServe(":8090", nil))
}
```

#### `go.mod`

```text
module example

go 1.14

require (
	github.com/brocaar/chirpstack-api/go/v3 v3.7.2
	github.com/golang/protobuf v1.3.5
)
```

### Python

The following code example demonstrates how to implement an HTTP endpoint using
[Python 3](https://www.python.org/) which decodes either a Protobuf or JSON
payload. If you run this example on the same host as ChirpStack,
then the endpoint for the HTTP integration is `http://localhost:8090`.

#### `main.py`

```python
from http.server import HTTPServer, BaseHTTPRequestHandler
from urllib.parse import urlparse, parse_qs

from chirpstack_api.as_pb import integration
from google.protobuf.json_format import Parse


class Handler(BaseHTTPRequestHandler):
    # True -  JSON marshaler
    # False - Protobuf marshaler (binary)
    json = False

    def do_POST(self):
        self.send_response(200)
        self.end_headers()
        query_args = parse_qs(urlparse(self.path).query)

        content_len = int(self.headers.get('Content-Length', 0))
        body = self.rfile.read(content_len)

        if query_args["event"][0] == "up":
            self.up(body)

        elif query_args["event"][0] == "join":
            self.join(body)

        else:
            print("handler for event %s is not implemented" % query_args["event"][0])

    def up(self, body):
        up = self.unmarshal(body, integration.UplinkEvent())
        print("Uplink received from: %s with payload: %s" % (up.dev_eui.hex(), up.data.hex()))

    def join(self, body):
        join = self.unmarshal(body, integration.JoinEvent())
        print("Device: %s joined with DevAddr: %s" % (join.dev_eui.hex(), join.dev_addr.hex()))

    def unmarshal(self, body, pl):
        if self.json:
            return Parse(body, pl)
        
        pl.ParseFromString(body)
        return pl

httpd = HTTPServer(('', 8090), Handler)
httpd.serve_forever()
```

#### `requirements.txt`

```text
chirpstack-api==3.7.2
```
