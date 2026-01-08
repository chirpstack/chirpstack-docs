# Tektelic

## Kona gateways

These instructions appy to the Linux based Tektelic Kona gateways.

- [Product detail page](https://www.tektelic.com/products/gateways/)

### SSH into the gateway

The first step is to login into the gateway using ssh:

```bash
ssh root@GATEWAY-IP-ADDRESS
```

The default password is the serial-number of the gateway which is printed on the
back of the gateway (the 9 characters above the 12V = 1A line).

### Packet-forwarder configuration

You must configure the packet-forwarder on the gateway to forward its data to
`127.0.0.1` at port `1700`. The file `/etc/default/config.json` must contain the
following lines:

```json
   ...
   "gateway_conf" {
       ...
       "server_address": "127.0.0.1",
    "serv_port_up": 1700,
    "serv_port_down": 1700,
    ...
   }
...
```

After updating the configuration file, make sure to restart the
packet-forwarder:

```bash
/etc/init.d/pkt_fwd restart
```

### Install ChirpStack MQTT Forwarder

#### Download IPK

Use the following commands to download the latest version of the
`chirpstack-mqtt-forwarder` package:

```bash
cd /tmp
wget https://artifacts.chirpstack.io/downloads/chirpstack-mqtt-forwarder/vendor/tektelic/kona/chirpstack-mqtt-forwarder_%MQTT_FORWARDER_VERSION-r1_kona.ipk
```

#### Install IPK

Use the `opkg` package-manager to install the downloaded package. Example:

```bash
sudo opkg install chirpstack-mqtt-forwarder_%MQTT_FORWARDER_VERSION-r1_kona.ipk
```

#### Configuration

To connect the ChirpStack MQTT Forwarder to your MQTT broker, you must update
the ChirpStack MQTT Forwarder configuration file. This file is located at:
`/etc/chirpstack-mqtt-forwarder/chirpstack-mqtt-forwarder.toml`.

#### (Re)start and stop commands

Use the following commands to (re)start and stop the ChirpStack MQTT Forwarder
service:

```bash
# Status
sudo monit summary

# start
monit start chirpstack-mqtt-forwarder

# stop
monit stop chirpstack-mqtt-forwarder

# restart
monit restart chirpstack-mqtt-forwarder
```

#### Log output

To view the ChirpStack MQTT Forwarder log output, use the following command:

```bash
tail -f -n 100 /var/log/messages |grep chirpstack-mqtt-forwarder
```
