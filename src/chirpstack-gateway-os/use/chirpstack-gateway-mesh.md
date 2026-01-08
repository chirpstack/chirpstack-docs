# ChirpStack Gateway Mesh

This pages describes how to configure the ChirpStack Gateway Mesh feature on the
ChirpStack Gateway OS. In case you have not yet configured you ChirpStack
Gateway OS instance, then please refer to
[Getting Started](../getting-started.md) first. For more information about the
ChirpStack Gateway Mesh feature, please refer to the
[ChirpStack Gateway Mesh documentation](../../chirpstack-gateway-mesh/index.md).

<!-- toc -->

## Before you start

### Minimum architecture

To start using the Gateway Mesh feature, you must at least configure one LoRa
Gateway as _Border Gateway_ and an other as _Relay Gateway_.

#### Border Gateway

This Gateway forwards (Mesh) uplinks to ChirpStack using the ChirpStack MQTT
Forwarder like a regular LoRa Gateway that you would use with ChirpStack. The
only difference is that the ChirpStack MQTT Forwarder uses the ChirpStack
Gateway Mesh as backend instead of the ChirpStack Concentratord.

#### Relay Gateway

This Gateway is not connected to ChirpStack and is able to operate without
internet connection. It will encapsulate received LoRaWAN uplink frames in the
[Mesh package format](../../chirpstack-gateway-mesh/protocol.md), and unwraps
Mesh encapsulated LoRaWAN downlinks.

### Concentrator slots

LoRa Gateways can have one or multiple concentrator modules. For example they
could contain a single EU868 module, one EU868 + one 2.4GHz module, two EU868
modules etc.

In case your LoRa Gateway only supports one module, then these must be
configured as _single slot gateways_, in any other case, you must configure
either _Slot 1_ or _Slot 2_ in the configuration. See also the ChirpStack
Concentratord configuration.

## Configuration

### ChirpStack Gateway Mesh

1. Within the web-interface, in the left menu click **ChirpStack**, then click
   **Gateway Mesh**.
2. In the **Global Configuration** tab you must at least configure:
   - _Enabled_: Check this box to enable the ChirpStack Gateway Mesh.
   - _Region_: Select the region for which you would like to setup the
     ChirpStack Gateway Mesh.
3. In the **Mesh Configuration** tab you must at least configure:
   - _Region_: Select the region to use for the Mesh communication. Note that
     this does not have to be the same region as in step 2 as this region is
     used for the Mesh protocol only. For example you could use the ISM2400
     region for Mesh communication (if the gateways support it).
   - _Signing key_: Configure the 128 bit (HEX) signing key used to signing and
     validate the Mesh communication.
   - _Border Gateway_: Enable this checkbox if the gateway must operate as a
     Border Gateway.
   - _Ignore direct uplinks_: Enable this checkbox if the gateway is a Border
     Gateway and you would like to ignore non-Mesh uplinks.
4. The **Mesh data-rate configuration** tab is automatically configured based on
   the selected _Region_ in the **Mesh Configuration**. It can be modified if
   you would like to use different data-rate configuration.
5. In the **Backend configuration** you must configure:
   - _Slot_: The concentrator slot used for communication with LoRaWAN End
     Devices.
6. In the **Mesh Backend configuration** you must configure:
   - _Slot_: The concentrator slot used for Mesh communication. Note that this
     can be the same slot as configured in the **Backend Configuration**.
7. Click **Save & Apply** to apply all changes.

### ChirpStack MQTT Forwarder

1. Within the web-interface, in the left menu click **ChirpStack**, then click
   **MQTT Forwarder**.
2. On the top of the page, you will find the following tabs, **Slot 1** (,
   **Slot 2**) and **Mesh**.
   - Make sure that the ChirpStack MQTT Forwarder is disabled for **Slot 1** and
     if available, **Slot 2**. Click **Save & Apply** before going to the next
     tab.
   - For the **Mesh** tab you must configure:
     - **Global configuration** tab:
       - _Enabled_: Only enable the ChirpStack MQTT Forwarder in case it is not
         configured as a Border Gateway.
     - In the **MQTT configuration** tab you must at least configure (for
       non-Border Gateways):
       - _Server_: The address of the MQTT broker.
     - Click **Save & Apply** to apply all changes.
