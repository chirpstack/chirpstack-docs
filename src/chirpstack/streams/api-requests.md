# API requests

The API requests stream exposes meta-data about [API](../api/index.md) requests
being made by the web-interface and / or external API clients.

## Redis key

This stream is published under the `api:stream:request` Redis key.

## Data

This stream exposes:

* gRPC API service (e.g. `DeviceService`)
* gRPC API method (e.g. `Create`)
* Metadata (e.g. the identifier of the created / updated / deleted object)

See also [api_request.proto](https://github.com/chirpstack/chirpstack/blob/master/api/proto/stream/api_request.proto).

## Example code

Example code can be found in the [examples](https://github.com/chirpstack/chirpstack/tree/master/examples)
directory of the [https://github.com/chirpstack/chirpstack](https://github.com/chirpstack/chirpstack)
repository.