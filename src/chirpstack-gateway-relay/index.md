# Introduction

The ChirpStack Gateway Relay is a software component that can run on a
LoRa<sup>&reg;</sup> gateway, turning the gateway into a relay.

The aim of this component is to extend LoRaWAN<sup>&reg;</sup> coverage without
the requirement to implement a Relay protocol into the end-device
(e.g. [TS011](https://resources.lora-alliance.org/technical-specifications/ts011-1-0-0-relay)).

## Architecture

```dot process
digraph G {
    node [shape=record,fontsize="10"];
    edge [fontsize="10"];
    fontsize="10";
    style=filled;
    color="#bbdefb";
    node [style=filled,color="#e3f2fd"];

    "End Device" -> "Relay Gateway" [dir="both", label="LoRaWAN payload"];
    "Relay Gateway" -> "Border Gateway" [dir="both", label="Relay encapsulated LoRaWAN payload"];
    "Border Gateway" -> "ChirpStack" [dir="both", label="LoRaWAN payload + Relay context"];
}
```

### End Device

LoRaWAN End Device, e.g. LoRaWAN 1.0.x or 1.1.x. Does not need modifications
or implementation TS011 specification.

### Relay Gateway

LoRa Gateway, e.g. SX1301/2/3 based (optionally + ISM2400 concentrator module).
This gateway does not have an internet backhaul and potentially could be solar
powered. It runs the ChirpStack Gateway Relay forwarder which handles the
relaying of uplink and downlink packages between the Relay Gateway and the
Border Gateway. For the Relay Gateway <> Border Gateway, it might use the same
radio band as the End Device, or it could use the ISM2400 band (based on
hardware capabilities + use-case requirements).

In this setup, only the ChirpStack Concentratord and ChirpStack Gateway Relay
must be installed on the gateway.

### Border Gateway

LoRa Gateway, which is internet connected. It handles the wrapping / unwrapping
of the Relay encapsulated LoRaWAN payloads. The Border Gateway <> ChirpStack
communication is as if the Border Gateway is directly communicating with the
End Device. The only exception is that on uplink it sets a `context` with Relay
specific data, which must be returned by ChirpStack on downlink.

In this setup, the ChirpStack Concentratord, ChirpStack Gateway Relay and
ChirpStack MQTT Forwarder must be installed on the gateway. ChirpStack MQTT
Forwarder will in this case be configured to use the ChirpStack Gateway Relay
API interface instead of the ChirpStack Concentratord.


```dot process
digraph G {
    node [shape=record,fontsize="10"];
    edge [fontsize="10"];
    fontsize="10";
    style=filled;
    color="#bbdefb";
    node [style=filled,color="#e3f2fd"];

    subgraph cluster_0 {
        label = "LoRa&reg; Gateway";
        node [style=filled,color="#e3f2fd"];

        "ChirpStack Gateway Relay" -> "ChirpStack Concentratord" [dir="both", label="ZeroMQ"];
        "ChirpStack MQTT Forwarder" -> "ChirpStack Gateway Relay" [dir="both", label="ZeroMQ"];
        "ChirpStack MQTT Forwarder" -> "MQTT Broker" [dir="both", label="MQTT"];
    }
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
