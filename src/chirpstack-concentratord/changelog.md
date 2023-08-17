# Changelog

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
