# FUOTA

Firmware update over the air (sometimes called FUOTA) makes it possible to
push firmware updates to one or multiple devices, making use of multicast.
It is standardized by the following LoRa<sup>&reg;</sup> Alliance specifications:

* LoRaWAN<sup>&reg;</sup> Application Layer Clock Synchronization
* LoRaWAN<sup>&reg;</sup> Fragmented Data Block Transport
* LoRaWAN<sup>&reg;</sup> Remote Multicast Setup

It is important to note that the implementation of this feature by devices
is optional and therefore, unless your device explicitly states that it
implements FUOTA it is safe to assume it does not.

ChirpStack does not provide integrated FUOTA support, but it does provide a
component called [ChirpStack FUOTA Server](https://github.com/brocaar/chirpstack-fuota-server).
