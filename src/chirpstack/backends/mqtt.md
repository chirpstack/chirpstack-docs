# MQTT

The MQTT backend is the default backend to communicate with the LoRa<sup>&reg;</sup>
gateways. If supported by the MQTT broker, it is recommended to make use of a
shared subscription, such that in case of multiple ChirpStack instances, each
gateway event message is delivered to only one ChirpStack instance instead of
all instances. In case this is not supported, please note that the ChirpStack
de-duplication handler will make sure that duplicates are handled correctly.

## Architecture

```dot process
digraph G {
	fontsize=10;
	style=filled;
	color="#bbdefb";
	node [shape=record, style=filled, color="#e3f2fd", fontsize=10];
	edge [fontsize=9];

	subgraph cluster_0 {
		label="LoRa&reg; Gateway";

		"chirpstack-gateway-bridge-gw" [label="Packet Forwarder +\nChirpStack Gateway Bridge"];
	}

	subgraph cluster_1 {
		label="Cloud / server / VM";

		"mqtt" [label="MQTT broker"];
		"chirpstack" [label="ChirpStack"];
	}


	"chirpstack-gateway-bridge-gw" -> "mqtt" [label="REGION/gateway/ID/event/EVENT"];
	"mqtt" -> "chirpstack-gateway-bridge-gw" [label="REGION/gateway/ID/command/COMMAND"];

	"chirpstack" -> "mqtt" [label="REGION/gateway/ID/command/COMMAND"];
	"mqtt" -> "chirpstack" [label="REGION/gateway/ID/event/EVENT"];
}
```

**Note:** In the graph above, the ChirpStack Gateway Bridge is
installed on the gateway. It is also possible to install the ChirpStack
Gateway Bridge in the cloud.
