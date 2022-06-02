# Geolocation

ChirpStack supports various forms of geolocation through the [LoRa Cloud](../integrations/loracloud.md)
integration. As ChirpStack does expose meta-data like gateway location, RSSI, SNR
and fine-timestamp (if supported by the gateway) in every uplink event, it is
also possible to implement your own geolocation resolver.

## Requirements

For TDOA (time difference of arrival) based geolocation, you must have a gateway
that has support for the fine-timestamp. If applicable, the fine-timestamp must
be encrypted by the packet-forwarder.
