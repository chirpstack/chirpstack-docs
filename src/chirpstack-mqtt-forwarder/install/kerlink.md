# Kerlink

<!-- toc -->

## Kerlink KerOS based gateways

**Note:** This was tested with the following firmware version: `5.5.4`.

### SSH into the gateway

The first step is to log in into the gateway using ssh:

```bash
ssh root@GATEWAY-IP-ADDRESS
```

Please refer to the [Kerlink wiki](http://wikikerlink.fr/wirnet-productline)
for login instructions.

### Packet-forwarder configuration

#### Option 1 - Kerlink CPF

##### Enable Kerlink CPF

By default, the Kerlink Common Packet-Forwarder (CPF) is disabled. Please
make sure it is enabled. The following command can be used to enable the CPF:

```bash
klk_apps_config --activate-cpf
```

##### Configure Kerlink CPF

You must configure the packet-forwarder on the gateway to forward its data to
`127.0.0.1` at port `1700`. The file `/etc/lorafwd.toml` must contain the
following lines under the `[ gwmp ]` section:

```toml
node = "127.0.0.1"
service.uplink = 1700
service.downlink = 1700
```

After updating this configuration file, make sure to restart the `lorafwd` service:

```bash
monit restart lorafwd
```

#### Option 2 - ChirpStack Concentratord

ChirpStack Concentratord supports some Kerlink gateways. Please refer to the
ChirpStack Concentratord [installation instructions](../../chirpstack-concentratord/installation/kerlink.md)
for installing the ChirpStack Concentratord on Kerlink gateways.

### Install ChirpStack MQTT Forwarder

#### Download IPK

Use the following commands to download the latest version of the
`chirpstack-mqtt-forwarder` package:

```bash
cd /user/.updates
wget https://artifacts.chirpstack.io/downloads/chirpstack-mqtt-forwarder/vendor/kerlink/klkgw/chirpstack-mqtt-forwarder_%MQTT_FORWARDER_VERSION-r1_klkgw.ipk
```

#### Install IPK

Run the following command to install the IPK package:

```bash
sync
kerosd -u
reboot
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
