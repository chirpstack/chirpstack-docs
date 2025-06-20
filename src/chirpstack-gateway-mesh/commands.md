# Commands

The ChirpStack Gateway Mesh provides the option to send commands to the Relay
Gateways (over MQTT).

These commands can be divided into the following ranges:

* 0 - 127: Known / built-in commands (encoding / decoding known)
* 128 - 255: Proprietary / custom commands (passed as binary payloads)


### Known commands

Currently there are no known commands defined in the [TLV payloads](./protocol/tlv-payloads.md).

### Proprietary commands

These commands can be freely configured in the [Configuration](./configuration.md)
file.

## MQTT topic

Commands are triggering by sending a [MeshCommand](https://github.com/chirpstack/chirpstack/blob/master/api/proto/gw/gw.proto)
payload to the following MQTT topic:

```
[PREFIX]/gateway/[GATEWAY_ID]/command/mesh
```

**Note:** Depending the ChirpStack MQTT Forwarder [Configuration](../chirpstack-mqtt-forwarder/configuration.md)
this payload must be Protobuf or JSON encoded.

## Example

**Note:** The following example assumes the ChirpStack MQTT Forwarder on the
Border Gateway is configured with `json=true`. The following example
demonstrates a command that once executed, writes the STDIN into a file
located at `/tmp/test.txt` after which it returns `OK`.

### Command script

The following script writes the STDIN to `/tmp/text.txt` and returns `OK`.

```bash
#!/bin/sh
cat > "/tmp/test.txt"
echo "OK"
```

### Configuration

```toml
[commands.commands]
  129 = ["/tmp/test.sh"]
```

### Execution

The above command can be triggered by publishing the following MQTT payload:

```json
{
   "gatewayId": "0016c001ff1a0a5f",
   "relayId": "f1137eb4",
   "commands": [
      {
         "proprietary": {
            "commandType": 129,
            "payload": "SGVsbG8gd29ybGQh"
         }
      }
   ]
}
```

Notes:

* `gatewayId` is the ID of the Border Gateway
* `relayId` is the ID of the Relay Gateway
* `payload` is the BASE64 encoded payload that is used as STDIN input

After execution and once it has been received and executed by the Relay Gateway
you should see a text-file appear under `/tmp/text.txt` containing `Hello world!`.
Shortly after, the following [Event](./events.md) payload will be published over MQTT:

```json
{
   "gatewayId": "0016c001ff1a0a5f",
   "relayId": "f1137eb4",
   "time": "2025-06-18T10:26:15+00:00",
   "events": [
      {
         "proprietary": {
            "eventType": 129,
            "payload": "T0sK"
         }
      }
  ]
}
```

Note that BASE64 decoding the `payload` will return `OK`, which is the output
of the command script.
