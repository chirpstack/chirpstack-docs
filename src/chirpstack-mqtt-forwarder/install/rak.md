# RAK

<!-- toc -->

## RAK7268V2

* [Product details](https://store.rakwireless.com/products/rak7268-8-channel-indoor-lorawan-gateway)

### Configure Packet Forwarder

In the RAK WisGate OS web-interface, you must configure the Packet Forwarder
such that it forwards to `localhost` on port `1700`.

Please refer to the [WisGate OS 2 manual](https://docs.rakwireless.com/Product-Categories/Software-APIs-and-Libraries/WisGateOS-2/Overview/)
for instructions on how to access the web-interface.

* In the left menu, click the **Configuration** icon (second icon from top to bottom).
* Configure the following settings:
  * **Work mode:** _Packet forwarder_
  * Under **Protocol**:
    * **Protocol:** _Semtech UDP GWMP Protocol_
  * Under **UDP Protocol parameters**:
    * **Server address:** _localhost_
    * **Server port up:** _1700_
    * **Server port down:** _1700_
* Click **Save changes**

### Install ChirpStack MQTT Forwarder

#### SSH login

First you must login into the gateway using SSH:

```bash
ssh root@GATEWAY-IP-ADDRESS
```

You must use the password as configured during the setup of the gateway.

#### Download IPK

Use the following commands to download the latest version of the
`chirpstack-mqtt-forwarder` package:

```bash
cd /tmp
wget https://artifacts.chirpstack.io/downloads/chirpstack-mqtt-forwarder/vendor/rak/mipsel_24kc/chirpstack-mqtt-forwarder_{{mqtt_forwarder_version}}-r1_mipsel_24kc.ipk
```

#### Install IPK

Use the `opkg` package-manager to install the downloaded package. Example:

```bash
opkg install chirpstack-mqtt-forwarder_{{mqtt_forwarder_version}}-r1_mipsel_24kc.ipk
```

#### Configuration

To connect the ChirpStack MQTT Forwarder to your MQTT broker, you must update
the ChirpStack MQTT Forwarder configuration file. This file is located at:
`/etc/chirpstack-mqtt-forwarder/chirpstack-mqtt-forwarder.toml`.

#### (Re)start and stop commands

Use the following commands to (re)start and stop the ChirpStack MQTT Forwarder service:

```bash
# start
/etc/init.d/chirpstack-mqtt-forwarder start

# stop
/etc/init.d/chirpstack-mqtt-forwarder stop

# restart
/etc/init.d/chirpstack-mqtt-forwarder restart
```

### Access log output

Use the following command to access the ChirpStack MQTT Forwarder log output:

```bash
logread -f -l 100 -e chirpstack-mqtt-forwarder
```
