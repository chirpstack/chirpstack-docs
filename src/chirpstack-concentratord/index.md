# Introduction

ChirpStack Concentratord is an open-source LoRa(WAN) concentrator daemon built
on top of the Semtech hardware abstraction layers. It exposes a [ZeroMQ](https://zeromq.org/)
based API that can be used by one or multiple applications to interact with
gateway hardware. By implementing and abstracting the the hardware specifics
in a separate daemon and exposing this over a ZeroMQ based API, packet forwarding
applications can be completely decoupled from the gateway hardware.
It also makes it possible to let multiple applications interact simultaneously
with the gateway hardware. For example multiple packet forwarders could forward
data to different LoRaWAN network servers.

## HAL support

ChirpStack Concentratord provides an unified ZeroMQ API for the following
Semtech HALs:

* [SX1301](https://github.com/Lora-net/lora_gateway)
* [SX1302/3](https://github.com/Lora-net/sx1302_hal/)
* [2g4](https://github.com/Lora-net/gateway_2g4_hal)

## Architecture example

```dot process
digraph G {
	node [shape=record,fontsize="10"];
	edge [fontsize="10"];
	fontsize="10";

	subgraph cluster_0 {
		style=filled;
		color="#bbdefb";
		node [style=filled,color="#e3f2fd"];
		label = "LoRa&reg; Gateway";

		"chirpstack-concentratord" -> "concentrator" [dir="both", label="Semtech HAL"];
		"chirpstack-mqtt-forwarder" -> "chirpstack-concentratord" [dir="both", label="ZeroMQ"];
		"chirpstack-udp-forwarder" -> "chirpstack-concentratord" [dir="both", label="ZeroMQ"];
		"monitoring" -> "chirpstack-concentratord" [dir="both", label="ZeroMQ"];

		"concentrator" [label="SX1301 / SX1302 / 2G4"];
		"chirpstack-concentratord" [label="ChirpStack Concentratord"];
		"chirpstack-mqtt-forwarder" [label="ChirpStack MQTT Forwarder"];
		"chirpstack-udp-forwarder" [label="ChirpStack UDP Forwarder"];
		"monitoring" [label="Third-party monitoring app"];
	}

}
```
