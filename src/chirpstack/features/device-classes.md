# Device classes

## Class-A

### Uplink

ChirpStack supports LoRaWAN<sup>&reg;</sup> Class-A devices. In Class-A a device is always in sleep
mode, unless it has something to transmit. Only after an uplink transmission
by the device, ChirpStack is able to schedule a downlink transmission.

Received frames are de-duplicated (in case it has been received by multiple
gateways), after which the mac-layer is handled by ChirpStack and the
decrypted application-payload is forwarded to the configured integrations.

### Downlink

ChirpStack persists downlink queue-items in its database. One a receive window
occurs (triggered by an uplink), ChirpStack will transmit the items in the
downlink queue FIFO to the device. Please note that mac-commands will get
priority over application-payloads.

#### Confirmed data

ChirpStack sends an ack event with `ack: true` to the configured integration
as soon the device has acknowledged the confirmed uplink. If the next uplink
does not contain an acknowledgement, an ack event with `ack: false` is sent
to the configured integrations.

**Note:** Always check the value of the `ack` field in the ack event! For
some integrations default values are omitted, which is equal to `ack: false`!

## Class-B

ChirpStack supports Class-B devices. A Class-B device synchronizes its
internal clock using Class-B beacons emitted by the gateway, this process
is also called a "beacon lock". Once in the state of a beacon lock, the
device negotiates its ping-interval. ChirpStack is then able to schedule
downlink transmissions on each occurring ping-interval. 

### Confirmed downlink

The ack event works the same as for Class-A devices, with the only exception
that an `ack: false` is triggered after the configured Class-B timeout. This
timeout defines the maximum time that the device has to acknowledge a confirmed
downlink. This timeout can be configured in the device-profile.

### Requirements

#### Device

The device must be able to operate in Class-B mode.

#### Gateway

The gateway must have a GNSS based time-source, as Class-B beacon and ping-slot
scheduling must be very precise. The packet-forwarder must be configured to
send Class-B beacons such that the device is able to synchronize its internal
clock.

## Class-C

### Downlink

ChirpStack supports Class-C devices and uses the same Class-A
downlink device-queue for Class-C downlink transmissions. 

### Confirmed data

The ack event works the same as for Class-A devices, with the only exception
that an `ack: false` is triggered after the configured Class-C timeout. This
timeout defines the maximum time that the device has to acknowledge a confirmed
downlink. This timeout can be configured in the device-profile.
