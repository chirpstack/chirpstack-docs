# Getting started

<!-- toc -->


These steps describe how to get started with ChirpStack Gateway OS _after_ you
have installed ChirpStack Gateway OS on your gateway (e.g. by writing the image
to a SD card).

**Important:** After the first boot, the gateway might reboot automatically to
apply some changes. The **Full** image will setup the PostgreSQL database on
its first boot. This could take a couple of minutes and during this time, the
gateway will be less responsive.

## Connecting

### Obtaining gateway IP address

#### Ethernet

After your gateway has booted, it will use DHCP to obtain an IP address. 
Many internet routers provide a web-interface where you will find the IP
addresses of connected devices.

#### Wi-Fi

If your gateway has Wi-Fi support, then the ChirpStack Gateway OS will create
an access-point to which you can connect with the name `ChirpStackAP-XXXXXX` and 
password `ChirpStackAP`.

Once connected with `ChirpStackAP` the IP of the gateway is `192.168.0.1`.

### Web-interface

The ChirpStack Gateway OS provides a web-interface. You can connect to the
web-interface by entering `http://GATEWAY-IP-ADDRESS/` in your browser. For
example if the IP address of your gateway is `192.168.0.1`, then you need to
enter `http://192.168.0.1`.

Your browser might show you a warning about a self-signed certificate.
Click proceed.

The default credentials are:

* Username: `root`
* Password: _(not set)_

## Set password

Once you have logged in, it is recommended to set a password for the _root_
user. In the left menu click **System**, then click **Administration**.

This will provide you with a form to set the _Router Password_.

## Configure Wi-Fi (optional)

If your gateway has Wi-Fi support, then follow the next steps to connect over
Wi-Fi (which will also disable the access-point mode).

1. In the left menu click **Network**, then click **Wireless**.
2. Click the **Scan** button, after a couple of seconds this will show the available Wi-Fi networks.
3. Click the **Join Network** button of the network you want to join.
4. In the Joining Network form, enter the following values and click **Submit**:
   * _Replace wireless configuration_: **Selected**
   * _WPA passphrase_: Enter the password of the selected Wi-Fi network
   * _Create / Assign firewall-zone_: Select **lan [lan: ] [wwan: ]**
5. Click **Save**
6. Click **Save & Apply**

After your gateway has connected to the Wi-Fi network, you need to obtain the
IP address again and re-connect. Please see the _Connecting_ steps.

## Flash concentrator firmware (optional)

In case you have a USB based concentrator module (USB based SX1302/3 module
or 2.4GHz module), you must flash the MCU firmware of the module before
it can be used.

To upgrade the MCU firmware, click **System** and then **Custom Commands**.
Click the **Run** button for the firmware upgrade option.

## Configure Concentratord

To configure the concentrator, click **ChirpStack** and then **Concentratord**
in the left menu. The important settings that you must configure are:

* Global configuration
   * _Enabled chipset_: Set this to the chipset matching your gateway
* SX1301, SX1302/SX1303, 2.4GHz tabs (matching the _Enabled chipset_)
   * _Shield model_: Select the shield model.
   * _Channel-plan_: Select the channel-plan. This will show an error if the channel-plan is not supported by the gateway or if the channel-plan has not been implemented yet.

If all fields are correctly configured, click **Save & Apply**.

Notes:

* **Save & Apply** will automatically configure the ChirpStack MQTT Forwarder MQTT prefix.
* If you are using the **Full** image, then this will also enable the region in ChirpStack.

### Confirm configuration

Once the Concentratord has been properly configured, the footer of the
web-interface should change from:

_Gateway ID: could not read gateway_id_

to (_0102030405060708_ replaced by the ID of your gateway):

_Gateway ID: 0102030405060708_

**Note:** It might take a couple of seconds for the Concentratord to initialize.
You might need to refresh the page a couple of times.

## Configure ChirpStack MQTT Forwarder (optional)

**Note:** If you are using the **Full** image, then ChirpStack MQTT Forwarder
is already configured to forward to the local ChirpStack Network Server
instance. In this case, no configuration is required.

To configure the ChirpStack MQTT Forwarder to forward to a (remote) MQTT broker,
click **ChirpStack** and then **MQTT Forwarder** in the left menu.

The important settings that you must configure are:

* _Topic prefix_: This must match the topic prefix as configured by ChirpStack.
  Note that this will be automatically set after the _Concentratord_
  configuration.
* _Use JSON_: Only enable this option for debugging purposes.
* _Server_: The address of the MQTT broker, e.g. `tcp://example.com:1883` or `ssl://example.com:8883` (TLS).

Depending the configuration of the MQTT broker, you might also want to edit
the _Username_ / _Password_ fields or _CA certificate_ / _TLS certificate_ and
_TLS key-file_ fields.

To apply the changes, click **Save & Apply**.

## Configure ChirpStack UDP Forwarder (optional)

The ChirpStack UDP Forwarder makes it possible to connect your gateway to
one or multiple Network Servers compatible with the Semtech UDP
Packet-Forwarder [protocol](https://github.com/Lora-net/packet_forwarder/blob/master/PROTOCOL.TXT).
This makes it possible to connect your gateway to multiple Network Servers
simultaneously.

To configure the ChirpStack UDP Forwarder, click **ChirpStack** and then
**UDP Forwarder** in the left menu.

Here you can configure multiple servers to which the ChirpStack UDP Forwarder
will forward simultaneously. Per server you must configure:

* _Server_: The hostname or IP address + port of the server, e.g. `example:1700`.

To apply the changes, click **Save & Apply**.

## Using ChirpStack LoRaWAN Network Server

If you have properly configured your gateway and if you have installed the
**Full** image (see [Image types](./image-types.md)), then you are ready
to use the ChirpStack LoRaWAN Network Server.

To open ChirpStack, click **Applications** in the left menu, and then click
the **ChirpStack** logo. This will open the ChirpStack web-interface in a
new window.

Please refer to the [Connecting a device](../guides/connect-device.md) guide
for connecting your first device.

**Note:** ChirpStack has been configured to allow gateway data from unknown
gateways, therefore it is not required to add your gateway to ChirpStack first.