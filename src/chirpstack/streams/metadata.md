# Uplink / downlink metadata

The uplink / downlink metadata stream exposes metadata that can be used for for
example billing-purposes.

## Redis key

This stream is published under the `stream:meta` Redis key.

## Data

This stream exposes:

- DevEUI
- Message-type
- RxInfo / TxInfo
- PHYPayload size
- MAC-command size
- Application payload size

See also
[meta.proto](https://github.com/chirpstack/chirpstack/blob/master/api/proto/stream/meta.proto).

## Example code

Example code can be found in the
[examples](https://github.com/chirpstack/chirpstack/tree/master/examples)
directory of the
[https://github.com/chirpstack/chirpstack](https://github.com/chirpstack/chirpstack)
repository.
