# Roaming with Helium

This guide explains step-by-step how to setup roaming with the
[Helium](https://www.helium.com/) network. Roaming with Helium allows you to run
your own ChirpStack instance and using the coverage provided by Helium. This
would allow you for example:

- To setup a LoRaWAN network without owning any gateways
- Use Helium to provide coverage for areas that are not covered by your own
  gateways

**Note:** Roaming described in this guide works one way, it will not share
uplinks received by your gateways with the Helium network.

<!-- toc -->

## In summary

You need to buy a Helium OUI. To this OUI you can assign your
[LoRa Alliance](https://lora-alliance.org/) NetID, or you can buy a block of
DevAddrs from Helium. This will be used to filter the data that will be
forwarded to your ChirpStack instance. For OTAA join-requests, you will need to
specify the DevEUIs and JoinEUIs of your devices.

Helium will forward the data to you using the
[Semtech UDP Packet Forwarder](https://github.com/Lora-net/packet_forwarder/blob/master/PROTOCOL.TXT)
protocol. Effectively, the Helium roaming works as a virtual gateway.

```d2
vars: {
	d2-config: {
		layout-engine: elk
	}
}

vm: Cloud / VM
helium: Helium Network

vm: {
  chirpstack: ChirpStack
  chirpstack_gateway_bridge: ChirpStack Gateway Bridge
  mqtt_broker: MQTT Broker

  chirpstack <> mqtt_broker: MQTT
  chirpstack_gateway_bridge <> mqtt_broker: MQTT
}

helium: {
  helium_packet_broker: Helium Packet Broker
}

helium.helium_packet_broker <> vm.chirpstack_gateway_bridge: UDP
```

## Step by step guide

### Setup ChirpStack

If you haven't done so already, you need to install ChirpStack first. For
getting started, you could follow one of the following guides:

- [Quickstart Docker Compose](../getting-started/docker.md)
- [Debian / Ubuntu](../getting-started/debian-ubuntu.md)
- [Ansible and Vagrant](../getting-started/ansible-vagrant.md)

It is important that you also install the
[ChirpStack Gateway Bridge](../chirpstack-gateway-bridge/index.md) on your
server, as the ChirpStack Gateway Bridge will be used to receive data from the
Helium network.

Helium uses the the following region configurations. Please make sure that you
have configured ChirpStack Gateway Bridge and ChirpStack accordingly:

| Helium region | Helium region ID | ChirpStack region configuration |
| ------------- | ---------------- | ------------------------------- |
| EU868         | eu868            | eu868                           |
| US915         | us915            | us915_1                         |

### Create Helium Wallet

Please refer to the
[Helium Wallet App documentation](https://docs.helium.com/wallets/helium-wallet-app/)
to learn how you can create a Helium Wallet if you haven't already done so.

### Install `helium-config-service-cli`

Download or compile the `helium-config-service-cli` utility from source. After
this step it is assumed that you can execute this binary using:

```bash
./helium-config-service-cli
```

#### Pre-compiled binaries

For pre-compiled binaries, please see
[https://github.com/helium/helium-config-service-cli](https://github.com/helium/helium-config-service-cli).

#### Compile from source

To compile the `helium-config-service-cli` from source, you must execute the
following commands:

```bash
git clone https://github.com/helium/helium-config-service-cli.git
cd helium-config-service-cli
cargo build --release
```

This does require that you have the Rust compiler installed. As well you need to
have Protobuf installed. Installation steps for these dependencies vary
depending your OS. After `cargo build --release` has compiled, you will find the
binary in `./target/release`.

### Generate key-pairs

In order to authenticate your interactions with the Helium config service, you
need to generate two key-pairs:

- **Delegate Keypair**: The delegate key is actively used to create and manage
  routes. In case it is ever compromised, it can be updated using the owner key.
- **Owner Keypair**: The owner key is irreplaceable and should be kept safe,
  secure, and private at all times. If lost or compromised, you could lose your
  OUI.

#### Generate key-pairs

Execute the following commands to generate the key-pairs:

```bash
./helium-config-service-cli env generate-keypair delegate.bin
./helium-config-service-cli env generate-keypair owner.bin
```

#### Print public-keys

Execute the following commands to print the public-keys of the key-pairs that
you have generated in the previous step, you will need this information in the
next steps:

```bash
./helium-config-service-cli env info --keypair delegate.bin
./helium-config-service-cli env info --keypair owner.bin
```

### Request Helium OUI

At this moment a Helium OUI must be requested through e-mail by sending an
e-mail to [hello@helium.foundation](mailto:hello@helium.foundation). This e-mail
must contain the following information:

- The public key of your owner key pair
- The public key of your delegate key pair
- The Helium Wallet address
- Your LoRa Alliance NetID **-or-**
- The number of DevAddrs you want to purchase (minimum 8)

#### Note on data-credits

Helium will charge US$100 for purchasing the Helium OUI. If you are buying
DevAddrs, then each DevAddr will cost US$100.

Once you have acquired you OUI you will be able to purchase more data-credits
online using credit-card **if you are based in the US**. If you are not based in
the US, you should indicate this in your e-mail to Helium.

Alternatively, you can burn HNT to buy data-credits. Note that in order to roam
with Helium, you must at least have US$35 worth of data-credits.

More in data-credits and information how to buy data-credits using credit-card
can be found in the
[Helium Data Credit documentation](https://docs.helium.com/tokens/data-credit).

### Configure ChirpStack

You must configure the ChirpStack configuration to make sure it is configured
with your LoRa Alliance NetID or with the Helium NetID and DevAddr range that
you have purchased. As well you must enable receiving data from unknown
gateways.

**Note:** After making modifications, please make sure to restart ChirpStack!

#### Allow unknown gateways

As it will not be known in advance which Helium Gateway IDs will send data to
your ChirpStack instance, you need to enable receiving data from unknown
gateways. Under `[gateway]` in your
[ChirpStack configuration](../chirpstack/configuration.md) configure:

```toml
allow_unknown_gateways=true
```

#### LoRa Alliance NetID

If you own a LoRa Alliance NetID and you have associated this with your OUI,
then confirm that the `net_id` under `[network]` is set correctly in your
[ChirpStack configuration](../chirpstack/configuration.md).

#### Helium DevAddrs

If you have purchased a Helium DevAddr range, then you must configure ChirpStack
with the Helium NetID **and** DevAddr range that you have purchased. This way,
ChirpStack will only assign DevAddrs in this range to devices that join the
network.

Under the `[network]` configuration inf your
[ChirpStack configuration](../chirpstack/configuration.md), configure:

- The `net_id` option to `00003C` (Helium NetID)
- The `dev_addr_prefixes` option to your Helium DevAddr prefix

The `dev_addr_prefixes` value can be obtained using the
`helium-config-service-cli` utility.

For example if your DevAddr range starts with `78000068` and ends with
`7800006F` (you have received this information by e-mail from Helium), then you
can execute the following command:

```bash
./target/release/helium-config-service-cli subnet-mask 78000068 7800006F
```

This will output:

```json
{
  "range": {
    "start_addr": "78000068",
    "end_addr": "7800006F"
  },
  "subnets": [
    "78000068/29"
  ]
}
```

Thus you will configure:

```toml
dev_addr_prefixes=[
    "78000068/29",
]
```

### Configure `helium-config-service-cli` environment

In this step you will configure the `helium-config-service-cli` environment
variables, to avoid that you have to enter the same configuration each time when
you execute this command. To start the configuration, execture the following
command:

```bash
./helium-config-service-cli env init
```

Enter the following information:

- **Config Service Host**: `http://mainnet-config.helium.io:6080/`
- **Keypair Location**: `./delegate.bin`
- **Net ID**: `00003C` if you are using the Helium NetID, or else your own NetID
- **Assigned OUI**: Your Helium OUI (you have received this information by
  e-mail from Helium)
- **Default Max Copies**: The max number of copies of each uplink that you would
  like to receive, e.g. `15`

Once you have completed all steps, this will print something like:

```
Put these in your environment
------------------------------------
HELIUM_CONFIG_HOST=http://mainnet-config.helium.io:6080/
HELIUM_KEYPAIR_BIN=./delegate.bin
HELIUM_NET_ID=00003C
HELIUM_OUI=xyz
HELIUM_MAX_COPIES=15
```

One way to put these values in your shell environment is to use `export`. To do
so, execute the following commands:

```bash
export HELIUM_CONFIG_HOST=http://mainnet-config.helium.io:6080/
export HELIUM_KEYPAIR_BIN=./delegate.bin
export HELIUM_NET_ID=00003C
export HELIUM_OUI=xyz
export HELIUM_MAX_COPIES=15
```

### Configure routing

In this step you will configure the routing from the Helium Package Router to
your [ChirpStack Gateway Bridge](../chirpstack-gateway-bridge/index.md). Helium
will route data that:

- Matches the configured DevEUI + JoinEUI (AppEUI)
- Matches the configured DevAddr range
- Matches the configured NetID (if you have your own NetID)

**Important notes:**

- The Helium Package Router will forward the data using the
  [Semtech UDP Packet Forwarder protocol](https://github.com/Lora-net/packet_forwarder/blob/master/PROTOCOL.TXT)
  thus this data must be routed to the ChirpStack Gateway Bridge.
- Make sure that your ChirpStack Gateway Bridge is publicly accessible. By
  default it uses port `1700` to receive gateway data. Please make sure that
  this port (UDP) is open in your firewall!

#### Create empty route

To create your (first) route, execute the following command:

```bash
./helium-config-service-cli route new --commit
```

You will need to use the response `id` as `<route-id>` in the next steps.

#### Configure endpoint

To point this route to your ChirpStack Gateway Bridge instance, execute the
following command:

```bash
./helium-config-service-cli route update server --host <server-ip> --port <server-port> --route-id <route-id> --commit
```

For example, if your ChirpStack Gateway Bridge instance is publicly accessible
at `example.com` on port `1700`, you would execute:

```bash
./helium-config-service-cli route update server --host example.com --port 1700 --route-id <route-id> --commit
```

To configure the (UDP) port for each LoRaWAN region, execute the following
command:

```bash
./helium-config-service-cli route update add-gwmp-region --route-id <route-id> us915 <us915-udp-port> eu868 <eu868-udp-port> --commit
```

For example, if you have two ChirpStack Gateway Bridge instances running:

- `1700` - EU868 configuration
- `1701` - US915 configuration

Then you would execute the following command:

```bash
./helium-config-service-cli route update add-gwmp-region --route-id <route-id> us915 1701 eu868 1700 --commit
```

#### Register device for roaming

The Helium Packet Router will only forward join-requests for registered devices
(DevEUI + JoinEUI). To register a device, execute the following command:

```bash
./helium-config-service-cli route euis add --route-id <route-id> -d <dev-eui> -a <app-eui> --commit
```

#### Register DevAddr range for roaming

While the above step will result in the Helium Packet Router forwarding
join-requests for the configured device(s), it will not forward uplink data, as
uplink messages will contain a DevAddr, not a DevEUI and thus the Helium Packet
Router is unable to filter these messages.

To let the Helium Packet Forwarder forward uplink messages, you have to
configure a DevAddr range (start and end DevAddr). To register a DevAddr range
to your route, execute the following command:

```bash
./helium-config-service-cli route devaddrs add -s <start DevAddr> -e <end DevAddr> --route-id <route-id> -- commit
```

In case you have purchased a Helium DevAddr range, this would match the first
and last DevAddr from this range. In case you are using your own NetID, this
would be the first and last DevAddr from your own DevAddr range.

### Validation

After completing the above steps, you should be able to OTAA activate your
device with the DevEUI and JoinEUI (AppEUI) that you added in the
`route euis add` step. In the ChirpStack LoRaWAN frames tab, you should see the
Join-Request uplink and Join-Accept downlink with the only difference that the
`rxInfo` will contain meta-data from the Helium gateways.

After activation, you device should have a DevAddr assigned that falls within
the range as configured in your ChirpStack configuration and that you have
configured in the `route devaddrs add` step.

### Next steps

At this point Helium will forward all uplinks within the configured DevAddr
range. Depending the use-case, you might want to set the `ignore_empty_skf` flag
to `true`.

If this option is enabled, then you must provide Helium with the `NwkSKey` of
each device. In this case, Helium will perform a MIC check before forwarding the
uplinks and only uplinks with passing MIC check will be forwarded.

Please refer to the CLI documentation for more information to provide this
information to Helium:

```bash
./helium-config-service-cli --help
```
