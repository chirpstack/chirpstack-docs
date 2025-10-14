# Changelog

## v4.9.0-test.1

### Features

#### Dragino LPS8N

This release adds support for the Dragino LPS8N gateway.

### Updates

* Update ChirpStack Concentratord to v4.5.2.
* Update ChirpStack to v4.15.0.

### Improvements

* Update configuration for RAK7391 gateway. ([#140](https://github.com/chirpstack/chirpstack-gateway-os/pull/140))
* Add `watchcat` service to images. ([#141](https://github.com/chirpstack/chirpstack-gateway-os/issues/141))
* Auto-generate API secret for ChirpStack service.
* Make ChirpStack Gateway Mesh filter configuration configurable.
* Update to OpenWrt v24.10.3.
* RAK7391: Allow to select `ttyACM2` for concentrator modules. ([#19](https://github.com/chirpstack/chirpstack-openwrt-feed/pull/19))

### Bugfixes

* Fix `udp_server` > `udp_forwarder` in ChirpStack UDP Forwarder configuration. ([#146](https://github.com/chirpstack/chirpstack-gateway-os/issues/146))

## v4.8.3

### Bugfixes

#### Raspberry Pi: fix failing boot after upgrade

This release fixes an issue with the Raspberry Pi images that could cause
the bootprocess of the Raspberry Pi to fail after an upgrade. For more details
about this issue, please see:

* Issue [#139](https://github.com/chirpstack/chirpstack-gateway-os/issues/139)
* OpenWrt issue [#9113](https://github.com/openwrt/openwrt/issues/9113)
* OpenWrt pull-request [#19997](https://github.com/openwrt/openwrt/pull/19997)

In case you are affected by this issue and you have access to the terminal
(e.g. USB keyboard + HDMI display) you could try the following command to
restore the Raspberry Pi: `firstboot && reboot`. `firstboot` will normally
reset the overlayfs containing the configuration, but as in case of this issue
OpenWrt fails to mount it, this will cause OpenWrt to re-create the overlayfs
and restore the configuration backup to it after the reboot.

## v4.8.2

### Updates

* Update ChirpStack to v4.14.1.
* Update ChirpStack Gateway Mesh to v4.1.0.
* Update Concentratord to v4.5.1.
* Update OpenWrt to v24.10.2.

### Improvements

* ChirpStack Gateway Mesh: Make Relay ID configurable.
* ChirpStack Concentratord: Make Gateway ID configurable for sx1302/3 and 2g4.

### Bugfixes

* Only automatically set hostname if the current value is `OpenWrt`.
* Fix broken `python3` dependency in `Dockerfile-devel` (after Debian update).

## v4.8.1

### Updates

* Update ChirpStack UDP Forwarder to v4.2.1.

## v4.8.0

### Features

#### Raspberry Pi 5

This release adds support for Raspberry Pi 5 based gateways.

### Updates

* Update ChirpStack to v4.13.0.
* Update ChirpStack Concentratord to v4.5.0.
* Update ChirpStack MQTT Forwarder to v4.4.0.
* Update ChirpStack UDP Forwarder to v4.2.0.
* Update ChirpStack Gateway Mesh to v4.1.0-test.1.
* Update Node-RED to v4.0.9.

## v4.7.1

### Updates

* Update ChirpStack to v4.12.0.
* Update ChirpStack Concentratord to v4.4.8.
* Update ChirpStack MQTT Forwarder to v4.3.2.
* Update ChirpStack UDP Forwarder to v4.1.10.
* Update Node-RED to v3.1.15.
* Update OpenWrt to v24.10.1.

### Improvements

* Seeed SenseCAP M2: Fix Wi-Fi mac address (to be consistent with package label).

## v4.7.0

### Updates

* Update ChirpStack to v4.11.1.
* Update ChirpStack Concentratord to v4.4.7.
* Update ChirpStack Gateway Mesh to v4.0.1.

### Features

#### Seeed SenseCAP M2 support

This release adds support for the [Seeed SenseCAP M2](https://wiki.seeedstudio.com/Network/SenseCAP_Network/SenseCAP_M2_Multi_Platform/SenseCAP_M2_Multi_Platform_Overview/)
gateway.

#### OpenWrt 24.10

This release updates OpenWrt to v24.10.0.

### ChirpStack Gateway Mesh

This release adds AS923 and IN865 region configuration.

### Bugfixes

* Fix missing `chirpstack-gateway-mesh` for `bcm2708` target. ([#122](https://github.com/chirpstack/chirpstack-gateway-os/issues/122))

## v4.6.2

### Updates

* Update ChipStack to v4.10.2.
* Update ChirpStack Concentratord to v4.4.6.

### Improvements

* ChirpStack MQTT Forwarder: Make commands and metadata configurable in UI.

## v4.6.1

This release only includes changes for the ChirpStack Gateway OS Full 
(Raspberry Pi) and the RAK7391 (Raspberry Pi CM4) images.

### Improvements

* Import LoRaWAN device repository after starting ChirpStack (such that the web-interface is immediately accessible after boot).

### Bugfixes

* Fix mkdir typo in `chirpstack` init script.

## v4.6.0

This release only includes changes for the ChirpStack Gateway OS Full 
(Raspberry Pi) and the RAK7391 (Raspberry Pi CM4) images.

### Features

#### PostgreSQL > SQLite migration

This release replaces the ChirpStack PostgreSQL database with SQLite.
By migrating to SQLite, PostgreSQL is no longer required and thus
reducing the dependencies. It also makes it easier to create backups
as the SQLite database is stored as a single file (`/srv/chirpstack/chirpstack.sqlite`).

### Before you upgrade

Before you upgrade, you should migrate the data to SQLite first, as after
upgrading the PostgreSQL binaries will no longer be available. To migrate
ChirpStack data, you can use the migration utility (for ARMv7) that will
automatically create the SQLite database and migrate the data.

You must download the `.ipk` package from [https://github.com/chirpstack/chirpstack-pg-to-sqlite/releases](https://github.com/chirpstack/chirpstack-pg-to-sqlite/releases).
In the ChirpStack Gateway OS v4.5.x web-interface, navigate to **System > Software**
click the **Upload package...** button and upload the `.ipk` package. After 
uploading it will immediately initialize the SQLite database and print the 
status of the migration.

### Updates

* ChirpStack to v4.10.1.

## v4.5.5

### Updates

* Update ChirpStack Concentratord to v4.4.5.

### Bugfixes

* Fix random wifi interface MAC for MIPS-based RAK gateways.
* UDP Forwarder: Fix missing slog configuration. ([#120](https://github.com/chirpstack/chirpstack-gateway-os/issues/120))

## v4.5.4

### Bugfixes

* Fix duplicated `[concentratord]` configuration section preventing ChirpStack UDP Forwarder from starting.

## v4.5.3

### Updates

* Update OpenWrt to v23.05.5.
* Update ChirpStack Concentratord to v4.4.4.

### Improvements

#### GPS time synchronization

This adds GPS time synchronization for the following gateways:

* RAK7267
* RAK7289v2

If these gateways are operating as Relay gateways (ChirpStack Gateway Mesh)
this assures that the heartbeat timestamp is accurate.

#### RAK7268v2 multi-slot

This changes the RAK7268v2 image to support multiple slots. While this gateway
is (currently) sold with a single concentrator module, it is possible to
install an additional concentrator module, for example a 2.4GHz module to use
together with the ChirpStack Gateway Mesh feature. ([#5](https://github.com/chirpstack/chirpstack-openwrt-feed/pull/5))

**Note:** After upgrading, you must re-configure the Concentratord / MQTT
Forwarder and Gateway Mesh (if used) to use the _Slot 1_ configuration.

#### Other improvements

* Show all Gateway IDs in web-interface footer. ([#7](https://github.com/chirpstack/chirpstack-openwrt-feed/pull/7))
* RAK7391: Add SPI device-paths to concentrator selector. ([#6](https://github.com/chirpstack/chirpstack-openwrt-feed/pull/6))
* RAK7391: Add reset GPIOs for SPI concentrators. ([#9](https://github.com/chirpstack/chirpstack-openwrt-feed/pull/9))

## v4.5.2

### Bugfixes

* Fix GNSS device path for RAK7289v2 and RAK7267.
* Include correct configuration for RAK7267.

## v4.5.1

### Bugfixes

* Fix regression in ChirpStack MQTT Forwarder init script.

## v4.5.0

### Upgrade notes

#### Backup and restore (Full image)

Before upgrading, please make sure to create a backup. Please see 
**System > Custom Commands**.

### Features

#### ChirpStack Gateway Mesh

This release adds the [ChirpStack Gateway Mesh](../chirpstack-gateway-mesh/index.md)
component to the ChirpStack Gateway OS images.

#### New targets

This release adds support for the following RAK gateways:

* [RAK7267 - WisGate Soho Pro Gateway](https://store.rakwireless.com/products/wisgate-soho-pro-gateway-rak7267)
* [RAK7268v2 - WisGate Edge Lite 2](https://store.rakwireless.com/products/rak7268-8-channel-indoor-lorawan-gateway)
* [RAK7289v2 - WisGate Edge Pro](https://store.rakwireless.com/products/rak7289-8-16-channel-outdoor-lorawan-gateway)
* [RAK7391 - WisGate Connect / Compute Module 4 (CM4) Carrier Board](https://store.rakwireless.com/products/wisgate-connect-base-kit-rak7391)

### Updates

* Update OpenWrt to v23.05.4.
* Update ChirpStack to v4.9.0.
* Update ChirpStack MQTT Forwarder to v4.3.1.
* Update ChirpStack Concentratord to v4.4.2.
* Update ChirpStack UDP Forwarder to v4.1.8.

## v4.4.0

### Upgrade notes

#### Backup and restore (Full image)

Before upgrading, please make sure to create a backup. Please see 
**System > Custom Commands**.

#### Enable services

This release adds an _Enabled_ configuration flag to the ChirpStack
Concentratord, ChirpStack MQTT Forwarder and ChirpStack UDP Forwarder
in the configuration. After upgrading, you must enable the services
you would like to use.

### Updates

* Update OpenWrt to v23.05.3.
* Update ChirpStack to v4.8.1.
* Update ChirpStack MQTT Forwarder to v4.3.0.
* Update ChirpStack Concentratord to v4.4.1.

### Bugfixes

* Add `-c` flag to `pg_restore` command (PostgreSQL database backup restore).

## v4.3.2

### Bugfixes

* Fix regression introduced by v4.3.1 causing the ChirpStack UDP Forwarder to not start (missing `chirpstack-udp-forwarder.sh`).

## v4.3.1

### Improvements

* RPi: Add FTDI kernel module for USB -> Serial devices. ([#105](https://github.com/chirpstack/chirpstack-gateway-os/issues/105))
* Refactor ChirpStack package scripts and configuration in preparation to support targets with multiple concentrator modules.

### Bugfixes

* Fix error in SX1301 init script. ([#108](https://github.com/chirpstack/chirpstack-gateway-os/issues/108))

## v4.3.0

### Upgrade notes

This release updates the PostgreSQL database version. If you would like to
retain all data, You must create a PostgreSQL + Redis backup **before**
upgrading. You must use the following commands (using SSH):

```bash
mkdir -p /srv/backup
chmod 777 /srv/backup
sudo -u postgres /usr/bin/pg_dump -h localhost -d chirpstack -F c -f /srv/backup/chirpstack.pg

service redis stop
cp /srv/redis/dump.rdb /srv/backup/chirpstack.redis
service redis start
```

After upgrading, you can use the _Restore ChirpStack backup_ command to
restore the backup. For future backups, you can use the _Create ChirpStack
Backup_ command (see features) after upgrading.

### Features

#### PostgreSQL & Redis backup / restore

This adds _Create ChirpStack backup_ and _Restore ChirpStack backup_
commands under **System > Custom Commands**.

#### Other features

* Add Wireguard VPN support.
* Add experimental support for RAK7268v2 gateways (to be documented).

### Updates

* Update OpenWrt to v23.05.2.
* Update ChirpStack to v4.6.0.
* Update ChirpStack Concentratord to v4.3.5.
* Update ChirpStack MQTT Forwarder to v4.1.3.
* Update ChirpStack UDP Forwarder to v4.1.6.

## v4.2.0

### Features

This is a very exciting release, as this release migrates the ChirpStack Gateway OS
from [Yocto](https://www.yoctoproject.org/) to [OpenWrt](https://openwrt.org/).
Thanks to [LuCI](https://github.com/openwrt/luci/) and the [UCI configuration system](https://openwrt.org/docs/guide-user/base-system/uci)
it is now possible to configure the provided ChirpStack components as well
things like network configuration through a web-interface instead of a CLI.
This release also splits the [packages](https://github.com/chirpstack/chirpstack-openwrt-feed)
from the ChirpStack Gateway OS [configuration](https://github.com/chirpstack/chirpstack-openwrt-config/)
such that these can be integrated in other OpenWrt based projects.
Main features that this release brings:

#### Web-interface

There is no need to configure the ChipStack Gateway OS using a terminal.
Powered by [LuCI](https://github.com/openwrt/luci/)), the ChirpStack Gateway OS
now provides an easy-to-use web-interface to configure ChirpStack components and
system configuration.

#### Custom packages

Through the web-interface (and cli) it is possible to install additional
software using the OpenWrt package repositories. This has been requested
several times, but until now would require to compile a custom ChirpStack
Gateway OS image which would take several hours to complete.

### Updates

* Update ChirpStack to v4.4.3.
* Update ChirpStack Concentratord to v4.2.3.
* Update ChirpStack UDP Forwarder to v4.1.2.
* Update ChirpStack MQTT Forwarder to v4.1.0.
* Update Node-RED to v3.0.2.

### Upgrade notes

Unfortunately it is not possible to migrate from ChirpStack Gateway OS v4.1.1
to ChirpStack Gateway OS v4.2.0. You need to re-flash your SD-Card with the
ChirpStack Gateway OS v4.2.0 image. 

## v4.1.1

### New shields

* Add support for Dragino PG1302 (EU868 + US915).

### Updates

* Update ChirpStack to v4.3.0.
* Update ChirpStack Concentratord to v4.1.1.
* Update ChirpStack UDP Forwarder to v4.1.1.

### Bugfixes

* Change `band_` to `_region_` after rename in Concentratord repo. ([#97](https://github.com/chirpstack/chirpstack-gateway-os/issues/97))

## v4.1.0

### Features

#### ChirpStack MQTT Forwarder

This release replaces the ChirpStack Gateway Bridge with the ChirpStack MQTT
Forwarder. The following MQTT configuration values will be automatically migrated
from the `chirpstack-gateway-bridge.toml` configuration file:

* `server`
* `topic_prefix` (detected based on `event_topic_template`)
* `username`
* `password`
* `ca_cert`
* `tls_cert`
* `tls_key`

### Updates

* ChirpStack to v4.2.0.
* ChirpStack Concentratord to v4.1.0.
* ChirpStack UDP Forwarder to v4.1.0.

## v4.0.2

### Updates

* ChirpStack to v4.1.1.
* ChirpStack Gateway Bridge to v4.0.3.
* Node-RED to v3.0.2.

### Bugfixes

* Fix NodeJS / Node-RED crashing with "Bus error" error message. ([#95](https://github.com/chirpstack/chirpstack-gateway-os/issues/95))

## v4.0.1

### Updates

* ChirpStack to v4.0.3.
* ChirpStack Concentratord to v4.0.1.
* The Things Network [lorawan-devices](https://github.com/TheThingsNetwork/lorawan-devices) to latest version.

### New shields

* Waveshare SX1302 LoRaWAN Gateway HAT

### Bugfixes

* Add missing `/etc/fw_env.config` file (see upgrade note below).
* Fix missing Wiregard binaries. ([#93](https://github.com/chirpstack/chirpstack-gateway-os/issues/93))

### Update note

If upgrading from v4.0.0 using the `software-update` utility, you must run the
following command first because of the missing `/etc/fw_env.config` file:

```
sudo echo "/boot/uboot.env 0x0000    0x4000" > /etc/fw_env.config
```

## v4.0.0

### Main changes

#### ChirpStack v4

This release replaces ChirpStack Network Server (v3) and ChirpStack Application
Server (v3) with ChirpStack v4.0.0.

#### MQTT certificates

If using the _full_ image, MQTT TLS certificates will be automatically generated on
first boot and configured in ChirpStack. For consuming device events throug the
MQTT integration, you can generate certificates in the web-interface under
the application integrations.

**Note:** localhost connections (e.g. if using Node-RED) to the MQTT do not
need TLS certificates for authentication and authorization.

#### gateway-config improvements

If using the _full_ ChirpStack Gateway OS image, which includes the ChirpStack
Network Server, the `gateway-config` script will automatically create the
gateway in the ChirpStack database for you.

#### Region configuration

The _full_ image contains update region configuration for the common regions.
The desired region can be selected using the `gateway-config` configuration script.

#### New shields

Support for the following shields has been added:

- RAK2247
- RAK5147
- Seeed WM1302

#### Yocto Kirkstone

ChirpStack Gateway OS has been updated to Yocto Kirkstone. This should fix
boot problems on some Pi 4 revisions, as this release contains an updated
bootloader.

### Other updates

* ChirpStack Concentratord to v4.0.0
* ChirpStack Gateway Bridge to v4.0.0
* Node-RED to v2.1.6

## v3.6.0

### Updates

* [ChirpStack Application Server](https://www.chirpstack.io/application-server) to v3.17.7
* [ChirpStack Gateway Bridge](https://www.chirpstack.io/gateway-bridge) to v3.14.2
* [ChirpStack Network Server](https://www.chirpstack.io/network-server) to v3.16.3

**Note:** The ChirpStack Gateway Bridge version included is compatible with
ChirpStack v4 which will soon be released.

## v3.5.1

### Updates

* Update [ChirpStack Concentratord](https://www.chirpstack.io/concentratord/) to v3.3.1.

### Improvements

* Additional region configurations
	* IMST Lite
		* RU864
		* IN865
	* Pi Supply - LoRa Gateway HAT
		* AU915
	* RAK2245
		* IN865
		* RU864
	* RAK2246(G)
		* IN865
		* RU864

### Bugfixes

* Fix configuration for RisingHF RHF0M301 shield (GPIO7 issue). [#72](https://github.com/brocaar/chirpstack-gateway-os/issues/72)

## v3.5.0

### Features

#### Node-RED integration

This release includes [Node-RED](https://nodered.org/) into the *full* image
version, with the [node-red-contrib-chirpstack](https://github.com/brocaar/node-red-contrib-chirpstack/)
package pre-installed. Please note that it must be enabled using the
`gateway-config` configuration utility first.

### Updates

* Update [ChirpStack Concentratord](https://www.chirpstack.io/concentratord/) to v3.3.0.
* Update [ChirpStack Gateway Bridge](https://www.chirpstack.io/gateway-bridge/) to v3.13.1.
* Update [ChirpStack Network Server](https://www.chirpstack.io/network-server/) to v3.15.0.
* Update [ChirpStack Application Server](https://www.chirpstack.io/application-server/) to v3.17.0.
* Update [Yocto](https://www.yoctoproject.org/) BSP and open-embedded layers from *dunfell* to *hardknott*.

### Notes

As this release increases the size of the rootfs and data partitions, updating
using a `.swu` image is not possible.

## v3.4.0

### Features

* Add support for Semtech 2.4 GHz gateway module.
* Add support for RAK2287 gateway module.

### Updates

* Update [ChirpStack Concentratord](https://www.chirpstack.io/concentratord/) to v3.2.0.
* Update [ChirpStack Network Server](https://www.chirpstack.io/network-server/) to v3.12.2.
* Update [ChirpStack Application Server](https://www.chirpstack.io/application-server/) to v3.14.0.

### Improvements

* Disable append only in Redis configuration.
* Align US915 and AU915 config examples (https://github.com/brocaar/chirpstack-docs/issues/38).
* Cleanup ChirpStack Gateway OS recipe structure.

### Bugfixes

* Fix disabling bluetooth on Raspberry Pi 3 (so that UART pins can be used for GNSS module).

## v3.3.3

### Features

* Add test-version of [ChirpStack UDP Bridge](https://github.com/brocaar/chirpstack-udp-bridge).

### Updates

* All meta-layers have been updated.
* Update [ChirpStack Concentratord](https://www.chirpstack.io/concentratord/) to v3.0.3.
* Update [ChirpStack Network Server](https://www.chirpstack.io/network-server/) to v3.11.0.
* Update [ChirpStack Application Server](https://www.chirpstack.io/application-server/) to v3.13.2.

### Bugfixes

* Fix bootfiles path in rpi-config. ([#63](https://github.com/brocaar/chirpstack-gateway-os/pull/63))
* Remove libubootenv. ([#64](https://github.com/brocaar/chirpstack-gateway-os/pull/64/))
* Update u-boot `CONFIG_SYS_BOOTM_LEN` to 16M.

## v3.3.2

### Updates

* [ChirpStack Concentratord](https://www.chirpstack.io/concentratord/) is updated to v3.0.2.

## v3.3.1

### Updates

* [ChirpStack Concentratord](https://www.chirpstack.io/concentratord/) is updated to v3.0.1.
* [ChirpStack Application Server](https://www.chirpstack.io/application-server/) is updated to v3.12.1.
* [ChirpStack Gateway Bridge](https://www.chirpstack.io/gateway-bridge/) is updated to v3.9.2.

## v3.3.0

This marks the first non-testing release of the ChirpStack Gateway OS!

### Updates

* [ChirpStack Concentratord](https://www.chirpstack.io/concentratord/) is updated to v3.0.0.

## v3.3.0-test.9

### Updates

* [ChirpStack Concentratord](https://github.com/brocaar/chirpstack-concentratord/) is updated to v3.0.0-test.11.

### Improvements

* `gateway-config` has been updated for the [Pi Supply LoRa Gateway HAT](https://uk.pi-supply.com/products/iot-lora-gateway-hat-for-raspberry-pi).
* New `gateway-config` option has been added to reload the Gateway ID.

## v3.3.0-test.8

## Bugfixes

* Fix ChirpStack Application Server Makefile execution to include web-interface statics in binary.

## v3.3.0-test.7

### Updates

* [ChirpStack Gateway Bridge](https://www.chirpstack.io/gateway-bridge/) is updated to v3.9.1.
* [ChirpStack Network Server](https://www.chirpstack.io/network-server/) is updated to v3.10.0.
* [ChirpStack Application Server](https://www.chirpstack.io/application-server/) is updated to v3.11.0.
* [ChirpStack Concentratord](https://github.com/brocaar/chirpstack-concentratord/) is updated to v3.0.0-test.10.

### Features

* Add `gateway-config` wizard for MQTT configuration.
* Update to Yocto Dunfell + build Go apps from source. ([#55](https://github.com/brocaar/chirpstack-gateway-os/pull/55))

### Bugfixes

* Fix AU915 selection bugs for RAK concentrators. ([#56](https://github.com/brocaar/chirpstack-gateway-os/pull/56))

## v3.3.0-test.6

### Features

* Support has been added for the Raspberry Pi 4.
* Support has been added for the Raspberry Pi Zero W.
* Support has been added for the RAK2246 and RAK2246G shields.
* Class-B beacon configuration has been added to the band configuration.

### Updates

* [ChirpStack Application Server](https://www.chirpstack.io/application-server/) is updated to v3.10.0.
* [ChirpStack Concentratord](https://github.com/brocaar/chirpstack-concentratord) is updated to v3.0.0-test.9.
* [ChirpStack Network Server](https://www.chirpstack.io/network-server/) is updated to v3.9.0.

### Improvements

* `gateway-config` shows ChirpStack Gateway OS version.
* `gateway-config` shows Gateway ID.
* Change ISM band names to their common name. ([#47](https://github.com/brocaar/chirpstack-gateway-os/pull/47))
* `sx1301-reset` script has been modified to leave the reset pin as output.

## v3.3.0-test.5

### Updates

* [ChirpStack Application Server](https://www.chirpstack.io/application-server/) is updated to v3.9.0.
* [ChirpStack Concentratord](https://github.com/brocaar/chirpstack-concentratord) is updated to v3.0.0-test.8.
* [ChirpStack Network Server](https://www.chirpstack.io/network-server/) is updated to v3.8.1.

### Features

#### Wifi Access Point mode

On initial installation on a Raspberry Pi 3, the Raspberry Pi will start Wifi
in Access Point mode, so that it is possible to connect directly to the
Raspberry Pi over WIFI for configuration of the concentrator shield and to
re-configure the WIFI.

### Improvements

* RAK2245 configuration has been improved (using defined Concentratord model name).
* RAK832 configuration has been improved.
* RAK2245 / RAK831 AS923 channel-plan has been added. ([#43](https://github.com/brocaar/chirpstack-gateway-os/pull/43))

### Bugfixes

* Fix passing incorrect model flags in `gateway-config` for some gateways.

## v3.3.0-test.4

When updating from a previous v3.3.0 version, it is recommended to re-run the
`gateway-config` utility to update the concentrator configuration.

### Updates

* [ChirpStack Concentratord](https://github.com/brocaar/chirpstack-concentratord) is updated to v3.0.0-test.6.

## v3.3.0-test.3

### Updates

* [ChirpStack Gateway Bridge](https://www.chirpstack.io/gateway-bridge/) is updated to v3.7.1.

## v3.3.0-test.2

### Updates

* [ChirpStack Gateway Bridge](https://www.chirpstack.io/gateway-bridge/) is updated to v3.7.0.
* [ChirpStack Network Server](https://www.chirpstack.io/network-server/) is updated to v3.7.0.
* [ChirpStack Application Server](https://www.chirpstack.io/application-server/) is updated to v3.8.0.
* [ChirpStack Concentratord](https://github.com/brocaar/chirpstack-concentratord) is updated to v3.0.0-test.5.

### Supported hardware

* Raspberry Pi 1 B+ support has been added (for IMST Lite Gateway)
* [IMST Lite Gateway](https://wireless-solutions.de/products/long-range-radio/lora-lite-gateway.html) has been added to gateway configuration script.

## v3.3.0-test.1

**This is a rewrite of the ChirpStack Gateway OS, you must re-flash your SD Card
to update!** Currently this version only targets the Raspberry Pi 3.

### Features

* Yocto has been updated to version 3.0.
* Software updates are now handled by [SWUpdate](https://github.com/sbabic/swupdate).
* The Semtech UDP Packet Forwarder has been replaced by [ChirpStack Concentratord](https://github.com/brocaar/chirpstack-concentratord).

### Updates

* [ChirpStack Gateway Bridge](https://www.chirpstack.io/gateway-bridge/) is updated to v3.7.0-test.2.
* [ChirpStack Network Server](https://www.chirpstack.io/network-server/) is updated to v3.7.0-test.1.
* [ChirpStack Application Server](https://www.chirpstack.io/application-server/) is updated to v3.8.0-test.1.

### Fixes

* Redis database does not start on boot after power failure that corrupts append only file. ([#32](https://github.com/brocaar/chirpstack-gateway-os/issues/32))

## v3.2.0test1

### General

* Update ChirpStack Gateway Bridge to v3.5.0.
* Update ChirpStack Network Server to v3.5.0.
* Update ChirpStack Application Server to v3.6.1.

### Features

* Add support for [Semtech SX1302 CoreCell](https://www.semtech.com/products/wireless-rf/lora-gateways/sx1302cxxxgw1) evaluation kit.

### Bugfixes

* Fix boot issue due to storage device not yet initialized. ([#9](https://github.com/brocaar/chirpstack-gateway-os/issues/9))
* Fix ChirpStack Network Server `enabled_uplink_channels` configuration. ([#26](https://github.com/brocaar/chirpstack-gateway-os/issues/26))

## v3.1.0test1

This release renames LoRa Gateway OS to ChirpStack Gateway OS.
See the [Rename Announcement](https://www.chirpstack.io/r/rename-announcement) for more information.

## v3.0.0test3

### LORIX One

* Fix Wiregard kernel module dependencies.

## v3.0.0test2

### General

* Update LoRa App Server to v3.2.0.
* Update LoRa Gateway Bridge to v3.1.0.
* Update LoRa Server to v3.1.0.
* Update Monit to 5.26.0 and set check interval to 10 seconds.
* Add `PersistentKeepalive = 25` to Wiregard example config.
* Update openembedded layers to latest versions.

### Raspberry Pi

* Fix concentrator ordering.

## v3.0.0test1

### General

* LoRa App Server v3.1.0.
* LoRa Server v3.0.2.
* LoRa Gateway Bridge v3.0.1.

### Raspberry Pi

* Add support for the [Pi Supply LoRa Gateway Hat](https://uk.pi-supply.com/products/iot-lora-gateway-hat-for-raspberry-pi).
* Fix HDMI related boot issue. ([#9](https://github.com/brocaar/lora-gateway-os/issues/9))

## v2.0.0test4

### General

* Add Wiregard VPN client
* Bump LoRa Server package versions

### Raspberry Pi

* Change SPI speed to 2MHz (required by RAK2245)
* Add IMST iC980A configuration
* Add RAK2245 configuration

### LORIX One 512MB

* Fix u-boot command

## v2.0.0test3

### LORIX One

* Fix setting the MAC address from EEPROM.

## v2.0.0test2

### General

* Implement Mender for (OTA) system updates.
* Implement OverlayFS over read-only root filesystem.
* Update LoRa Gateway Bridge to v2.6.2.
* [lora-gateway-os-full] Update LoRa Server to v2.4.1.

### Raspberry Pi

* Add support for Sandbox Electronics LoRaGo PORT concentrator.
* Implement all US915 and AU915 channel-blocks. ([#2](https://github.com/brocaar/lora-gateway-os/pull/2))
* [lora-gateway-os-full] Automatic (re)configure LoRa Server on setting the concentrator channel-plan.

## v2.0.0test1

* Initial test release.
