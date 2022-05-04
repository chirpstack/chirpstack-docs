# Getting started

These steps describe how to get started with ChirpStack Gateway OS **after** you
have installed ChirpStack Gateway OS on your gateway.

**Important:** The **chirpstack-gateway-os-full** image will setup the PostgreSQL
on its first boot. This could take a couple of minutes and during this time,
the gateway will be less responsive!

## Connecting

After booting the gateway, you need to login using SSH. In case the gateway
running the ChirpStack Gateway OS supports Wi-Fi, then it will be configured
as access-point with the name `ChirpStackAP` and password `ChirpStackAP`.
Once connected with `ChirpStackAP` the IP of the gateway is `192.168.0.1`.

If you are connected using ethernet, then it uses DHCP to obtain an IP address.
Many internet routers provide a web-interface with the IP addresses of connected
devices.

If the IP of your gateway is `192.168.0.1`:

```bash
ssh admin@192.168.0.1
```

The default username is `admin`, the default password is `admin`.

This will prompt the following message:

```text
   ________    _           _____ __             __     _     
  / ____/ /_  (_)________ / ___// /_____ ______/ /__  (_)___ 
 / /   / __ \/ / ___/ __ \\__ \/ __/ __ `/ ___/ //_/ / / __ \
/ /___/ / / / / /  / /_/ /__/ / /_/ /_/ / /__/ ,< _ / / /_/ /
\____/_/ /_/_/_/  / .___/____/\__/\__,_/\___/_/|_(_)_/\____/ 
                 /_/

Documentation and copyright information:
> www.chirpstack.io

Commands:
> sudo gateway-config  - configure the gateway
> sudo monit status    - display service monitor
```

## Configuration

Execute the `sudo gateway-config` to configure the concentrator shield model
and the channel-configuration that the gateway must use.

### Base image

After the board and channel-plan have been configured, you must update the
ChirpStack Gateway Bridge configuration, such that it connects to your MQTT
broker. This can be done using the **Edit ChirpStack Gateway Bridge config**
option in the `gateway-config` menu.

Use the **Edit configuration file** option to edit the configuration file or
the **MQTT connection wizzard**. In case you are using (client-)certificate
authentication / authorization, the latter is recommended as it allows you
do directly paste the certificate files.

## Full image

Unlike the **chirpstack-gateway-os-base** image, you **should not** update the
ChirpStack Gateway Bridge configuration. It is configured to point to the MQTT broker
which comes with the **chirpstack-gateway-os-full** image.

If using the **chirpstack-gateway-os-full** image and after configuring the
channel-plan, the gateway will be automatically added to ChirpStack. In case
the `gateway-config` displays an error during this step, it is recommended to
wait ~ 30 seconds and try again. On first boot, it might take some time to
setup the database (which happens on the background) and during this process
it is normal to see an error message.

### Access ChirpStack web-interface

To access the web-interface, you need the IP address of the Raspberry Pi. If it
is still in Wi-Fi access-point mode, the IP will be `192.168.0.1`, else you
can usually retrieve the assigned IP address from your internet modem.

Assuming its IP address is `192.168.0.1`, enter the following address in your
browser to access the web-interface: `http://192.168.0.1:8080`.

The default login credentials are: `admin` / `admin`.

### Create device profile

After login, the first thing to do is creating a device profile. A device
profile stores the properties for a certain device type (e.g. brand, model, ...)
that are important for ChirpStack to correctly handle your devices. To do
so, click **Device profiles** and then **Add device profile**.

**Note:** It is important that the _Region_ field matches the region that you
have configured using `gateway-config`.

### Create application

To group simmilar devices together, you need to create an application. This can
be done by clicking **Applications** and then **Add application**.

### Create first device

After creating the application, you are directly redirected to the created
application, with the _Devices_ tab enabled. You can add your device by clicking
the **Add device** button.

## Important to know

### SD Card wearout

Although ChirpStack will try to minimize the number of database writes, there
will be regular writes to the SD Card (PostgreSQL and Redis snapshots).
According to [Is it true that a SD/MMC Card does wear levelling with its own controller?](https://electronics.stackexchange.com/questions/27619/is-it-true-that-a-sd-mmc-card-does-wear-levelling-with-its-own-controller)
it might make a difference which SD Card brand you use.
