# Downlink payload format

<!-- toc -->

On downlink, the Border Gateway creates and re-transmits a Relay encapsulated
to the Relay Gateway in the following format. The Relay encapsulation overhead
is 15 bytes.

Bytes:

| 1 byte        | 6 bytes           | 4 bytes  | n bytes            | 4 bytes |
| ------------- | ----------------- | -------- | ------------------ | ------- |
| Downlink MHDR | Downlink Metadata | Relay ID | LoRaWAN PHYPayload | MIC     |

## Downlink MHDR

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

## Downlink Metadata

Bytes:

| 2 bytes                        | 3 bytes            | 1 byte           |
| ------------------------------ | ------------------ | ---------------- |
| Uplink ID + Downlink data-rate | Downlink Frequency | Delay + TX Power |

### Uplink ID + Downlink data-rate

Bits:

| 15..4     | 3..0               |
| --------- | ------------------ |
| Uplink ID | Downlink data-rate |

#### Uplink ID

Unique identifier (to the Relay Gateway) identifying the uplink to which this
downlink relates. E.g. the Relay Gateway uses this to retrieve the uplink
counter value.

#### Data-rate

Downlink data-rate, unsigned integer with a minimum value of `0` and a maximum
value of `15`.

### Downlink Frequency

Encoded as Frequency<sub>Hz</sub> = `Downlink Frequency x 100`.

### Delay + TX Power

Bits:

| 7..4     | 3..0  |
| -------- | ----- |
| TX Power | Delay |

#### TX Power

The TX Power which must be used for sending the downlink.

#### Delay

The relative delay in seconds to the Uplink ID which must be used
for the downlink transmission. Delay is an unsigned integer, encoded as:

Delay<sub>Seconds</sub> = `Delay + 1`.

## Relay ID

This contains the Relay ID that must relay the downlink to the
End Device using the provided Downlink metadata.

## LoRaWAN PHYPayload

The LoRaWAN PHYPayload that must be sent to the End-Device.

## MIC

Message integrity code, used by other Relay gateways to check the
data integrity of the packet. This is obtained by calculating the CMAC over
the downlink payload (- MIC bytes), and using the first 4 bytes of the calculated
CMAC as MIC.
