# Introduction

ChirpStack MQTT Forwarder is a MQTT packet forwarder for LoRa gateways.
By default it forwards packets in [Protobuf](https://developers.google.com/protocol-buffers)
binary format, optionally it can be configured to use JSON encoding for
debugging. In contrast to the [ChirpStack Gateway Bridge](../chirpstack-gateway-bridge/index.md),
this component must always be installed on the gateway.

## Backends

The following backends are provided:

* [ChirpStack Concentratord](../chirpstack-concentratord/index.md)
* [Semtech UDP Packet Forwarder](https://github.com/Lora-net/packet_forwarder)

### ChirpStack Concentratord

```dot process
digraph G {
	node [shape=record,fontsize="10"];
	edge [fontsize="10"];
	fontsize="10";

	subgraph cluster_0 {
		style=filled;
		color="#bbdefb";
		node [style=filled,color="#e3f2fd"];
		label="LoRa&reg; Gateway";

		"chirpstack-mqtt-forwarder" -> "chirpstack-concentratord" [dir="both" label="ZeroMQ"];

		"chirpstack-concentratord" [label="ChirpStack Concentratord"];
		"chirpstack-mqtt-forwarder" [label="ChirpStack MQTT Forwarder"];

	}

	subgraph cluster_1 {
		style=filled;
		color="#bbdefb";
		node [style=filled,color="#e3f2fd"];
		label="Cloud / server / VM";

		"mqtt-broker" [label="MQTT broker"];
	}

	"chirpstack-mqtt-forwarder" -> "mqtt-broker" [dir="both" label="MQTT"];
}
```

### Semtech UDP Packet Forwarder

```dot process
digraph G {
	node [shape=record,fontsize="10"];
	edge [fontsize="10"];
	fontsize="10";

	subgraph cluster_0 {
		style=filled;
		color="#bbdefb";
		node [style=filled,color="#e3f2fd"];
		label="LoRa&reg; Gateway";

		"chirpstack-mqtt-forwarder" -> "semtech-udp-packet-forwarder" [dir="both" label="UDP"];

		"semtech-udp-packet-forwarder" [label="Semtech UDP Packet Forwarder"];
		"chirpstack-mqtt-forwarder" [label="ChirpStack MQTT Forwarder"];

	}

	subgraph cluster_1 {
		style=filled;
		color="#bbdefb";
		node [style=filled,color="#e3f2fd"];
		label="Cloud / server / VM";

		"mqtt-broker" [label="MQTT broker"];
	}

	"chirpstack-mqtt-forwarder" -> "mqtt-broker" [dir="both" label="MQTT"];
}
```
