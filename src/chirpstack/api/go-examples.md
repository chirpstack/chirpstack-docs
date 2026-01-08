# Go examples

- [Go gRPC reference](https://pkg.go.dev/google.golang.org/grpc)
- [ChirpStack Go SDK reference](https://pkg.go.dev/github.com/chirpstack/chirpstack/api/go/v4)

## Enqueue downlink

The example below demonstrates:

- Configuration of gRPC _dial options_ including API token
- Connect to a gRPC API
- Define service client (in this case for the `DeviceService`)
- Perform an API call for a service (in this case `Enqueue`)

### `main.go`

```go
{{#include ../../../examples/chirpstack/api/go/enqueue_downlink.go}}
```
