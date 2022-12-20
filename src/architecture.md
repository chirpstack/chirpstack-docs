# Architecture

<!-- toc -->

A typical ChirpStack deployment has the following architecture. Note that in
this diagram, the gateways are connected in different ways to ChirpStack to
highlight the different possible connectivity options.

```dot process
digraph G {
	node [shape=record,fontsize="10"];
	edge [fontsize="10"];
	fontsize="10";

	subgraph cluster_0 {
		style=filled;
		color="#bbdefb";
		node [style=filled,color="#e3f2fd"];

		"chirpstack" -> "pub-sub" [dir="both",label="MQTT"];
		"chirpstack-gateway-bridge-cloud" -> "pub-sub" [dir="both" label="MQTT"];

		"chirpstack" [label="ChirpStack"];
		"pub-sub" [label="MQTT broker"];
		"chirpstack-gateway-bridge-cloud" [label="ChirpStack Gateway Bridge"];

		label = "Cloud / server / VM";
	}

	subgraph cluster_1 {
		style=filled;
		color="#bbdefb";
		node [style=filled,color="#e3f2fd"];
		label="LoRa&reg; Gateway";

		"chirpstack-mqtt-forwarder-gw1" -> "pub-sub" [label="MQTT",dir="both"];
		"chirpstack-mqtt-forwarder-gw1" [label="UDP Packet Forwarder +\nChirpStack MQTT Forwarder"];
	}

	subgraph cluster_2 {
		style=filled;
		color="#bbdefb";
		node [style=filled,color="#e3f2fd"];
		label="LoRa&reg; Gateway";

		"chirpstack-mqtt-forwarder-gw2" -> "pub-sub" [label="MQTT",dir="both"];
		"chirpstack-mqtt-forwarder-gw2" [label="ChirpStack Concentratord +\nChirpStack MQTT Forwarder"];
	}

	subgraph cluster_3 {
		style=filled;
		color="#bbdefb";
		node [style=filled,color="#e3f2fd"];
		label="LoRa&reg; Gateway";

		"packet-forwarder-gw2" -> "chirpstack-gateway-bridge-cloud" [label="UDP",dir="both"];
		"packet-forwarder-gw2" [label="UDP Packet Forwarder"];
	}

	subgraph cluster_4 {
		style=filled;
		color="#bbdefb";
		node [style=filled,color="#e3f2fd"];
		label="LoRa&reg; Gateway";

		"packet-forwarder-gw3" -> "chirpstack-gateway-bridge-cloud" [label="Websockets",dir="both"];
		"packet-forwarder-gw3" [label="Basics Station"];
	}

	subgraph cluster_5 {
		style=filled;
		color="#bbdefb";
		node [style=filled,color="#e3f2fd"];
		label="Integrations";

		"http-int" [label="HTTP"];
		"mqtt-int" [label="MQTT"];
		"other-int" [label="Etc ..."];
	}


	"chirpstack" -> "http-int";
	"chirpstack" -> "mqtt-int";
	"chirpstack" -> "other-int";

	"as-api-client" -> "chirpstack" [label="gRPC"];
	"as-api-client" [label="API client"];
}
```

## Multi-region example 1

This example displays a multi-region ChirpStack installation (EU868 and US915)
using a single MQTT broker. Please note the MQTT prefix. In this case for US915
the prefix `us915_0` is used, indicating that this region uses the first 8
channels of the US915 region (the prefix is user-configurable). To keep the
diagram simple, some elements of the previous diagram have been left out.

```dot process
digraph G {
	node [shape=record,fontsize="10"];
	edge [fontsize="10"];
	fontsize="10";

	subgraph cluster_0 {
		style=filled;
		color="#bbdefb";
		node [style=filled,color="#e3f2fd"];

		"chirpstack" -> "pub-sub" [dir="both",label="MQTT"];
		"chirpstack-gateway-bridge-cloud-eu868" -> "pub-sub" [dir="both" label="MQTT\neu868/gateway/..."];
		"chirpstack-gateway-bridge-cloud-us915" -> "pub-sub" [dir="both" label="MQTT\nus915_0/gateway/..."];

		"chirpstack" [label="ChirpStack"];
		"pub-sub" [label="MQTT broker"];
		"chirpstack-gateway-bridge-cloud-eu868" [label="ChirpStack Gateway Bridge (EU868)"];
		"chirpstack-gateway-bridge-cloud-us915" [label="ChirpStack Gateway Bridge (US915)"];

		label = "Cloud / server / VM";
	}

	subgraph cluster_1 {
		style=filled;
		color="#bbdefb";
		node [style=filled,color="#e3f2fd"];
		label="LoRa&reg; Gateway (EU868)";

		"chirpstack-gateway-bridge-gw-eu868" -> "pub-sub" [label="MQTT\neu868/gateway/...",dir="both"];
		"chirpstack-gateway-bridge-gw-eu868" [label="UDP Packet Forwarder +\nChirpStack MQTT Forwarder"];
	}


	subgraph cluster_2 {
		style=filled;
		color="#bbdefb";
		node [style=filled,color="#e3f2fd"];
		label="LoRa&reg; Gateway (EU868)";

		"packet-forwarder-gw2-eu868" -> "chirpstack-gateway-bridge-cloud-eu868" [label="UDP",dir="both"];
		"packet-forwarder-gw2-eu868" [label="UDP Packet Forwarder"];
	}

	subgraph cluster_3 {
		style=filled;
		color="#bbdefb";
		node [style=filled,color="#e3f2fd"];
		label="LoRa&reg; Gateway (US915)";

		"chirpstack-gateway-bridge-gw-us915" -> "pub-sub" [label="MQTT\nus915_0/gateway/...",dir="both"];
		"chirpstack-gateway-bridge-gw-us915" [label="UDP Packet Forwarder +\nChirpStack MQTT Forwarder"];
	}


	subgraph cluster_4 {
		style=filled;
		color="#bbdefb";
		node [style=filled,color="#e3f2fd"];
		label="LoRa&reg; Gateway (US915)";

		"packet-forwarder-gw2-us915" -> "chirpstack-gateway-bridge-cloud-us915" [label="UDP",dir="both"];
		"packet-forwarder-gw2-us915" [label="UDP Packet Forwarder"];
	}
}
```

## Multi-region example 2

As an alternative to the previous diagram, it is also possible to setup a
separate MQTT broker per region. In this it is no longer a requirement to use
the region prefix in the MQTT topic.

```dot process
digraph G {
	node [shape=record,fontsize="10"];
	edge [fontsize="10"];
	fontsize="10";

	subgraph cluster_0 {
		style=filled;
		color="#bbdefb";
		node [style=filled,color="#e3f2fd"];

		"chirpstack" -> "pub-sub-eu868" [dir="both",label="MQTT"];
		"chirpstack" -> "pub-sub-us915" [dir="both",label="MQTT"];
		"chirpstack-gateway-bridge-cloud-eu868" -> "pub-sub-eu868" [dir="both" label="MQTT"];
		"chirpstack-gateway-bridge-cloud-us915" -> "pub-sub-us915" [dir="both" label="MQTT"];

		"chirpstack" [label="ChirpStack"];
		"pub-sub-eu868" [label="MQTT broker EU868"];
		"pub-sub-us915" [label="MQTT broker US915"];
		"chirpstack-gateway-bridge-cloud-eu868" [label="ChirpStack Gateway Bridge (EU868)"];
		"chirpstack-gateway-bridge-cloud-us915" [label="ChirpStack Gateway Bridge (US915)"];

		label = "Cloud / server / VM";
	}

	subgraph cluster_1 {
		style=filled;
		color="#bbdefb";
		node [style=filled,color="#e3f2fd"];
		label="LoRa&reg; Gateway (EU868)";

		"chirpstack-gateway-bridge-gw-eu868" -> "pub-sub-eu868" [label="MQTT",dir="both"];
		"chirpstack-gateway-bridge-gw-eu868" [label="UDP Packet Forwarder +\nChirpStack MQTT Forwarder"];
	}


	subgraph cluster_2 {
		style=filled;
		color="#bbdefb";
		node [style=filled,color="#e3f2fd"];
		label="LoRa&reg; Gateway (EU868)";

		"packet-forwarder-gw2-eu868" -> "chirpstack-gateway-bridge-cloud-eu868" [label="UDP",dir="both"];
		"packet-forwarder-gw2-eu868" [label="UDP Packet Forwarder"];
	}

	subgraph cluster_3 {
		style=filled;
		color="#bbdefb";
		node [style=filled,color="#e3f2fd"];
		label="LoRa&reg; Gateway (US915)";

		"chirpstack-gateway-bridge-gw-us915" -> "pub-sub-us915" [label="MQTT",dir="both"];
		"chirpstack-gateway-bridge-gw-us915" [label="UDP Packet Forwarder +\nChirpStack MQTT Forwarder"];
	}


	subgraph cluster_4 {
		style=filled;
		color="#bbdefb";
		node [style=filled,color="#e3f2fd"];
		label="LoRa&reg; Gateway (US915)";

		"packet-forwarder-gw2-us915" -> "chirpstack-gateway-bridge-cloud-us915" [label="UDP",dir="both"];
		"packet-forwarder-gw2-us915" [label="UDP Packet Forwarder"];
	}
}
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
