# Dragino

<!-- toc -->

## LPS8(N) Indoor LoRaWAN Gateway

* [Product details](https://www.dragino.com/products/lora-lorawan-gateway/item/200-lps8n.html)

**Note:** This was tested with the following firmware version: `lgw-5.4.1668567157`.

### Configure Packet Forwarder

In the Dragino web-interface, you must configure the Packet Forwarder such
that it forwards to `localhost` on port `1700`.

By default, the web-interface can be accessed by entering the following URL
in your browser: `https://GATEWAY-IP-ADDRESS:8000` (replace `GATEWAY-IP-ADDRESS`
by the actual IP address of your gateway). The default credentials are
`root` / `dragino`.

* In the **LoRaWAN** menu, click **LoRaWAN -- Semtech UDP**
* Configure the following settings:
  * **Service Provider:** _Custom / Private LoRaWAN_
  * **Server Address:** _localhost_
  * **Uplink Port:** _1700_
  * **Downlink Port:** _1700_
* Click **Save & Apply**

### Install ChirpStack MQTT Forwarder

#### SSH login

First you must login into the gateway using SSH:

```bash
ssh -p 2222 root@GATEWAY-IP-ADDRESS
```

The default password is _dragino_.

#### Download IPK

Use the following commands to download the latest version of the
`chirpstack-mqtt-forwarder` package:

```bash
cd /tmp
wget https://artifacts.chirpstack.io/downloads/chirpstack-mqtt-forwarder/vendor/dragino/mips_24kc/chirpstack-mqtt-forwarder_$MQTT_FORWARDER_VERSION-r1_mips_24kc.ipk
```

#### Install IPK

Use the `opkg` package-manager to install the downloaded package. Example:

```bash
opkg install chirpstack-mqtt-forwarder_$MQTT_FORWARDER_VERSION-r1_mips_24kc.ipk
```

**Note:** In case of an upgrade, it is recommeded to first uninstall the
`chirpstack-mqtt-forwarder` package using `opkg remove ...` and then install the
new version using `opkg install ...`. Configuration files will be maintained.

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

## LPS8V2 LoRaWAN Gateway

* [Product details](https://www.dragino.com/products/lora-lorawan-gateway/item/228-lps8v2.html)

**Note:** This was tested with the following firmware version: `1.0.1`

### Configure Packet Forwarder

In the Dragino web-interface, you must configure the Packet Forwarder such
that it forwards to `localhost` on port `1700`.

By default, the web-interface can be accessed by entering the following URL
in your browser: `http://GATEWAY-IP-ADDRESS` (replace `GATEWAY-IP-ADDRESS`
by the actual IP address of your gateway). The default credentials are
`root` / `dragino`.

* In the **LoRaWAN** menu, click **LoRaWAN -- Semtech UDP**
* Configure the following settings:
  * **Service Provider:** _Custom / Private LoRaWAN_
  * **Server Address:** `localhost`
  * **Uplink Port:** `1700`
  * **Downlink Port:** `1700`
* Click **Save & Apply**

### Install ChirpStack MQTT Forwarder

#### SSH login

First you must login into the gateway using SSH:

```bash
ssh root@GATEWAY-IP-ADDRESS
```

The default password is `dragino`.

#### Download .deb

Use the following commands to download the latest version of the
ChirpStack MQTT Forwarder package:

```bash
cd /tmp
wget https://artifacts.chirpstack.io/downloads/chirpstack-mqtt-forwarder/chirpstack-mqtt-forwarder_$MQTT_FORWARDER_VERSION_linux_armhf.deb
```

#### Install .deb

Use the `dpkg` package-manager to install the downloaded package. Example:

```bash
dpkg -i chirpstack-mqtt-forwarder_$MQTT_FORWARDER_VERSION_linux_armhf.deb
```

#### Configuration

To connect the ChirpStack MQTT Forwarder to your MQTT broker, you must update
the ChirpStack MQTT Forwarder configuration file. This file is located at:
`/etc/chirpstack-mqtt-forwarder/chirpstack-mqtt-forwarder.toml`.

#### (Re)start and stop commands

Use the following commands to (re)start and stop the ChirpStack MQTT Forwarder service:

```bash
# start
systemctl start chirpstack-mqtt-forwarder

# stop
systemctl stop chirpstack-mqtt-forwarder

# restart
systemctl restart chirpstack-mqtt-forwarder
```


## LG308(N) LoRaWAN Gateway

* [Product details](https://www.dragino.com/products/lora-lorawan-gateway/item/229-lg308n.html)

Please see LPS8N Indoor LoRaWAN Gateway instructions.
