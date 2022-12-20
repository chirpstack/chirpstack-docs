# Multitech

<!-- toc -->

## Multitech Conduit

* [Product detail page](https://www.multitech.com/brands/multiconnect-conduit)
* [Product documentation page](http://www.multitech.net/developer/products/multiconnect-conduit-platform/)

**Note:** This was tested with the following firmware version: `6.0.0`.

### Supported cards

| mCard | Concentratord IPK | Concentratord configuration directory |
| ----- | ----------------- | ----------------- |
| MTAC-003E00 | [chirpstack-concentratord-sx1302_{{concentratord_version}}_r1_arm926ejste.ipk](https://artifacts.chirpstack.io/downloads/chirpstack-concentratord/vendor/multitech/conduit/chirpstack-concentratord-sx1302_{{concentratord_version}}-r1_arm926ejste.ipk) | `/var/config/chirpstack-concentratord-sx1302` |
| MTAC-003U00 | [chirpstack-concentratord-sx1302_{{concentratord_version}}_r1_arm926ejste.ipk](https://artifacts.chirpstack.io/downloads/chirpstack-concentratord/vendor/multitech/conduit/chirpstack-concentratord-sx1302_{{concentratord_version}}-r1_arm926ejste.ipk) | `/var/config/chirpstack-concentratord-sx1302` |
| MTAC-LORA-H-915 | [chirpstack-concentratord-sx1301_{{concentratord_version}}_r1_arm926ejste.ipk](https://artifacts.chirpstack.io/downloads/chirpstack-concentratord/vendor/multitech/conduit/chirpstack-concentratord-sx1301_{{concentratord_version}}-r1_arm926ejste.ipk) | `/var/config/chirpstack-concentratord-sx1301` |
| MTAC-LORA-H-868 | [chirpstack-concentratord-sx1301_{{concentratord_version}}_r1_arm926ejste.ipk](https://artifacts.chirpstack.io/downloads/chirpstack-concentratord/vendor/multitech/conduit/chirpstack-concentratord-sx1301_{{concentratord_version}}-r1_arm926ejste.ipk) | `/var/config/chirpstack-concentratord-sx1301` |
| MTAC-LORA-2G4 | [chirpstack-concentratord-2g4_{{concentratord_version}}_r1_arm926ejste.ipk](https://artifacts.chirpstack.io/downloads/chirpstack-concentratord/vendor/multitech/conduit/chirpstack-concentratord-2g4_{{concentratord_version}}-r1_arm926ejste.ipk) | `/var/config/chirpstack/concentratord-2g4` |

### Disable internal packet forwarder

In order to install the ChirpStack Concentratord, you must disable any packet
forwarder provided by Multitech.

#### Web-interface log in

The mPower web-interface can be accessed by entering `https://IP-ADDRESS/` into
your browser. As the web-interface uses a self-signed certificate, your browser
will probably raise a warning.

#### Disable packet-forwarder

1. In the left menu, click _LoRaWAN&reg;_
2. Under _LoRa mode_, select **DISABLED**
3. Click _Submit_, and then _Save and Appy_ in the left menu.

### Install ChirpStack Concentratord

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

You must select the correct Concentratord IPK package for the card(s) that you
have installed in your Multitech Conduit gateway. Please refer to the _Supported cards_
table and copy the link of the corresponding IPK package.

**Note:** Multiple ChirpStack Concentratord variants can be installed
simultaneously, e.g. if you have a MTAC-003E00 installed in AP1 and a
MTAC-LORA-2G4 in AP2.

Use the following command to download the IPK on your gateway:

```bash
cd /tmp
wget <Concentratord IPK URL>
```

#### Install IPK

Use the `opkg` package-manager to install the downloaded package. Example:

```bash
# E.g. sudo opkg install chirpstack-concentratord-sx1302_{{concentratord_version}}_r1_arm926ejste.ipk
sudo opkg install <Concentratord IPK>
```

#### Configuration

On first start, the ChirpStack Concentratord will automatically detect the hardware
region of your card. After the first start, you will the Concentratord either
in the `ap1` or `ap2` sub-directory of the Concentratord configuration directory.

To start the ChirpStack Concentratord, use the following command:

```bash
sudo /etc/init.d/chirpstack-concentratord-<variant>-<port> start
```

* `<variant>` can be: `sx1301`, `sx1302` or `2g4`.
* `<port>` can be: `ap` or `ap2`.

Example:

```bash
# To start the MTAC-003E00 module in AP1
sudo /etc/init.d/chirpstack-concentratord-sx1302-ap1 start

# To start the MTAC-LORA-2G4 module in AP2
sudo /etc/init.d/chirpstack-concentratord-2g4-ap2 start
```

#### Enabling ChirpStack Concentratord

To automatically enable ChirpStack Concentratord on boot, you must copy
the [Monit](https://mmonit.com/monit/) configuration file(s) for the
corresponding APs:

```bash
# Copy Monit configuration file(s)
sudo cp /var/config/chirpstack-concentratord-<variant>/chirpstack-concentratord-<variant>-<port>.monit /etc/monit.d

# Reload Monit
sudo monit reload
```

Example:

```bash
# To start the MTAC-003E00 module in AP1 on boot
sudo cp /var/config/chirpstack-concentratord-sx1302/chirpstack-concentratord-sx1302-ap1.monit /etc/monit.d

# To start the MTAC-LORA-2G4 module in AP2 on boot
sudo cp /var/config/chirpstack-concentratord-2g4/chirpstack-concentratord-2g4-ap2.monit /etc/monit.d

# Reload Monit
sudo monit reload
```

#### (Re)start and stop commands

Use the following commands to (re)start and stop the ChirpStack Concentratord
service(s):

```bash
# Status
sudo monit summary

# start
monit start chirpstack-concentratord-<variant>-<port>

# stop
monit stop chirpstack-mqtt-forwarder-<variant>-<port>

# restart
monit restart chirpstack-mqtt-forwarder-<variant>-<port>
```


#### Log output

To view the ChirpStack Concentratord log output, use the following command:

```bash
tail -f -n 100 /var/log/messages |grep chirpstack-concentratord
```

## Multitech Conduit AP

* [Product detail page](https://www.multitech.com/brands/multiconnect-conduit-ap)
* [Product documentation page](https://www.multitech.net/developer/products/multiconnect-conduit-access-point/)

**Note:** This was tested with the following firmware version: `6.0.1`.


### Disable internal packet forwarder

In order to install the ChirpStack Concentratord, you must disable any packet
forwarder provided by Multitech.

#### Web-interface log in

The mPower web-interface can be accessed by entering `https://IP-ADDRESS/` into
your browser. As the web-interface uses a self-signed certificate, your browser
will probably raise a warning.

#### Disable packet-forwarder

1. In the left menu, click _LoRaWAN&reg;_
2. Under _LoRa mode_, select **DISABLED**
3. Click _Submit_, and then _Save and Appy_ in the left menu.

### Install ChirpStack Concentratord

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
`chirpstack-concentratord` package:

```bash
cd /tmp
wget https://artifacts.chirpstack.io/downloads/chirpstack-concentratord/vendor/multitech/conduit-ap/chirpstack-concentratord_{{concentratord_version}}-r1_arm926ejste.ipk
```

#### Install IPK

Use the `opkg` package-manager to install the downloaded package. Example:

```bash
sudo opkg install chirpstack-concentratord_{{concentratord_version}}-r1_arm926ejste.ipk
```

#### Configuration

The first time ChirpStack Concentratord is started, it will automatically
detect the region of the gateway, and copy the configuration.

To start the Concentratord, use the following command:

```bash
sudo monit start chirpstack-concentratord
```

Configuration files can be found at:

* `/var/config/chirpstack-concentratord/concentratord.toml` - Concentratord configuration
* `/var/config/chirpstack-concentratord/channels.toml` - Channel configuration
* `/var/config/chirpstack-concentratord/examples` - Directory containing example configuration

#### (Re)start and stop commands

Use the following commands to (re)start and stop the ChirpStack Concentratord
service:

```bash
# Status
sudo monit summary

# start
monit start chirpstack-concentratord

# stop
monit stop chirpstack-concentratord

# restart
monit restart chirpstack-concentratord
```

#### Log output

To view the ChirpStack Concentratord log output, use the following command:

```bash
tail -f -n 100 /var/log/messages |grep chirpstack-concentratord
```
