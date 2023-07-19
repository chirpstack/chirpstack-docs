# Downloads

Please refer to the [Installation](./install/index.md) section for per-target
downloads.

## Older version

ChirpStack Gateway OS has been migrated from [Yocto](https://www.yoctoproject.org/)
to [OpenWrt](https://openwrt.org/) since version v4.2.0, making it possible to
configure all options through a web-interface. If you are looking for the older
version using the `gateway-config` command for configuration, then please use
one of the download links below:

### Image types

There are two file types:

* `.wic.gz` - Image to use for an initial installation
* `.swu` - Software update file, see [Software update](../use/software-update.md)

### Image links

* [Raspberry Zero W](http://artifacts.chirpstack.io/downloads/chirpstack-gateway-os/raspberrypi/raspberrypi0-wifi/{{yocto_gateway_os_version}}/)
* [Raspberry Pi 1](http://artifacts.chirpstack.io/downloads/chirpstack-gateway-os/raspberrypi/raspberrypi/{{yocto_gateway_os_version}}/)
* [Raspberry Pi 3](http://artifacts.chirpstack.io/downloads/chirpstack-gateway-os/raspberrypi/raspberrypi3/{{yocto_gateway_os_version}}/)
* [Raspberry Pi 4](http://artifacts.chirpstack.io/downloads/chirpstack-gateway-os/raspberrypi/raspberrypi4/{{yocto_gateway_os_version}}/)

### SD card flashing

#### Using Balena Etcher

* Download the SD card image for your Raspberry Pi (ending with `.wic.gz`).
* Click the **Flash from file** option in the [Balena Etcher](https://www.balena.io/etcher/) interface.
* Flash the SD card.

#### Using Win32DiskImager

* Download the SD card image for your Raspberry Pi (ending with `.wic.gz`).
* Extract the `.wic.gz` image using for example [7-Zip](https://www.7-zip.org/).
* Within the [Win32DiskImager](http://sourceforge.net/projects/win32diskimager/) interface, select the extracted `.wic` file.
  Please note that you must select the `*.*` filter in the _Select a disk image_ popup before you can select the `.wic` file.
* Flash the SD card.
