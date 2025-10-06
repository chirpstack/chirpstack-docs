# Changelog

## v4.5.2

### Improvements

* Align `asap_count` in JIT queue with Semtech SX1302 implementation.
* Update internal dependencies and workspace configuration.

## v4.5.1

### Improvements

* Make Gateway ID configurable (2g4).
* Update internal dependencies.

## v4.5.0

### Features

#### ZMQ API

This updates the ZMQ API to make it easier and less error-prone to handle all
events and commands implemented by the Concentratord. This has been achieved
by refactoring the multi-part ZMQ messages by a single ZMQ message containing
a "wrapper" Protobuf message containing a `oneof` field containing the
different event / command messages.

#### Get location API

This adds a new ZMQ API command, to retrieve the current gateway location.

## v4.4.8

### Improvements

* Update internal dependencies.

## v4.4.7

### Improvements

* Make reset compatible with cdev uAPI v2. 

## v4.4.6

### Improvements

* Set read-timeout + fail if no 'u-blox' device was found (Gpsd). ([#198](https://github.com/chirpstack/chirpstack-concentratord/issues/198))
* Internal dependency updates.

### Bugfixes

* Fix incorrect FSK datarate multiplication.

## v4.4.5

### Improvements

* Provide option to override SX1302/3 embedded Gateway ID. ([#187](https://github.com/chirpstack/chirpstack-concentratord/pull/187))
* Update internal dependencies.

### Bugfixes

* Fix default `sx1302` model name. ([#188](https://github.com/chirpstack/chirpstack-concentratord/issues/188))

## v4.4.4

### Bugfixes

* Fix parsing fix status in case NMEA RMC has 12 fields (e.g. using Gpsd).

## v4.4.3

### Features

#### Override pin / device-path config

It is now possible to override all (reset)pin configuration and SPI / GNSS / 
I2C device-paths of all concentrator modules. A configuration example is
added to the `configfile` template.

#### GPSd support

This adds support for using Concentratord together with [GPSd](https://gpsd.gitlab.io/gpsd/).
If GPSd has been configured (e.g. `gnss_dev_path="gpsd://localhost:2947"`), the
Concentratord will automatically enable the `NAV-TIMEGPS` binary message of
the U-blox module on start. By using GPSd, it is possible to also use the GNSS
module for time-synchronization by using GPSd as a time-source for NTPD.

### Improvements

* Add more log context in case of (beacon) enqueue error.
* Refactor min/max tx frequency validation (such that gateway can also use uplink frequencies).

### Bugfixes

* Update `multitech_mtac_lora_h_915` SPI path. ([#167](https://github.com/chirpstack/chirpstack-concentratord/issues/167))
* Set `stats_interval` on receiving gateway configuration. ([#181](https://github.com/chirpstack/chirpstack-concentratord/issues/181))

## v4.4.2

### Improvements

* Update internal dependencies.
* Fix typo (Ctatus > Status). ([#179](https://github.com/chirpstack/chirpstack-concentratord/pull/179))

## v4.4.1

### Bugfixes

* Fix duty-cycle tracking causes downlink to be rejected for antenna gains < 2dBi. ([#157](https://github.com/chirpstack/chirpstack-concentratord/issues/157))

## v4.4.0

### Features

#### Duty-cycle support

This implements duty-cycle tracking and reporting support for the `EU868` region.
By default this does not enforce the duty-cycle limitations, enforcing can be
enabled by setting the model-flag `ENFORCE_DC`.

#### Concentrator module updates

* Add `EU433` support to `semtech_sx1302c490gw1` module.
* Add `IN865` and `RU864` support to `semtech_sx1302c868gw1` module.
* Add `IN865` and `RU864` support to `semtech_sx1302css868gw1` module.
* Add USB support to `seeed_wm1302` module (`USB` model-flag). ([#151](https://github.com/chirpstack/chirpstack-concentratord/pull/151))

#### Other features

* Make it possible to use environment variables in configuration (like ChirpStack).

### Improvements

* Improve error handling in threads.
* Update internal dependencies.

### Bugfixes

* Add missing `region_as923_2.toml`, `_3.toml` and `_4.toml` config files. ([#150](https://github.com/chirpstack/chirpstack-concentratord/issues/150))

## v4.3.5

### Features

* Add support for `embit_emb_lr1302_mpcie` module.

### Improvements

* Refactor (reset) pin configuration. ([#126](https://github.com/chirpstack/chirpstack-concentratord/pull/126))
* Update internal dependencies.

### Bugfixes

* `sx1302`: No need to trigger (reset) pins if using USB.

## v4.3.4

### Improvements

* Mention in config that antenna gain is in dBi.
* Implement `preamble` and `no_crc` fields.
* Update internal dependencies.

## v4.3.3

### Features

* Add support for Multitech Conduit AP3 (EU868 and US915).
* Add missing `waveshare_sx1302_lorawan_gateway_hat` regions. ([#92](https://github.com/chirpstack/chirpstack-concentratord/issues/92))

### Improvements

* Update internal dependencies.

### Bugfixes

* Update (modified) SX1302 HAL version to fix `qsort` issue. ([#107](https://github.com/chirpstack/chirpstack-concentratord/issues/107))

## v4.3.2

### Features

* Add AMD64 build. ([#99](https://github.com/chirpstack/chirpstack-concentratord/pull/99))
* Add `AU915` support for `risinghf_rhf0m301` module. ([#104](https://github.com/chirpstack/chirpstack-concentratord/pull/104))

### Improvements

* Refactor gnss / spi / i2c device path overwrites.
* Refactor (reset)pin overwrites.


## v4.3.0

### Features

* Add missing regions for `seeed_wm1302_spi` module.

### Improvements

* Update dependencies.

### Important

This release removes the region suffix in the model names, e.g. `..._eu868`.
Instead, you must configure the `region` configuration option. Please refer
to the v4.2.0 release note.

## v4.2.3

### Features

* Add support for [RAK5148](https://store.rakwireless.com/products/2-4-ghz-mini-pcie-concentrator-module-for-lora-based-on-sx1280-rak5148).

### Internal

* Update dependencies.

## v4.2.2

### Features

* Add `AU915` support for Semtech SX1302 LoRa(R) CoreCell modules.

### Internal

* Update dependencies.

### Bugfixes

* Remove GNSS config from Seeed WM1302 configuration (the GNSS module is not supported by the Semtech HAL).

## v4.2.1

### Internal

* Fix filling of `wrapper::timespec` fields in case of `bindgen` generating `_bitfield_...` fields for memory alignment.

### Bugfixes

* Update documentation urls.

## v4.2.0

### Features

#### Refactor model configuration

This release adds a `region` configuration option and removes the region suffix
from the gateway / shield model. This changes the configuration from:

```toml
model="rak_rak2247_eu868"
```

to:

```toml
model="rak_rak2247"
region="EU868"
```

For backwards compatibility, the old `model` configuration remains valid in
this version, but will be removed in `v4.3.0`.

#### Other features

* Add `CN470` support for `semtech_sx1302c490gw1` shield.
* Add `AU915` support for `sandbox_lorago_port` shield.
* Add fine-timestamp for SX1302/SX1302. ([#66](https://github.com/chirpstack/chirpstack-concentratord/pull/66))

### Improvements

#### Build improvements

The provided pre-compiled binaries are fully static and based on _musl libc_.
This removes the need to compile against a very old verion of _glibc_ to
stay compatible with old gateway firmwares.

#### Other improvements

* Update internal dependencies.
* Implement time fallback option in case no GPS time-source is available. ([#49](https://github.com/chirpstack/chirpstack-concentratord/issues/49))

### Bugfixes

* Fix Class-B beacon channel-hopping.
* Fix typo in `2g4` logs (it would log under the `chirpstack-concentratord-sx1301` process name).

## v4.1.1

### Features

* Add `dragino_pg1302_eu868` model support.
* Add `dragino_pg1302_us915` model support.

### Improvements

* Update internal dependencies.

## v4.1.0

### Features

* Add `multitech_mtac_lora_2g4` model support.
* Add `multitech_mtac_003e00_eu868` model support.
* Add `multitech_mtac_003u00_us915` model support.
* Add support for Class-B beacons for SX1302 version.
* Expose CRC status + option to disable CRC check.

### Improvements

* Update Multitech Conduit packaging.
* Update Multitech Conduit AP packaging.
* Update Kerlink iFemtoCell packaging.
* Refactor reset-pin configuration.
* Implement support for using GPSd.
* Update parsing of UBX binary format parsing.
* Make I2C temperature sensor compatible + add support for TMP-102 sensor (used by the Multitech SX1302 MTAC modules).
* Update internal dependencies.

## v4.0.1

### Features

* Add `waveshare_sx1302_lorawan_gateway_hat_eu868` model support.

### Improvements

* Update internal dependencies.

## v4.0.0

### Features

This release implements the ChirpStack v4 API interface. Please note that this
version is not compatible with ChirpStack Gateway Bridge v3.

## v3.3.2

This release adds the following gateway model configurations:

* pi_supply_lora_gateway_hat_as923
* pi_supply_lora_gateway_hat_in865
* pi_supply_lora_gateway_hat_kr920
* pi_supply_lora_gateway_hat_ru864
* rak_2246_cn470
* rak_2246_eu433
* rak_2247_as923
* rak_2247_au915
* rak_2247_cn433
* rak_2247_eu433
* rak_2247_eu868
* rak_2247_in865
* rak_2247_kr920
* rak_2247_ru864
* rak_2247_us915
* rak_2287_cn470
* rak_2287_eu433
* rak_5146_as923
* rak_5146_au915
* rak_5146_cn470
* rak_5146_eu433
* rak_5146_eu868
* rak_5146_in865
* rak_5146_kr920
* rak_5146_ru864
* rak_5146_us915

## v3.3.1

This release adds the following gateway model configurations:

* imst_ic880a_ru864
* imst_ic880a_in865
* pi_supply_lora_gateway_hat_au915
* risinghf_rhf0m301_eu868
* risinghf_rhf0m301_us915

## v3.3.0

### Features

* Implement and expose various gateway stats aggregations (uplinks / downlinks per frequency and modulation parameters and downlinks per ack status).

### Bugfixes

* Remove Class-B beacon frequency correction on enqueue.

## v3.2.0

### Features

* Implement support for 2.4 GHz concentrator.
* Add `configfile` sub-command to binaries for printing the configuration template.

## v3.1.0

### Features

* Add support for setting static gateway location.
* Add support for RAK2287 module (SPI & USB). ([#16](https://github.com/brocaar/chirpstack-concentratord/pull/16))

### Improvements

* Update SX1302 HAL to v2.1.0.

## v3.0.3

### Bugfixes

* Fix `channel_max` calculation.

### Features

* Implement overwriting reset-pin for RPi shields.

## v3.0.2

### Bugfixes

* Fix sending 0, 0 coordinates when GPS is unavailable.

## v3.0.1

### Bugfixes

* Fix beacon loop termination on re-configuration and improve debug logging.

## v3.0.0

Initial stable release.
