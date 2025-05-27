# TLV payloads

<!-- toc -->

| Type        | Event / Command | Length   | Name            |
| ----------- | --------------- | -------- | --------------- |
| 0x00        | Event           | Variable | Heartbeat       |
| 0x80 - 0xff | Event / Command | Proprietary payloads       |

## Heartbeat

Periodically, each Relay Gateway will send a heartbeat message to indicate that
that it is still online. Each Relay Gateway relaying this heartbeat payload will
add its own Relay ID to the path, such that after the Border Gateway has
receive the payload, the full path from sending Relay Gateway to Border Gateway
can be obtained.

As the payload size is variable and will increment each time it is relayed, caution
must be taken when mixing this event with other events.

Bytes:

| 0 - 28 bytes          |
| --------------------- |
| Relay path (repeated) |

### Relay path

Bytes (repeated):

| 4 bytes  | 1 byte | 1 byte |
| -------- | ------ | ------ |
| Relay ID | RSSI   | SNR    |

Each Relay Gateway relaying the heartbeat payload adds its Relay ID, and the
RSSI and SNR of the received heartbeat payload to the end of the Relay path field.
Initially this field is empty. If the heartbeat payload is not relayed by any
other Relay Gateway, then this field remains empty.

#### RSSI

Encoded as RSSI<sub>dBm</sub> = `-1 x RSSI`

#### SNR

Bits:

| 7..6 | 5..0 |
| ---- | ---- |
| RFU  | SNR  |

SNR is a signed integer with a minimum value of `-32` and a maximum value of
`31`.
