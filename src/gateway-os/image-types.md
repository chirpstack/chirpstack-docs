# Image types

## Base

The **chirpstack-gateway-os-base** image provides all the features and components
to use a Raspberry Pi as LoRa gateway, forwarding data to an external server.

- ChirpStack Concentratord
- ChirpStack Gateway Bridge
- Gateway configuration utility (`gateway-config`)

## Full

The **chirpstack-gateway-os-full** provides all the features of the
**chirpstack-gateway-os-base** image, but is also bundled with the ChirpStack
Network Server and all requirements to run a complete LoRaWAN stack on the
Raspberry Pi. As all components are pre-configured, this is ideal when getting
started with LoRaWAN.
