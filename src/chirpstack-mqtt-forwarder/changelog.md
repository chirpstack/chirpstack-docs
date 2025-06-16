# Changelog

## v4.4.0

### Features

#### Concentratord ZMQ API

This updates the Concentratord ZMQ API interface to be compatible with the
ChirpStack Concentratord v4.5+ and ChirpStack Gateway Mesh v4.1+. If using
the Concentratord backend, make sure to update this / these component(s)
too.

#### Mesh events / commands

This refactors the Gateway Mesh heartbeat payload into a more generic Mesh Event
payload, which can be used by both "known" and "proprietary" event types. The
first implemented "known" type is the heartbeat. It also implements Mesh
commands that can be used to execute commands configured on the Relay gateway.

### Bugfixes

* Handle `ant` payload sent by Tektelic gateways. ([#691](https://github.com/chirpstack/chirpstack/issues/691))
* Fix handling of negative altitude. ([#66](https://github.com/chirpstack/chirpstack-mqtt-forwarder/issues/66))

## v4.3.2

### Features

* Add MQTT connection callback configuration (e.g. to exec. command to turn on / off LED on gateway).

### Improvements

* Make MQTT re-connect interval configurable through `reconnect_interval` config option. ([#63](https://github.com/chirpstack/chirpstack-mqtt-forwarder/pull/63))
* On startup, validate that Gateway ID is 8-bytes. ([#65](https://github.com/chirpstack/chirpstack-mqtt-forwarder/pull/65))
* Update internal dependencies.

## v4.3.1

### Features

* Implement handling of ChirpStack Gateway Mesh heartbeat messages.

### Improvements

* Update internal dependencies.

### Bugfixes

* Spawn message-callback to not block MQTT eventloop. ([#57](https://github.com/chirpstack/chirpstack-mqtt-forwarder/issues/57))

## v4.3.0

### Features

* Expose duty-cycle stats exposed by ChirpStack Concentratord.
* Make it possible to use environment variables in configuration (like ChirpStack).

### Improvements

* Update internal dependencies.

## v4.2.3

### Improvements

* Add `keep_alive_interval` config option for MQTT. ([#54](https://github.com/chirpstack/chirpstack-mqtt-forwarder/issues/54)) 
* Update internal dependencies.

## v4.2.2

### Bugfixes

* Fix mips corrupted payloads regression re-introduced by v4.2.1. ([#50](https://github.com/chirpstack/chirpstack-mqtt-forwarder/issues/50))

## v4.2.1

### Bugfixes

* Fix firewall security issue on Kerlink gateways.

## v4.2.0

### Features

#### `paho-mqtt` to `rumqttc` refactor

This refactors the code to use `rumqttc` instead of `paho-mqtt` as MQTT client
library. As `rumqttc` is a pure-Rust MQTT client, which uses `rustls` for TLS,
it is no longer needed to statically link against OpenSSL.

### Improvements

* Add AMD64 build. ([#46](https://github.com/chirpstack/chirpstack-mqtt-forwarder/pull/46))
* Reconnect ZeroMQ REQ socket on error.
* Update dependencies.

### Bugfixes

* Fix extending meta-data in Concentratord backend.
* Fix `udp_bind` > `bind` error in config examples. ([#49](https://github.com/chirpstack/chirpstack-mqtt-forwarder/pull/49))
* Fix corrupted payloads on mips / mipsel platforms. ([#50](https://github.com/chirpstack/chirpstack-mqtt-forwarder/issues/50), [#52](https://github.com/chirpstack/chirpstack-mqtt-forwarder/issues/52))

## v4.1.3

### Improvements

* Update RAK packaging for WisGateOS2 v2.2+.
* Implement `preamble` and `no_crc` fields.
* Update internal dependencies.

## v4.1.2

### Features

* Provide package for Multitech Conduit AP3.

### Improvements

* Update internal dependencies.

## v4.1.1

### Improvements

* Add packaging for RAK `ramips_24kec` based gateways.
* Update internal dependencies.

### Bugfixes

* Fix setting empty username and password in MQTT. ([#257](https://github.com/chirpstack/chirpstack/issues/257))

## v4.1.0

### Features

### Improvements

#### DevAddr and JoinEUI filters

This adds a `[backend.filters]` configuration section under which it is
possible to configure the prefix-filters for DevAddr (data uplink) and
JoinEUIs (join-requests).

#### Time fallback

This adds a `time_fallback_enabled` configuration option to the `semtech_udp`
section. If enabled and no RX time is provided by the packet-forwarder, then
the ChirpStack MQTT Forwarder server-time will be used as fallback.

#### Build improvements

The provided pre-compiled binaries are fully static and based on _musl libc_.
This removes the need to compile against a very old verion of _glibc_ to
stay compatible with old gateway firmwares.

#### Other improvements

* Update dependencies.

### Bugfixes

* Fix bootstript for Dragino `.ipk` package.

## v4.0.0

This marks the initial release of the ChirpStack MQTT Forwarder.

The ChirpStack MQTT Forwarder is a re-implementation of the ChirpStack Gateway
Bridge for gateway installations only. The aim of this component is to provide
a more lightweight solution, that also can be installed on gateways with a
limited amount of memory (for example Dragino OpenWRT based gateways).
