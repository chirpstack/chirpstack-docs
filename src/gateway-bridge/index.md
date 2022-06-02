# Introduction

ChirpStack Gateway Bridge is a service which converts LoRa<sup>&reg;</sup> Packet Forwarder protocols
into a ChirpStack [common data-format](https://github.com/chirpstack/chirpstack/blob/master/api/proto/gw/gw.proto).

## Backends

The following Packet Forwarder backends are provided:

* [ChirpStack Concentratord](../concentratord/index.md)
* [Semtech UDP Packet Forwarder](https://github.com/Lora-net/packet_forwarder)
* [Basic Station Packet Forwarder](https://github.com/lorabasics/basicstation)

## Integrations

The following integrations are provided:

* Generic MQTT broker
* [GCP Cloud IoT Core MQTT Bridge](https://cloud.google.com/iot-core/)
* [Azure IoT Hub MQTT Bridge](https://docs.microsoft.com/en-us/azure/iot-hub/iot-hub-mqtt-support)
