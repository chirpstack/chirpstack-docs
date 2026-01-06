# Dragino

<!-- toc -->

**Note:** These are custom firmware images that are not supported by Dragino and might
might void your warranty. Use at your own risk.

## Downloads

Please use one of the download options below to download the latest
(v$GATEWAY_OS_VERSION) ChirpStack Gateway OS image.

| Model | Factory image | Sysupgrade image |
| ----- | ------------- | ---------------- |
| Dragino LPS8N | | [Download](https://artifacts.chirpstack.io/downloads/chirpstack-gateway-os/$GATEWAY_OS_VERSION/dragino/ath79/generic/chirpstack-gateway-os-$GATEWAY_OS_VERSION-lps8n-ath79-generic-dragino_lps8n-squashfs-sysupgrade.bin) |

## Installation

### LPS8N

**Important note:** The LPS8N device has limited memory. When running the
ChirpStack Gateway OS image on the gateway, it might run out of memory
when flashing a new firmware. To avoid this, it is recommended to stop
all `chirpstack-*` services first. This can be done in the web-interface
under **System > Startup**, by clicking the **Stop** button for each
`chirpstack-` service.

#### LPS8N to Gateway OS

This replaces the Dragino LPS8N firmware of the gateway with the
Chirpstack Gateway OS.

1. First download the latest ChirpStack Gateway OS image for the Dragino LPS8N (see above table).
2. In the Dragino LPS8N web-interface, nativate to **System > Firmware Upgrade**.
3. Click the **Browse** button under _Upload Firmware File_, select the ChirpStack Gateway OS image and click **Upload**.
4. Once the image has been uploaded, click **Proceed** (do not enable any of the _Preserve Settings_ checkboxes!).
5. Wait until the firmware is flashed, this will take some minutes.

#### Gateway OS to LPS8N

This replaces the ChirpStack Gateway OS with the Dragino LPS8N firmware.

1. First download the latest firmware from [https://www.dragino.com/downloads/index.php?dir=LoRa_Gateway/LPS8/Firmware/Release/](https://www.dragino.com/downloads/index.php?dir=LoRa_Gateway/LPS8/Firmware/Release/).
2. In the ChirpStack Gateway OS web-interface, navigate to **System > Backup / Flash Firmware**.
3. Click the **Flash image...** button at the bottom of the page and select the Dragino image and click **Upload**.
4. De-select the **Keep settings and retain the current configuration** option and **Continue**.
   * You might see a _Image metadata not present_ warning, in which case you need to enable the **Force upgrade** checkbox.
5. Wait until the firmware is flashed, this will take some minutes.
6. Continue with the [Dragino LPS8N documentation](https://wiki.dragino.com/xwiki/bin/view/Main/User%20Manual%20for%20All%20Gateway%20models/LPS8N%20-%20LoRaWAN%20Gateway%20User%20Manual/).

## First usage

After flashing the ChirpStack Gateway OS image, follow the [Getting started](../getting-started.md)
documentation to setup your gateway.
