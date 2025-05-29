# Stored context

This section describes which context must be known at the different components
in the Mesh architecture. Unless otherwise specified, this context must be
known at both the Relay Gateway and Border Gateway.

## Mesh root-key

An AES128 key which is used to add a message integrity code to each mesh
packet. In case of events and commands packets, it is also used for
encryption / decryption.  Each Relay and Border gateway must be configured
with this same key.

### MIC key

The MIC key is generated as follow:

```
MIC key = aes128_encrypt(Mesh Root Key, 0x00 | pad16)
```

### Encryption key

The encryption key is generated as follow:

```
Encryption key = aes128_encrypt(Mesh Root Key, 0x01 | pad16)
```

## TX Power table

A TX Power table, which maps the (downlink) TX Power to an integer (0 - 15)
and back. Ideally this matches the TX Power table of the Concentratord
configuration (if there is no exact match, the Concentratord will use the
highest value, that is less than the requested TX Power).

## Data-rate table

A data-rate table, which maps the data-rate parameters to an integer value
(and back). There is no requirement that this follows the data-rate indices
as specified in RP002.

## Uplink channel index table

An uplink channel table, which maps an uplink channel to an integer value
(and back). There is no requirement that this follows the channel indices as
specified in RP002.

## Uplink table

The Relay Gateway stores a list of Uplink IDs (integer value) mapped to data
required to transmit a downlink. Typically the Uplink ID maps to the
concentrator counter value of the received uplink. On Class-A downlink to the
same device, the Relay Gateway will receive the Uplink ID + a relative delay
to the Uplink ID.
