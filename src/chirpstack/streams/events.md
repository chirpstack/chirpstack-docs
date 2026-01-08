# Integration events

The events stream exposes [Integration Events](../integrations/events.md). This
stream can be used to build external integrations, for example using the
[chirpstack_integration](https://crates.io/crates/chirpstack-integration) crate.

## Redis keys

This stream is published under the following Redis keys:

- `device:[DEV_EUI]:stream:event` (per DevEUI stream)
- `device:stream:event` (global stream)

## Data

This stream exposes all [Integration Events](../integrations/events.md).

See also
[integration.proto](https://github.com/chirpstack/chirpstack/blob/master/api/proto/integration/integration.proto).

## Example code

Example code can be found in the
[examples](https://github.com/chirpstack/chirpstack/tree/master/examples)
directory of the
[https://github.com/chirpstack/chirpstack](https://github.com/chirpstack/chirpstack)
repository.
