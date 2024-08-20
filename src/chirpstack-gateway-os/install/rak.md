# RAK

<!-- toc -->

**Note:** These are custom firmware images that are not supported by RAK and might
might void your warranty. Use at your own risk.

## Downloads

Please use one of the download options below to download the latest
(v{{gateway_os_version}}) ChirpStack Gateway OS image.

| Model | Factory image | Sysupgrade image |
| ----- | ------------- | ---------------- |
| RAK7267 | | [Download](https://artifacts.chirpstack.io/downloads/chirpstack-gateway-os/{{gateway_os_version}}/rak/ramips/mt76x8/chirpstack-gateway-os-{{gateway_os_version}}-rak7267-ramips-mt76x8-rakwireless_rak636-squashfs-sysupgrade.bin) |
| RAK7268v2 | | [Download](https://artifacts.chirpstack.io/downloads/chirpstack-gateway-os/{{gateway_os_version}}/rak/ramips/mt76x8/chirpstack-gateway-os-{{gateway_os_version}}-rak7268v2-ramips-mt76x8-rakwireless_rak636-squashfs-sysupgrade.bin) |
| RAK7289v2 | | [Download](https://artifacts.chirpstack.io/downloads/chirpstack-gateway-os/{{gateway_os_version}}/rak/ramips/mt76x8/chirpstack-gateway-os-{{gateway_os_version}}-rak7289v2-ramips-mt76x8-rakwireless_rak636-squashfs-sysupgrade.bin) |
| RAK7391 | [Download](https://artifacts.chirpstack.io/downloads/chirpstack-gateway-os/{{gateway_os_version}}/rak/bcm27xx/bcm2709/chirpstack-gateway-os-{{gateway_os_version}}-rak7391-bcm27xx-bcm2709-rpi-2-squashfs-factory.img.gz) | [Download](https://artifacts.chirpstack.io/downloads/chirpstack-gateway-os/{{gateway_os_version}}/rak/bcm27xx/bcm2709/chirpstack-gateway-os-{{gateway_os_version}}-rak7391-bcm27xx-bcm2709-rpi-2-squashfs-sysupgrade.img.gz) |

## Installation

### RAK72XX

#### WisGateOS2 to Gateway OS

This replaces the WisGateOS2 firmware of the gateway with the ChirpStack GatewayOS
OS.

```bash
cd /tmp

# Download image
# Note: Replace this URL with the correct image URL from the Downloads!
wget https://artifacts.chirpstack.io/downloads/chirpstack-gateway-os/{{gateway_os_version}}/rak/ramips/mt76x8/chirpstack-gateway-os-{{gateway_os_version}}-rakXXXX-ramips-mt76x8-rakwireless_rak636-squashfs-sysupgrade.bin

# Flash image
# Note: Replace filename with the downloaded image!
sysupgrade -n chirpstack-gateway-os-{{gateway_os_version}}-rakXXXX-ramips-mt76x8-rakwireless_rak636-squashfs-sysupgrade.bin
```

To avoid conflicting configuration, you must use the `-n` flag. This will mean
all data stored on the gateway will be erased.

**Note:** This will take a couple of minutes. During the process the gateway
might disconnect. After rebooting, the gateway might have a different IP 
address. Do not power-cycle the gateway!

#### Gateway OS to WisGateOS2

This restores the WisGateOS2 firmware of a RAK gateway that previously has been
flashed with a ChirpStack Gateway OS image. **Note:** this will not install the
latest WisGateOS2 image version. Once you have restored the WisGateOS2 firmware
you should upgrade your gateway to the latest WisGateOS2 image.

##### Set `compat_version`

You must set the `compat_version` parameter to `2.1.4`, or else the upgrade
will fail with:

> The device is supported, but this image is incompatible for sysupgrade based on the image version (1.0->2.1.4)

To set this parameter, execute the following command:


```bash
uci set system.@system[0].compat_version='2.1.4'
uci commit
```

##### Download and flash

```bash
cd /tmp

# Download restore image
wget https://artifacts.chirpstack.io/downloads/chirpstack-gateway-os/rak/wisgateos-ramips-mt76x8-rakwireless_rak636-squashfs-sysupgrade.bin

# Flash restore image
sysupgrade -n wisgateos-ramips-mt76x8-rakwireless_rak636-squashfs-sysupgrade.bin
```

To avoid conflicting configuration, you must use the `-n` flag. This will mean
all data stored on the gateway will be erased.

**Note:** This will take a couple of minutes. During the process the gateway
might disconnect. After rebooting, the gateway might have a different IP 
address. Do not power-cycle the gateway!

##### Upgrading to latest WisGateOS2

Once the WisGateOS2 image is restored, you should see the WisGateOS2 web-interface
which will guide you through the gateway configuration. Once setup, you should
update the firmware to the latest WisGateOS2 version! You can download the
latest WisGateOS2 firmware image from the [RAK Download Center](https://downloads.rakwireless.com/#LoRa/WisGateOS2/).

### RAK7391

For instructions on how to flash an image to the Compute Module 4, please refer to
[Flash an image to a Compute Module](https://www.raspberrypi.com/documentation/computers/compute-module.html#flash-an-image-to-a-compute-module).

## First usage

After flashing the ChirpStack Gateway OS image, follow the [Getting started](../getting-started.md)
documentation to setup your gateway.
