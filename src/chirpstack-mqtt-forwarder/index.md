# Introduction

ChirpStack MQTT Forwarder is a MQTT packet forwarder for LoRa gateways. By
default it forwards packets in
[Protobuf](https://developers.google.com/protocol-buffers) binary format,
optionally it can be configured to use JSON encoding for debugging. In contrast
to the [ChirpStack Gateway Bridge](../chirpstack-gateway-bridge/index.md), this
component must always be installed on the gateway.

## Backends

The following backends are provided:

- [ChirpStack Concentratord](../chirpstack-concentratord/index.md)
- [Semtech UDP Packet Forwarder](https://github.com/Lora-net/packet_forwarder)

### ChirpStack Concentratord

```d2
vars: {
	d2-config: {
		layout-engine: elk
	}
}

lora_gateway: LoRa Gateway
vm: Cloud / VM

lora_gateway: {
	chirpstack_concentratord: ChirpStack Concentratord
	chirpstack_mqtt_forwarder: ChirpStack MQTT Forwarder

	chirpstack_concentratord <> chirpstack_mqtt_forwarder: ZeroMQ
}

vm: {
	mqtt_broker: MQTT Broker
}

lora_gateway.chirpstack_mqtt_forwarder <-> vm.mqtt_broker: MQTT
```

### Semtech UDP Packet Forwarder

```d2
vars: {
	d2-config: {
		layout-engine: elk
	}
}

lora_gateway: LoRa Gateway
vm: Cloud / VM

lora_gateway: {
	udp_forwarder: Semtech UDP Packet Forwarder
	chirpstack_mqtt_forwarder: ChirpStack MQTT Forwarder

	udp_forwarder <> chirpstack_mqtt_forwarder: UDP
}

vm: {
	mqtt_broker: MQTT Broker
}

lora_gateway.chirpstack_mqtt_forwarder <-> vm.mqtt_broker: MQTT
```
