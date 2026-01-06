# Architecture

<!-- toc -->

A typical ChirpStack deployment has the following architecture. Note that in
this diagram, the gateways are connected in different ways to ChirpStack to
highlight the different possible connectivity options.

```d2
vars: {
	d2-config: {
		layout-engine: elk
	}
}

vm: Cloud / VM
lora_gateway_udp: LoRa Gateway
lora_gateway_bs: LoRa Gateway
lora_gateway_mqtt: LoRa Gateway
integrations: Integrations
api_client: API Client

vm: {
	chirpstack: ChirpStack
	chirpstack_gateway_bridge: ChirpStack Gateway Bridge
	mqtt_broker: MQTT Broker

	chirpstack <> mqtt_broker
	chirpstack_gateway_bridge <> mqtt_broker

}

lora_gateway_udp: {
	udp_packet_forwarder: UDP Packet Forwarder
}

lora_gateway_bs: {
	basics_station: Basics Station
}

lora_gateway_mqtt: {
	chirpstack_mqtt_forwarder: ChirpStack MQTT Forwarder
}

integrations: {
	http: HTTP
	mqtt: MQTT
	etc: Etc..
}

lora_gateway_udp.udp_packet_forwarder <> vm.chirpstack_gateway_bridge: UDP
lora_gateway_bs.basics_station <> vm.chirpstack_gateway_bridge: Websockets
lora_gateway_mqtt.chirpstack_mqtt_forwarder <> vm.mqtt_broker: MQTT
vm.chirpstack -> integrations.http
vm.chirpstack -> integrations.mqtt
vm.chirpstack -> integrations.etc
api_client -> vm.chirpstack: gRPC
```

## Multi-region example 1

This example displays a multi-region ChirpStack installation (EU868 and US915)
using a single MQTT broker. Please note the MQTT prefix. In this case for US915
the prefix `us915_0` is used, indicating that this region uses the first 8
channels of the US915 region (the prefix is user-configurable). To keep the
diagram simple, some elements of the previous diagram have been left out.

```d2
vars: {
	d2-config: {
		layout-engine: elk
	}
}

vm: Cloud / VM
lora_gateway_udp_eu868: LoRa Gateway (EU868)
lora_gateway_udp_us915: LoRa Gateway (US915)
lora_gateway_mqtt_eu868: LoRa Gateway (EU868)
lora_gateway_mqtt_us915: LoRa Gateway (US915)

vm: {
	chirpstack_gateway_bridge_eu868: ChirpStack Gateway Bridge (EU868)
	chirpstack_gateway_bridge_us915: ChirpStack Gateway Bridge (US915)
	chirpstack: ChirpStack
	mqtt_broker: MQTT Broker

	chirpstack_gateway_bridge_eu868 <> mqtt_broker: MQTT\neu868/gateway/...
	chirpstack_gateway_bridge_us915 <> mqtt_broker: MQTT\nus915_0/gateway/...
	chirpstack <> mqtt_broker: MQTT
	chirpstack <> mqtt_broker: MQTT
	
}

lora_gateway_udp_eu868: {
	udp_packet_forwarder: UDP Packet Forwarder
}

lora_gateway_udp_us915: {
	udp_packet_forwarder: UDP Packet Forwarder
}

lora_gateway_mqtt_eu868: {
	chirpstack_mqtt_forwarder: ChirpStack MQTT Forwarder
}

lora_gateway_mqtt_us915: {
	chirpstack_mqtt_forwarder: ChirpStack MQTT Forwarder
}

lora_gateway_udp_eu868.udp_packet_forwarder <> vm.chirpstack_gateway_bridge_eu868: UDP
lora_gateway_udp_us915.udp_packet_forwarder <> vm.chirpstack_gateway_bridge_us915: UDP
lora_gateway_mqtt_eu868.chirpstack_mqtt_forwarder <> vm.mqtt_broker: MQTT\neu868/gateway/...
lora_gateway_mqtt_us915.chirpstack_mqtt_forwarder <> vm.mqtt_broker: MQTT\nus915_0/gateway/...
```


## Multi-region example 2

As an alternative to the previous diagram, it is also possible to setup a
separate MQTT broker per region. In this it is no longer a requirement to use
the region prefix in the MQTT topic.

```d2
vars: {
	d2-config: {
		layout-engine: elk
	}
}

vm: Cloud / VM
lora_gateway_udp_eu868: LoRa Gateway (EU868)
lora_gateway_udp_us915: LoRa Gateway (US915)
lora_gateway_mqtt_eu868: LoRa Gateway (EU868)
lora_gateway_mqtt_us915: LoRa Gateway (US915)

vm: {
	chirpstack_gateway_bridge_eu868: ChirpStack Gateway Bridge (EU868)
	chirpstack: ChirpStack
	chirpstack_gateway_bridge_us915: ChirpStack Gateway Bridge (US915)
	mqtt_broker_eu868: MQTT Broker EU868
	mqtt_broker_us915: MQTT Broker US915

	chirpstack_gateway_bridge_eu868 <> mqtt_broker_eu868: MQTT
	chirpstack_gateway_bridge_us915 <> mqtt_broker_us915: MQTT
	chirpstack <> mqtt_broker_eu868: MQTT
	chirpstack <> mqtt_broker_us915: MQTT
	
}

lora_gateway_udp_eu868: {
	udp_packet_forwarder: UDP Packet Forwarder
}

lora_gateway_udp_us915: {
	udp_packet_forwarder: UDP Packet Forwarder
}

lora_gateway_mqtt_eu868: {
	chirpstack_mqtt_forwarder: ChirpStack MQTT Forwarder
}

lora_gateway_mqtt_us915: {
	chirpstack_mqtt_forwarder: ChirpStack MQTT Forwarder
}

lora_gateway_udp_eu868.udp_packet_forwarder <> vm.chirpstack_gateway_bridge_eu868: UDP
lora_gateway_udp_us915.udp_packet_forwarder <> vm.chirpstack_gateway_bridge_us915: UDP
lora_gateway_mqtt_eu868.chirpstack_mqtt_forwarder <> vm.mqtt_broker_eu868: MQTT
lora_gateway_mqtt_us915.chirpstack_mqtt_forwarder <> vm.mqtt_broker_us915: MQTT
```

## ChirpStack components

### ChirpStack Concentratord

[ChirpStack Concentratord](./chirpstack-concentratord/index.md) is an
open-source LoRa concentrator daemon. It exposes a [ZeroMQ](https://zeromq.org/)
based API that can be used by one or multiple (forwarder) applications to
interact with the gateway hardware.

### ChirpStack MQTT Forwarder
[ChirpStack MQTT Forwarder](./chirpstack-mqtt-forwarder/index.md) is an
open-source Protobuf or JSON MQTT packet forwarder, which can either use
the [Semtech UDP Packet Forwarder](https://github.com/Lora-net/packet_forwarder)
or [ChirpStack Concentratord](./chirpstack-concentratord/index.md) as
gateway backend. It is intended to be installed on each gateway.

### ChirpStack Gateway Bridge

[ChirpStack Gateway Bridge](./chirpstack-gateway-bridge/index.md) is an
open-source bridge which converts messages received from the
[Semtech UDP Packet Forwarder](https://github.com/Lora-net/packet_forwarder)
or [Semtech Basics Station](https://github.com/lorabasics/basicstation)
into MQTT. It can be installed on the gateway, or in the cloud.

### ChirpStack

[ChirpStack](./chirpstack/index.md) is an open-source LoRaWAN Network Server which can be used to
to setup private or public LoRaWAN networks. ChirpStack provides a web-interface
for the management of gateways, devices and tenants as well to setup data
integrations with the major cloud providers, databases and services commonly
used for handling device data. ChirpStack provides a gRPC based API that can
be used to integrate or extend ChirpStack.
