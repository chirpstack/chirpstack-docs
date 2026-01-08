# HTTP

If configured, the HTTP integration will make `POST` requests to the configured
event endpoint or endpoints (multiple URLs can be configured, comma separated).
The `event` URL query parameter indicates the type of the event.

## Events

The HTTP integration exposes all events as documented by
[Event types](events.md).

## Example code

The following code examples are for demonstration purposes only to demonstrate
how integration events can be decoded it the most simple way without taking
performance or security in mind. Additional code might be required for
production usage.

### Go

The following code example demonstrates how to implement an HTTP endpoint using
[Go](https://golang.org/) which decodes either a Protobuf or JSON payload. If
you run this example on the same host as ChirpStack, then the endpoint for the
HTTP integration is `http://localhost:8090`.

`main.go`:

```go
{{#include ../../../examples/chirpstack/integrations/http/go/main.go}}
```

### Python

The following code example demonstrates how to implement an HTTP endpoint using
[Python 3](https://www.python.org/) which decodes either a Protobuf or JSON
payload. If you run this example on the same host as ChirpStack, then the
endpoint for the HTTP integration is `http://localhost:8090`.

`main.py`:

```python
{{#include ../../../examples/chirpstack/integrations/http/python/main.py}}
```
