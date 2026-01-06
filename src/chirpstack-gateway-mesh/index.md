# Introduction

The ChirpStack Gateway Mesh is a software component that can run on
LoRa<sup>&reg;</sup> gateways, turning these gateways either into
Relay Gateways (gateways relaying data, most likely these gateways
are solar powered) or Border Gateways (gateways terminating the mesh 
protocol and directly communicating with ChirpStack).

The aim of this component is to extend LoRaWAN<sup>&reg;</sup> coverage,
by adding LoRa gateways that are not connected to the internet which
will repeat uplink and downlink LoRaWAN payloads. This can be useful in
remote areas where internet coverage is sparse.
This solution is different from the  LoRa Alliance Relay Protocol,
as no software implementation changes are required in the device stack.

It is possible that there are multiple Relay Gateways inbetween the End-Device
and the Border Gateway. The technical limitation is 8 hops (the first Relay
Gateway counts as the 1st), but this is also influenced by required airtime
to receive and relay these payloads.

The ChirpStack Gateway Mesh component has been developed in collaboration
with [RAKwireless](https://learn.rakwireless.com/hc/en-us/articles/26826770321559-How-To-Set-Up-Gateway-Mesh-Quick-Start-Using-ChirpStackOS-on-RAK-Gateways)
and [Smart Parks](https://www.smartparks.org/).

## ChirpStack Gateway Mesh architecture

```d2
vars: {
	d2-config: {
		layout-engine: elk
	}
}

End Device -> Relay Gateway 1: LoRaWAN payload
Relay Gateway 1 <> Relay Gateway 2: Mesh encapsulated LoRaWAN payload
Relay Gateway 2 <> Border Gateway: Mesh encapsulated LoRaWAN payload
Border Gateway <> ChirpStack: LoRaWAN payload + Mesh context blob

```

### End Device

LoRaWAN End Device, e.g. LoRaWAN 1.0.x or 1.1.x. Does not need modifications
or implementation TS011 specification.

### Relay Gateway

LoRa Gateway, e.g. SX1301/2/3 based (optionally + ISM2400 concentrator module).
This gateway does not have an internet backhaul and could be solar
powered. It runs the ChirpStack Gateway Mesh component which handles the
relaying of uplink and downlink packages between the Relay Gateway and the
Border Gateway. For the Relay Gateway <> Border Gateway, it might use the same
radio band as the End Device, or it could use the ISM2400 band (based on
hardware capabilities + use-case requirements).

In this setup, only the ChirpStack Concentratord and ChirpStack Gateway Mesh
components must be installed on the gateway.

```d2
vars: {
	d2-config: {
		layout-engine: elk
	}
}

ChirpStack Gateway Mesh <> ChirpStack Concentratord
```

### Border Gateway

LoRa Gateway, which is internet connected. It handles the wrapping / unwrapping
of the Mesh encapsulated LoRaWAN payloads. The Border Gateway <> ChirpStack
communication is as if the Border Gateway is directly communicating with the
End Device. The only exception is that on uplink it sets a `context` with Mesh
specific data, which must be returned by ChirpStack on downlink.

In this setup, the ChirpStack Concentratord, ChirpStack Gateway Mesh and
ChirpStack MQTT Forwarder must be installed on the gateway. ChirpStack MQTT
Forwarder will in this case be configured to use the ChirpStack Gateway Mesh
API interface instead of the ChirpStack Concentratord.

```d2
vars: {
	d2-config: {
		layout-engine: elk
	}
}

lora_gateway: LoRa Gateway {
    ChirpStack Gateway Mesh <> ChirpStack Concentratord
    ChirpStack MQTT Forwarder <> ChirpStack Gateway Mesh
    ChirpStack MQTT Forwarder <> MQTT Broker
}
```

### ChirpStack

ChirpStack will receive the uplink from the Border Gateway, with the RX
parameters from the Relay Gateway. Thus it receives the original:

* Data-rate
* Frequency
* RSSI
* SNR

Returning the uplink `context` "blob" on downlink, it is able to transmit
downlinks to the End Device.
