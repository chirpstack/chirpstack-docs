# Rejoin configuration

ChirpStack supports the handling and configuration of rejoin-requests for
LoRaWAN<sup>&reg;</sup> 1.1.x devices.
Using a rejoin-request, a device is able to request the network-server to
re-activate, without being 'disconnected' until receiving the activation.
In other words, when it doesn't receive a new activation, the device will
continue to use the old security context.

With the rejoin-request, it is possible to reset the device context
including all radio parameters (device address, frame counters, session-keys,
radio parameters, ...) restore a lost session context or rekey a device
(device address, session-keys and frame-counters).

## Rejoin interval

ChirpStack supports configuring devices so that they will send a rejoin-request
based on a time and / or frame-counter based interval.

For more details on rejoin-requests, please refer to the LoRaWAN specification.
