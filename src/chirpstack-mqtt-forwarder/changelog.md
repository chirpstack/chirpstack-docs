# Changelog

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
