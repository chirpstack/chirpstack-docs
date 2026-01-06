# Seeed

<!-- toc -->

**Note:** These are custom firmware images that are not supported by Seeed and might
might void your warranty. Use at your own risk.

## Downloads

Please use one of the download options below to download the latest
(v%GATEWAY_OS_VERSION) ChirpStack Gateway OS image.

| Model | Factory image | Sysupgrade image |
| ----- | ------------- | ---------------- |
| SenseCAP M2 | | [Download](https://artifacts.chirpstack.io/downloads/chirpstack-gateway-os/%GATEWAY_OS_VERSION/seeed/ramips/mt76x8/chirpstack-gateway-os-%GATEWAY_OS_VERSION-sensecap-m2-ramips-mt76x8-sensecap_wm7628n-squashfs-sysupgrade.bin) |

## Installation

### SenseCAP M2

#### SenseCAP M2 to Gateway OS

This replaces the (open-source) SenseCAP M2 firmware of the gateway with the
Chirpstack Gateway OS.

1. First download the latest ChirpStack Gateway OS image for the SenseCAP M2 (see above table).
2. In the SenseCAP M2 web-interface, nativate to **System > Backup / Flash Firmware**.
3. Click the **Flash image...** button at the bottom of the page and select the ChirpStack Gateway OS image and click **Upload**.
4. Depending the SenseCAP M2 firmware, you might get the option **Keep settings and retain the current configuration**, if this
   is the case, then de-select this option and **Continue**.
5. Wait until the firmware is flashed, this will take some minutes.

If you did not see the **Keep settings and retain...** option (step 4), then
the previous credentials (and some other settings) will be retained. It is advised
to perform a factory reset after logging into the gateway to avoid conflicting
configuration. The ChirpStackAP-xxxxxx access-point might not work until the
factory reset is performed. A factory-reset can be performed in the web-interface
under **System > Backup / Flash Firmware** by clicking the **Perform reset**
button.


#### Gateway OS to SenseCAP M2

This replaces the Chirpstack Gateway OS with the SenseCAP M2 open-source firmware.

1. First download the latest SenseCAP M2 open-source firmware from [https://github.com/Seeed-Solution/LoRa_Gateway_OpenWRT/releases](https://github.com/Seeed-Solution/LoRa_Gateway_OpenWRT/releases).
2. In the ChirpStack Gateway OS web-interface, navigate to **System > Backup / Flash Firmware**.
3. Click the **Flash image...** button at the bottom of the page and select the SenseCAP M2 image and click **Upload**.
4. De-select the **Keep settings and retain the current configuration** option and **Continue**.
5. Wait until the firmware is flashed, this will take some minutes.
6. Continue with the [Login into Console](https://wiki.seeedstudio.com/flash_opensource_firmware_to_m2_gateway/#login-into-console) documentation.

## First usage

After flashing the ChirpStack Gateway OS image, follow the [Getting started](../getting-started.md)
documentation to setup your gateway.
