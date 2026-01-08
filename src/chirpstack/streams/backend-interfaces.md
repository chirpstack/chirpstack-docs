# Backend Interfaces API

The Backend Interfaces API stream exposes the Backend Interfaces requests and
responses (both from and to an external server). This includes:

- Passive Roaming related API calls
- Join Server related API calls

## Redis key

This stream is published under the `backend_interfaces:stream:request` Redis
key.

## Data

This stream exposes:

- Sender / Receiver ID
- Timestamp
- Transaction ID
- Message-type
- Result code
- Request and response body
- Request error

See also
[backend_interfaces.proto](https://github.com/chirpstack/chirpstack/blob/master/api/proto/stream/backend_interfaces.proto).

## Example code

Example code can be found in the
[examples](https://github.com/chirpstack/chirpstack/tree/master/examples)
directory of the
[https://github.com/chirpstack/chirpstack](https://github.com/chirpstack/chirpstack)
repository.
