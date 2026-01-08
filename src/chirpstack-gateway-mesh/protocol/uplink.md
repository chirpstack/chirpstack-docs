# Uplink payload format

<!-- toc -->

On receiving an uplink, the Relay creates and re-transmits a Relay encapsulated
LoRaWAN payload. The Relay encapsulation overhead is 14 bytes.

Bytes:

| 1 byte      | 5 bytes         | 4 bytes  | n bytes            | 4 bytes |
| ----------- | --------------- | -------- | ------------------ | ------- |
| Uplink MHDR | Uplink Metadata | Relay ID | LoRaWAN PHYPayload | MIC     |

## Uplink MHDR

Bits:

| 7..5  | 4..3         | 2..0      |
| ----- | ------------ | --------- |
| MType | Payload type | Hop count |

- MType = `111` (= Proprietary LoRaWAN MType)
- Payload type = `00` (= Relayed uplink)
- Hop count = `000` = 1, ... `111` = 8

**Note:** The hop count is incremented each time the uplink payload is relayed
by an other Relay Gateway. As this changes the uplink payload, the MIC must be
re-calculated.

## Uplink Metadata

Bytes:

| 2 bytes               | 1 byte | 1 byte | 1 byte  |
| --------------------- | ------ | ------ | ------- |
| Uplink ID + Data-rate | RSSI   | SNR    | Channel |

### Uplink ID + Data-rate

Bits:

| 15..4     | 3..0      |
| --------- | --------- |
| Uplink ID | Data-rate |

#### Uplink ID

Unique identifier (to the Relay Gateway) identifying the received uplink.
Internally the Relay Gateway must temporarily store this in the Uplink table.

#### Data-rate

Uplink data-rate, unsigned integer with a minimum value of `0` and a maximum
value of `15`.

### RSSI

Encoded as RSSI<sub>dBm</sub> = `-1 x RSSI`

### SNR

Bits:

| 7..6 | 5..0 |
| ---- | ---- |
| RFU  | SNR  |

SNR is a signed integer with a minimum value of `-32` and a maximum value of
`31`.

### Channel

Uplink channel, unsigned integer.

## Relay ID

This contains the Relay ID which received the uplink from the End Device.

Bytes:

| 4 bytes  |
| -------- |
| Relay ID |

## LoRaWAN PHYPayload

The received LoRaWAN PHYPayload.

## MIC

Message integrity code, used by other Relay and Border gateways to check the
data integrity of the packet. This is obtained by calculating the CMAC over the
uplink payload (- MIC bytes), and using the first 4 bytes of the calculated CMAC
as MIC.
