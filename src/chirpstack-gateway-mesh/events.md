# Events

The ChirpStack Gateway Mesh provides the option to send events from the Relay
Gateways to the server (through the Border Gateway(s), ChirpStack MQTT Forwarder
to the MQTT broker).

These events can be divided into the following ranges:

* 0 - 127: Known / built-in events (encoding / decoding known)
* 128 - 255: Proprietary / custom events (passed as binary payloads)


### Known events

These events are known by the ChirpStack Gateway Mesh. An example is the
periodic heartbeat sent by Relay Gateways. See also the list of
[TLV payloads](./protocol/tlv-payloads.md).

### Proprietary events

These events can be freely configured in the [Configuration](./configuration.md)
file. Configuration is done in two steps:

* Configuration of type (number) to command
* Configuration of event-sets (list of type numbers and the interval)

## MQTT topic

Events are published to the following MQTT topic using the [MeshEvent](https://github.com/chirpstack/chirpstack/blob/master/api/proto/gw/gw.proto)
payload:

```
[PREFIX]/gateway/[GATEWAY_ID]/event/mesh
````

**Note:** Depending the ChirpStack MQTT Forwarder [Configuration](../chirpstack-mqtt-forwarder/configuration.md)
this payload will be Protobuf or JSON encoded.

## Example

**Note:** The following example assumes the ChirpStack MQTT Forwarder on the
Border Gateway is configured with `json=true`.

The following example demonstrates how to periodically send the output of the
`uptime` command as event-type `128`. Please in this example, for simplicity
the output of the command will be sent as text. To keep the payload as small
as possible, it is recommended to use a binary encoding.

```toml
[events.commands]
  128 = ["/usr/bin/uptime"]

[[events.sets]]
  interval = "5min"
  events = [128]
```

This will be exposed over MQTT using the `[PREFIX]/gateway/[GATEWAY_ID]/event/mesh`
topic. Example payload:

```json
{
   "gatewayId": "0016c001ff1a0a5f",
   "relayId": "f1137eb4",
   "time": "2025-05-18T10:28:25+00:00",
   "events": [
      {
         "proprietary": {
            "eventType": 128,
            "payload": "IDEwOjI4OjI1IHVwIDMwIG1pbiwgIGxvYWQgYXZlcmFnZTogMC4wNiwgMC4wNiwgMC4wNAo="
         }
      }
  ]
}
```

Notes:

* `gatewayId` is the ID of the Border Gateway
* `relayId` is the ID of the Relay Gateway
* `time` is the time when the event was sent by the Relay Gateway

Note that the `payload` is BASE64 encoded, as the event can be in binary form.
In this example, the output is text. BASE64 decoding `payload` gives:

```
 10:28:25 up 30 min,  load average: 0.06, 0.06, 0.04
```
