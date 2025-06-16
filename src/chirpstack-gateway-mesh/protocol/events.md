# Relay event payload

<!-- toc -->

The Relay event payload allows a Relay Gateway to broadcast data like a heartbeat,
location information, battery status, ... The event payloads are encoded as [TLV](https://en.wikipedia.org/wiki/Type%E2%80%93length%E2%80%93value)
such that the event payload can contain data for known and unknown (proprietary)
event types.

Bytes:

| 1 byte     | 4 bytes   | 4 bytes  | n bytes                 | 4 bytes |
| ---------- | --------- | -------- | ----------------------- | ------- |
| Event MHDR | Timestamp | Relay ID | TLV payload (encrypted) | MIC     |


## Event MHDR

Bits:

| 7..5  | 4..3         | 2..0      |
| ----- | -------------| --------- |
| MType | Payload type | Hop count |

* MType = `111` (= Proprietary LoRaWAN MType)
* Payload type = `10` (= Relay event)
* Hop count = `000` = 1, ... `111` = 8

**Note:** The hop count is incremented each time the event payload is relayed
by an other Relay Gateway. The MIC must be re-calcalculated on each relay.

## Timestamp

Unix timestamp (seconds).

### Relay ID

The Relay ID of the Relay Gateway sending the event payload.

## TLV payload

Bytes:

| 1 byte | 1 byte | n bytes |
| ------ | ------ | ------- |
| Type   | Length | Payload |

Please see the [TLV payloads](./tlv-payloads.md) page for the known payload types.

### Encryption

The TLV payload is encrypted using the same encryption scheme as the
LoRaWAN FrmPayload (see section 4.3.3 of the LoRaWAN 1.0.4 specs). For the
encryption the block Ai is defined as follow:

Octets:

| 1    | 4        | 1    | 4        | 4         | 1    | 1 |
| ---- | -------- | ---- | -------- | --------- | ---- | - |
| 0x01 | 4 x 0x00 | 0x00 | Relay ID | Timestamp | 0x00 | i |


## MIC

Message integrity code, used by other Relay and Border gateways to check the
data integrity of the packet. This is obtained by calculating the CMAC over
the event payload (- MIC bytes), and using the first 4 bytes of the calculated
CMAC as MIC. **Note:** The MIC must be calculated after encrypting the TLV payload!
