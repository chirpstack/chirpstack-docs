# Comparison

The ChirpStack Gateway Mesh and the LoRa Alliance Relay Specification each have
their own benefits. Depending the use-case one or the other might provide a
better solution. To understand which solution might be better, this page
highlights the main differences between the two.

## Implementation

* **Gateway Mesh**: Does not require changes to LoRaWAN end-devices. The Gateway
    Mesh protocol only needs to be implemented by the LoRa Gateway. Border Gateway
    to ChirpStack communication is the same as a regular LoRa Gateway communicating
    to ChirpStack.
* **Relay Specification**: Requires implementation by the LoRaWAN Network Server
    and the LoRaWAN end-device. ChirpStack does implement the Relay Specification.

## Security

* **Gateway Mesh**: Can filter received data based on NetID and JoinEUI. Is not
    LoRaWAN end-device aware, and forwards all data passing the configured
    filters. The Mesh protocol uses signing key to protect against spoofing of
    Mesh payloads.
* **Relay Specification**: Can filter received data based on JoinEUI + DevEUI.
    is LoRaWAN end-device aware and only forwards data for these Devices
    (trusted end-device list).

## Capacity

* **Gateway Mesh**: Relay Gateway limitations are defined by max. throughput of
    LoRa concentrator. Mesh communication can optionally use LoRaWAN 2.4 GHz to
    increase throughput.
* **Relay Sepcification**: The trusted end-device list of the Relay device
    can store max. 16 end-devices. The Relay device uses the same band as the
    end-devices.

## Mobility

* **Gateway Mesh**: Relay Gateways within the Mesh are not end-device
    aware, devices can move freely without re-synchronizing between ChirpStack
    and the Relay Gateways.
* **Relay Specification**: If an end-device moves from one Relay device to
    an other Relay, then ChirpStack must re-synchronize the trusted end-device
    list (removing the end-device from on list, adding it to the other).

## Hardware

* **Gateway Mesh**: Gateway Mesh functionality is a software layer on top of
    LoRa Gateway hardware.
* **Relay Specification**: Relay devices are technically LoRaWAN end-devices.    
