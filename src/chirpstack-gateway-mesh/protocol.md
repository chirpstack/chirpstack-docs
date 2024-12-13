# Protocol

This page describes the Relay encapsulated payload that is used for relaying
uplink and downlink transmissions.

<!-- toc -->

## Stored context

This section describes which context must be known at the different components
in the Mesh architecture. Unless otherwise specified, this context must be
known at both the Relay Gateway and Border Gateway.

### Signing key

An AES128 key which is used to add a message integrity code to each mesh
packet. Each Relay and Border gateway must be configured with this same key
in order to sign and validate all mesh packets.

### TX Power table

A TX Power table, which maps the (downlink) TX Power to an integer (0 - 15)
and back. Ideally this matches the TX Power table of the Concentratord
configuration (if there is no exact match, the Concentratord will use the
highest value, that is less than the requested TX Power).

### Data-rate table

A data-rate table, which maps the data-rate parameters to an integer value
(and back). There is no requirement that this follows the data-rate indices
as specified in RP002.

### Uplink channel index table

An uplink channel table, which maps an uplink channel to an integer value
(and back). There is no requirement that this follows the channel indices as
specified in RP002.

### Uplink table

The Relay Gateway stores a list of Uplink IDs (integer value) mapped to data
required to transmit a downlink. Typically the Uplink ID maps to the
concentrator counter value of the received uplink. On Class-A downlink to the
same device, the Relay Gateway will receive the Uplink ID + a relative delay
to the Uplink ID.

## Uplink payload format

On receiving an uplink, the Relay creates and re-transmits a Relay encapsulated
LoRaWAN payload. The Relay encapsulation overhead is 14 bytes.

Bytes:

| 1 byte      | 5 bytes         | 4 bytes  | n bytes            | 4 bytes |
| ----------- | --------------- | -------- | ------------------ | ------- |
| Uplink MHDR | Uplink Metadata | Relay ID | LoRaWAN PHYPayload | MIC     |


### Uplink MHDR

Bits:

| 7..5  | 4..3         | 2..0      |
| ----- | -------------| --------- |
| MType | Payload type | Hop count |

* MType = `111` (= Proprietary LoRaWAN MType)
* Payload type = `00` (= Relayed uplink)
* Hop count = `000` = 1, ... `111` = 8

**Note:** The hop count is incremented each time the uplink payload is relayed
by an other Relay Gateway. As this changes the uplink payload, the MIC must be
re-calculated.

### Uplink Metadata

Bytes:

| 2 bytes               | 1 byte | 1 byte | 1 byte  |
| --------------------- | ------ | ------ | ------- |
| Uplink ID + Data-rate | RSSI   | SNR    | Channel |

#### Uplink ID + Data-rate

Bits:

| 15..4     | 3..0      |
| --------- | --------- |
| Uplink ID | Data-rate |

##### Uplink ID

Unique identifier (to the Relay Gateway) identifying the received uplink.
Internally the Relay Gateway must temporarily store this in the Uplink table.

##### Data-rate

Uplink data-rate, unsigned integer with a minimum value of `0` and a maximum
value of `15`.

#### RSSI

Encoded as RSSI<sub>dBm</sub> = `-1 x RSSI`

#### SNR

Bits:

| 7..6 | 5..0 |
| ---- | ---- |
| RFU  | SNR  |

SNR is a signed integer with a minimum value of `-32` and a maximum value of
`31`.

#### Channel

Uplink channel, unsigned integer.

### Relay ID

This contains the Relay ID which received the uplink from the End Device.

Bytes:

| 4 bytes  |
| -------- |
| Relay ID |

### LoRaWAN PHYPayload

The received LoRaWAN PHYPayload.

### MIC

Message integrity code, used by other Relay and Border gateways to check the
data integrity of the packet. This is obtained by calculating the CMAC over
the uplink payload (- MIC bytes), and using the first 4 bytes of the calculated
CMAC as MIC.


## Downlink payload format

On downlink, the Border Gateway creates and re-transmits a Relay encapsulated
to the Relay Gateway in the following format. The Relay encapsulation overhead
is 15 bytes.

Bytes:

| 1 byte        | 6 bytes           | 4 bytes  | n bytes            | 4 bytes |
| ------------- | ----------------- | -------- | ------------------ | ------- |
| Downlink MHDR | Downlink Metadata | Relay ID | LoRaWAN PHYPayload | MIC     |

### Downlink MHDR

Bits:

| 7..5  | 4..3         | 2..0      |
| ----- | -------------| --------- |
| MType | Payload type | Hop count |

* MType = `111` (= Proprietary LoRaWAN MType)
* Payload type = `01` (Relayed downlink)
* Hop count = `000`

**Note:** The hop count is incremented each time the downlink payload is relayed
by an other Relay Gateway. As this changes the downlink payload, the MIC must be
re-calculated.

### Downlink Metadata

Bytes:

| 2 bytes                        | 3 bytes            | 1 byte           |
| ------------------------------ | ------------------ | ---------------- |
| Uplink ID + Downlink data-rate | Downlink Frequency | Delay + TX Power |

#### Uplink ID + Downlink data-rate

Bits:

| 15..4     | 3..0               |
| --------- | ------------------ |
| Uplink ID | Downlink data-rate |

##### Uplink ID

Unique identifier (to the Relay Gateway) identifying the uplink to which this
downlink relates. E.g. the Relay Gateway uses this to retrieve the uplink
counter value.

##### Data-rate

Downlink data-rate, unsigned integer with a minimum value of `0` and a maximum
value of `15`.

#### Downlink Frequency

Encoded as Frequency<sub>Hz</sub> = `Downlink Frequency x 100`.

#### Delay + TX Power

Bits:

| 7..4     | 3..0  |
| -------- | ----- |
| TX Power | Delay |

##### TX Power

The TX Power which must be used for sending the downlink.

##### Delay

The relative delay in seconds to the Uplink ID which must be used
for the downlink transmission. Delay is an unsigned integer, encoded as:

Delay<sub>Seconds</sub> = `Delay + 1`.

### Relay ID

This contains the Relay ID that must relay the downlink to the
End Device using the provided Downlink metadata.

### LoRaWAN PHYPayload

The LoRaWAN PHYPayload that must be sent to the End-Device.

### MIC

Message integrity code, used by other Relay gateways to check the
data integrity of the packet. This is obtained by calculating the CMAC over
the downlink payload (- MIC bytes), and using the first 4 bytes of the calculated
CMAC as MIC.

## Relay event payload

The Relay event payload allows a Relay Gateway to broadcast data like a heartbeat,
location information, battery status, ... The information is encoded as [TLV](https://en.wikipedia.org/wiki/Type%E2%80%93length%E2%80%93value)
such that the event payload can contain data for known and unknown (proprietary)
event types.

Bytes:

| 1 byte     | 4 bytes   | n bytes     | 4 bytes |
| ---------- | --------- | ----------- | ------- |
| Event MHDR | Timestamp | TLV payload | MIC     |


### Event MHDR

Bits:

| 7..5  | 4..3         | 2..0      |
| ----- | -------------| --------- |
| MType | Payload type | Hop count |

* MType = `111` (= Proprietary LoRaWAN MType)
* Payload type = `10` (= Relay event)
* Hop count = `000` = 1, ... `111` = 8

**Note:** The hop count is incremented each time the event payload is relayed
by an other Relay Gateway. In the case this changes the payload (e.g. in case
of a heartbeat), the MIC must be re-calculated.

### Timestamp

Unix timestamp (seconds).

### TLV payload

The Type and Length are both encoded as single bytes. Please see the TLV eventpayloads
payloads section for the known payload types.

## TLV event payloads

| Type | Length   | Name      | Mixed |
| ---- | -------- | --------- | ----- |
| 0x00 | Variable | Heartbeat | No    |
| 0x80 - 0xff | Proprietary event payloads |

### Heartbeat

Periodically, each Relay Gateway will send a heartbeat message to indicate that
that it is still online. Each Relay Gateway relaying this heartbeat payload will
add its own Relay ID to the path, such that after the Border Gateway has
receive the payload, the full path from sending Relay Gateway to Border Gateway
can be obtained.

As the payload size is variable and will increment each time it is relayed, this
event type must not be mixed with other TLV payloads.

Bytes:

| 4 bytes  | 0 - 28 bytes          |
| -------- | --------------------- |
| Relay ID | Relay path (repeated) |

#### Relay ID

The Relay ID of the Relay Gateway sending the heartbeat message.

#### Relay path

Bytes (repeated):

| 4 bytes  | 1 byte | 1 byte |
| -------- | ------ | ------ |
| Relay ID | RSSI   | SNR    |

Each Relay Gateway relaying the heartbeat payload adds its Relay ID, and the
RSSI and SNR of the received heartbeat payload to the end of the Relay path field.
Initially this field is empty. If the heartbeat payload is not relayed by any
other Relay Gateway, then this field remains empty.

##### RSSI

Encoded as RSSI<sub>dBm</sub> = `-1 x RSSI`

##### SNR

Bits:

| 7..6 | 5..0 |
| ---- | ---- |
| RFU  | SNR  |

SNR is a signed integer with a minimum value of `-32` and a maximum value of
`31`.
