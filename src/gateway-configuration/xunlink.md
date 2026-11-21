# Xunlink

<!-- toc -->

## Prerequisites

### Device details

- [TM-CP2B product page](https://utilityedgegateways.com/products/tm-cp2b)
- [TM-CP2B user guide](https://utilityedgegateways.com/products/tm-cp2b/user-guide)

The steps below were tested on the TM-CP2B-P (V11), a model in the same
series as the TM-CP2B linked above. The TM-CP2B-P is an industrial Linux
gateway (Ubuntu 24.04) with a Mini PCIe slot for an optional SX1302-based
LoRaWAN concentrator module, connected over SPI (`/dev/spidev0.0`).

### Network connectivity

- Ensure the gateway can reach the server running the ChirpStack Gateway
  Bridge (UDP port `1700`).

## Configuring the gateway

### Build the concentrator HAL and packet forwarder

The SX1302 HAL source for the TM-CP2B series is provided by Xunlink on
request. Please contact Xunlink through the product page linked above to
obtain the download link.

Log in to the gateway (serial console or SSH) and run:

    git clone <SX1302 HAL source URL provided by Xunlink>
    cd sx1302
    git checkout cp2b
    make clean all
    make -j8
    cp tools/reset_lgw.sh util_chip_id/
    cp tools/reset_lgw.sh packet_forwarder/
    cp tools/reset_lgw.sh libloragw/

### Read the gateway EUI

    cd util_chip_id/
    sudo ./chip_id

Note the EUI printed by `chip_id`. You need it when adding the gateway in
ChirpStack.

### Point the packet forwarder at ChirpStack

In `packet_forwarder/`, edit the `global_conf.json.sx1250.<REGION>` file
that matches your module's frequency band (tested with
`global_conf.json.sx1250.EU868`), and set the `gateway_conf` section:

    "gateway_ID": "<EUI from chip_id>",
    "server_address": "<ChirpStack Gateway Bridge host>",
    "serv_port_up": 1700,
    "serv_port_down": 1700

Start the forwarder, passing the configuration file:

    cd packet_forwarder/
    sudo ./lora_pkt_fwd -c global_conf.json.sx1250.EU868

On units without a GPS module, a warning about `/dev/ttyS0` appears at
startup. This does not affect packet forwarding.

## Adding the gateway to ChirpStack

1. Go to `Tenants -> Gateways` and click `Add gateway`.
2. Enter a name and the `Gateway ID (EUI64)` read with `chip_id`.
3. Click `Submit`.

## Verifying the connection

In `Tenants -> Gateways`, the gateway status should show `Online`.

(Tested with ChirpStack 4.20.0 on the TM-CP2B-P V11.)
