# JavaScript examples

* [gRPC documentation for Node](https://grpc.io/docs/languages/node/)

## Enqueue

The example below demonstrates:

* Configuration of gRPC _dial options_ including API token
* Connect to a gRPC API
* Define service client (in this case for the `DeviceService`)
* Perform an API call for a service (in this case `Enqueue`)

### `enqueue_downlink.js`

```js
{{#include ../../../examples/chirpstack/api/js/enqueue_downlink.js}}
```
