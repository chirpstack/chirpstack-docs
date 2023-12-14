# Protocol

This page describes the Relay encapsulated payload that is used for relaying
uplink and downlink transmissions.

<!-- toc -->

## Stored context

This section describes which context must be known at the different components
in the Relay architecture. Unless otherwise specified, this context must be
known at both the Relay Gateway and Border Gateway.

### Relay Gateway ID length

Indicates the number of Gateway ID bytes that are included in the Relay
Gateway ID field. Default `1`. The Relay Gateway might use the full 8 byte
Gateway ID and uses the N least significant bytes, or it might make the
Relay Gateway ID configurable.

### Data-rate table

A data-rate table, which maps the data-rate parameters to an integer value
(and back). There is no requirement that this follows the data-rates as
specified in RP002.

### Uplink channel index table

An uplink channel table, which maps an uplink channel to an integer value
(and back). There is no requirement that this follows the channels as
specified in RP002.

### Uplink table

The Relay Gateway stores a list of Uplink IDs (integer value) mapped to data
required to transmit a downlink. Typically the Uplink ID maps to the
concentrator counter value of the received uplink. On Class-A downlink to the
same device, the Relay Gateway will receive the Uplink ID + a relative delay
to the Uplink ID.

## Uplink payload format

On receiving an uplink, the Relay creates and re-transmits a Relay encapsulated
LoRaWAN payload. The minimum Relay encapsulation overhead is 7 bytes (assuming
a Relay Gateway ID of 1 byte).

Bytes:

| 1 byte      | 5 bytes         | n bytes          | n bytes            |
| ----------- | --------------- | ---------------- | ------------------ |
| Uplink MHDR | Uplink Metadata | Relay Gateway ID | LoRaWAN PHYPayload |


### Uplink MHDR

Bits:

| 7..5  | 4         | 3..2      | 1..0    |
| ----- | --------- | --------- | ------- |
| MType | Direction | Hop count | Version |

* MType = `111` (= Proprietary LoRaWAN MType)
* Direction = `0` (uplink)
* Hop count = `00`
* Version = `00`

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

### Relay Gateway ID

This contains the Relay Gateway ID, using the number of bytes as configured.
E.g. for a small network, 1 byte might be sufficient to uniquely identify all
the Relay Gateways. For larger networks, it might be needed to increase this
number.

Bytes:

| n bytes          |
| ---------------- |
| Relay Gateway ID |


## Downlink payload format

On downlink, the Border Gateway creates and re-transmits a Relay encapsulated
to the Relay Gateway in the following format. The minimum Relay encapsulation
overhead is 7 bytes (assuming a Relay Gateway ID of 1 byte).

Bytes:

| 1 byte        | 5 bytes           | n bytes          | n bytes            |
| ------------- | ----------------- | ---------------- | ------------------ |
| Downlink MHDR | Downlink Metadata | Relay Gateway ID | LoRaWAN PHYPayload |

### Downlink MHDR

Bits:

| 7..5  | 4         | 3..2      | 1..0    |
| ----- | --------- | --------- | ------- |
| MType | Direction | Hop count | Version |

* MType = `111` (= Proprietary LoRaWAN MType)
* Direction = `1` (downlink)
* Hop count = `00`
* Version = `00`

### Downlink Metadata

Bytes:

| 2 bytes                        | 3 bytes            | 1 byte |
| ------------------------------ | ------------------ | ------ |
| Uplink ID + Downlink data-rate | Downlink Frequency | Delay  |

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

#### Delay

This holds the relative delay in seconds to the Uplink ID which must be used
for the downlink transmission:

Bits:

| 15..4 | 3..0  |
| ----- | ----- |
| RFU   | Delay |

Delay is an unsigned integer, encoded as:

Delay<sub>Seconds</sub> = `Delay + 1`.

### Relay Gateway ID

This contains the Relay Gateway ID that must relay the downlink to the
End Device using the provided Downlink metadata.
