# LoRaWAN frames

The frames stream exposes the raw LoRaWAN<sup>&reg;</sup> payloads.
This stream can be used for monitoring purposes. **Note:** this stream
should never be used for building application-integrations. For this
you should always use the [Events stream](./events.md).

## Redis keys

This stream is published under the following Redis keys:

* Frames received / sent by the gateways:
    * `gw:[GATEWAY_ID]:stream:frame` (per Gateway ID stream)
    * `gw:stream:frame` (global stream)
* Authenticated frames:
    * `device:[DEV_EUI]:stream:frame` (per DevEUI stream)
    * `device:stream:frame` (global stream)

## Data

This stream exposes the [UplinkFrame](https://github.com/chirpstack/chirpstack/blob/master/api/proto/gw/gw.proto#L473)
and [DownlinkFrame](https://github.com/chirpstack/chirpstack/blob/master/api/proto/gw/gw.proto#L501)
messages.

## Example code

Example code can be found in the [examples](https://github.com/chirpstack/chirpstack/tree/master/examples)
directory of the [https://github.com/chirpstack/chirpstack](https://github.com/chirpstack/chirpstack)
repository.
