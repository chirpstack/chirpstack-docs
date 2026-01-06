# Multitech

<!-- toc -->

## Multitech Conduit

* [Product detail page](https://www.multitech.com/brands/multiconnect-conduit)
* [Product documentation page](http://www.multitech.net/developer/products/multiconnect-conduit-platform/)

**Note:** This was tested with the following firmware version: `6.0.0`.

### Configure Packet Forwarder

#### Web-interface log in

The mPower web-interface can be accessed by entering `https://IP-ADDRESS/` into
your browser. As the web-interface uses a self-signed certificate, your browser
will probably raise a warning.

#### Option 1 - Packet-forwarder configuration

1. In the left menu, click _LoRaWAN&reg;_
2. Under _LoRa mode_, select **PACKET FORWARDER**
3. Under _LoRa Packet Forwarder Configuration_ enter / select the following settings:
	* Network Settings
		* Network: **Manual**
		* Channel Plan: The desired channel-plan
	* Server Settings
		* Server Address: **127.0.0.1**
		* Upstream Port: **1700**
		* Downstream Port: **1700**
4. Click _Submit_, and then _Save and Appy_ in the left menu.

After completing these steps, you should see _RUNNING_ under the Packet Forwarder Status.

#### Option 2 - ChirpStack Concentratord

Please refer to the ChirpStack Concentratord [installation instructions](../../chirpstack-concentratord/installation/multitech.md)
for installing the ChirpStack Concentratord on the Multitech Conduit.

### Install ChirpStack MQTT Forwarder

#### Enable SSH

In order to install ChirpStack packages on the gateway, you must first enable SSH
in the web-interface.

1. In the left menu, click _Administration_ and then _Access configuration_. 
2. Under _SSH Settings_, make sure this option is **Enabled**.
3. In case of changes, click _Submit_ and then _Save and Apply_ in the left menu.

#### SSH log in

To log in into your gateway, use the following command:

```bash
ssh USERNAME@IP-ADDRESS
```

Where `USERNAME` is the username that you use to gain access to the web-interface
of the gateway and `IP-ADDRESS` with the IP address of the gateway.

#### Download IPK

Use the following commands to download the latest version of the
`chirpstack-mqtt-forwarder` package:

```bash
cd /tmp
wget https://artifacts.chirpstack.io/downloads/chirpstack-mqtt-forwarder/vendor/multitech/conduit/chirpstack-mqtt-forwarder_$MQTT_FORWARDER_VERSION-r1_arm926ejste.ipk
```

#### Install IPK

Use the `opkg` package-manager to install the downloaded package. Example:

```bash
sudo opkg install chirpstack-mqtt-forwarder_$MQTT_FORWARDER_VERSION-r1_arm926ejste.ipk
```

#### Configuration

The Multitech Conduit has two accessory ports (AP1 and AP2) which can be used
simultaneously, for example AP1 for EU868 and AP2 for 2.4GHz. For this reason
it is possible to start two instances of the ChirpStack MQTT Forwarder.

To connect the ChirpStack MQTT Forwarder to your MQTT broker, you must update
the ChirpStack MQTT Forwarder configuration file(s). These are located at:

* **AP1**: `/var/config/chirpstack-mqtt-forwarder/ap1/chirpstack-mqtt-forwarder.toml`
* **AP2**: `/var/config/chirpstack-mqtt-forwarder/ap2/chirpstack-mqtt-forwarder.toml`

#### Enabling ChirpStack MQTT Forwarder

To enable the ChirpStack MQTT Forwarder instance or instances, you must copy
the [Monit](https://mmonit.com/monit/) configuration file(s) for the
corresponding APs:

```bash
# To enable AP1
sudo cp /var/config/chirpstack-mqtt-forwarder/ap1/examples/chirpstack-mqtt-forwarder-ap1.monit /etc/monit.d/

# To enable AP2
sudo cp /var/config/chirpstack-mqtt-forwarder/ap2/examples/chirpstack-mqtt-forwarder-ap2.monit /etc/monit.d/

# Reload Monit
sudo monit reload
```

#### (Re)start and stop commands

Use the following commands to (re)start and stop the ChirpStack MQTT Forwarder
services:

```bash
# Status
sudo monit summary

# Start
sudo monit start chirpstack-mqtt-forwarder-[ap1|ap2]

# Restart
sudo monit restart chirpstack-mqtt-forwarder-[ap1|ap2]

# Stop
sudo monit stop chirpstack-mqtt-forwarder-[ap1|ap2]
```

#### Log output

To view the ChirpStack MQTT Forwarder log output, use the following command:

```bash
tail -f -n 100 /var/log/messages |grep chirpstack-mqtt-forwarder
```

## Multitech Conduit AP

* [Product detail page](https://www.multitech.com/brands/multiconnect-conduit-ap)
* [Product documentation page](https://www.multitech.net/developer/products/multiconnect-conduit-access-point/)

**Note:** This was tested with the following firmware version: `6.0.1`.

### Configure Packet Forwarder

#### Web-interface log in

The mPower web-interface can be accessed by entering `https://IP-ADDRESS/` into
your browser. As the web-interface uses a self-signed certificate, your browser
will probably raise a warning.

#### Option 1 - Packet-forwarder configuration

1. In the left menu, click _LoRaWAN&reg;_
2. Under _LoRa mode_, select **PACKET FORWARDER**
3. Under _LoRa Packet Forwarder Configuration_ enter / select the following settings:
	* Network Settings
		* Network: **Manual**
		* Channel Plan: The desired channel-plan
	* Server Settings
		* Server Address: **127.0.0.1**
		* Upstream Port: **1700**
		* Downstream Port: **1700**
4. Click _Submit_, and then _Save and Appy_ in the left menu.

After completing these steps, you should see _RUNNING_ under the Packet Forwarder Status.

#### Option 2 - ChirpStack Concentratord

Please refer to the ChirpStack Concentratord [installation instructions](../../chirpstack-concentratord/installation/multitech.md)
for installing the ChirpStack Concentratord on the Multitech Conduit AP.

### Install ChirpStack MQTT Forwarder

#### Enable SSH

In order to install ChirpStack packages on the gateway, you must first enable SSH
in the web-interface.

1. In the left menu, click _Administration_ and then _Access configuration_. 
2. Under _SSH Settings_, make sure this option is **Enabled**.
3. In case of changes, click _Submit_ and then _Save and Apply_ in the left menu.

#### SSH log in

To log in into your gateway, use the following command:

```bash
ssh USERNAME@IP-ADDRESS
```

Where `USERNAME` is the username that you use to gain access to the web-interface
of the gateway and `IP-ADDRESS` with the IP address of the gateway.

#### Download IPK

Use the following commands to download the latest version of the
`chirpstack-mqtt-forwarder` package:

```bash
cd /tmp
wget https://artifacts.chirpstack.io/downloads/chirpstack-mqtt-forwarder/vendor/multitech/conduit_ap/chirpstack-mqtt-forwarder_$MQTT_FORWARDER_VERSION-r1_arm926ejste.ipk
```

#### Install IPK

Use the `opkg` package-manager to install the downloaded package. Example:

```bash
sudo opkg install chirpstack-mqtt-forwarder_$MQTT_FORWARDER_VERSION-r1_arm926ejste.ipk
```

#### Configuration

To connect the ChirpStack MQTT Forwarder to your MQTT broker, you must update
the ChirpStack MQTT Forwarder configuration file. This file is located at:
`/var/config/chirpstack-mqtt-forwarder/chirpstack-mqtt-forwarder.toml`.

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

## Multitech Conduit AP3

* [Product detail page](https://www.multitech.com/brands/conduit-ap-300)
* [Product documentation page](https://www.multitech.net/developer/products/multiconnect-conduit-access-point/)

**Note:** This was tested with the following firmware version: `6.3.0`.

### Configure Packet Forwarder

#### Web-interface log in

The mPower web-interface can be accessed by entering `https://IP-ADDRESS/` into
your browser. As the web-interface uses a self-signed certificate, your browser
will probably raise a warning.

#### Option 1 - Packet-forwarder configuration

1. In the left menu, click _LoRaWAN&reg;_
2. Under _LoRa mode_, select **PACKET FORWARDER**
3. Under _LoRa Packet Forwarder Configuration_ enter / select the following settings:
	* Network Settings
		* Network: **Manual**
		* Channel Plan: The desired channel-plan
	* Server Settings
		* Server Address: **127.0.0.1**
		* Upstream Port: **1700**
		* Downstream Port: **1700**
4. Click _Submit_, and then _Save and Appy_ in the left menu.

After completing these steps, you should see _RUNNING_ under the Packet Forwarder Status.

#### Option 2 - ChirpStack Concentratord

Please refer to the ChirpStack Concentratord [installation instructions](../../chirpstack-concentratord/installation/multitech.md)
for installing the ChirpStack Concentratord on the Multitech Conduit AP.

### Install ChirpStack MQTT Forwarder

#### Enable SSH

In order to install ChirpStack packages on the gateway, you must first enable SSH
in the web-interface.

1. In the left menu, click _Administration_ and then _Access configuration_. 
2. Under _SSH Settings_, make sure this option is **Enabled**.
3. In case of changes, click _Submit_ and then _Save and Apply_ in the left menu.

#### SSH log in

To log in into your gateway, use the following command:

```bash
ssh USERNAME@IP-ADDRESS
```

Where `USERNAME` is the username that you use to gain access to the web-interface
of the gateway and `IP-ADDRESS` with the IP address of the gateway.

#### Download IPK

Use the following commands to download the latest version of the
`chirpstack-mqtt-forwarder` package:

```bash
cd /tmp
wget https://artifacts.chirpstack.io/downloads/chirpstack-mqtt-forwarder/vendor/multitech/conduit_ap3/chirpstack-mqtt-forwarder_$MQTT_FORWARDER_VERSION-r1_mtcap3.ipk
```

#### Install IPK

Use the `opkg` package-manager to install the downloaded package. Example:

```bash
sudo opkg install chirpstack-mqtt-forwarder_$MQTT_FORWARDER_VERSION-r1_mtcap3.ipk
```

#### Configuration

To connect the ChirpStack MQTT Forwarder to your MQTT broker, you must update
the ChirpStack MQTT Forwarder configuration file. This file is located at:
`/var/config/chirpstack-mqtt-forwarder/chirpstack-mqtt-forwarder.toml`.

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
