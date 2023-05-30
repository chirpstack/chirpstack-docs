# Changelog

## v4.4.0 (in development)

### Features

#### Relay support (TS011)

This adds support for the Relay specification (TS011). In the Device Profile
it is possible to configure the Relay and Relay capable end-device parameters.
Under the Application view it is possible to assign devices to a Relay
(required for filtering and exchanging the device list with the Relay).

**Note:** Please note that this requires a Relay and Relay capable end-device.

### Improvements

#### Build changes

The build configuration has been updated to generate fully static binaries
based on _musl libc_. This solves the issue where in some cases ChirpStack
would not connect over TLS to a PostgreSQL database. ([#156](https://github.com/chirpstack/chirpstack/issues/156)).

This also changes the Docker base image to `alpine`, reducing the Docker image
size by ~ 50% compared to `debian:buster-slim`. Within the `Dockerfile` we now
`COPY` the already compiled binaries, which also reduces the build time on
release.

If you are compiling ChirpStack from source, please refer to the `README.md`
in the source repository as some commands have changed.

#### IFTTT integration

Configuration has been added to configure the prefix of the event name and
to send arbitrary JSON payloads instead of the 3 value payload.

#### Other improvements

* Add `lrwn_filters` crate for filtering LoRaWAN PHYPayloads.
* Dependencies have been updated.

### Bugfixes

* Fix `netid_type` method `panic` in case of invalid DevAddr prefix type.
* Fix missing device `search` filter (API).

## v4.3.2

### Improvements

#### Disable v3 compatibility option

This add a configuration option, to enable / disable compatibility with the
latest ChirpStack Gateway Bridge v3 version. If compatibility is not needed,
then add `v4_migrate=false` to the
`[regions.gateway.backend.mqtt]` section (in `region_...toml`). This will
save some bandwidth in the GW <> NS communication. The current default is
`true`, in ChirpStack v4.4+, the default will change to `false` (and thus
it must be explicitly enabled).

#### Other improvements

* Add back web-interface option to download events and frames as JSON.
* Update internal dependencies.

### Bugfixes

* Fix sending to multiple URLs in case one endpoint fails.
* Enable new `tls-rustls` feature of `redis` dependency (fixes Redis TLS issues). ([#170](https://github.com/chirpstack/chirpstack/issues/170))

## v4.3.1

### Improvements

#### LoRa Cloud integration

This adds a _Forward messages on these FPorts to LoRa Cloud_ configuration
option, which based on FPort, let LoRa Cloud Modem & Geolocation Services
automatically handle the payload according.

This also adds a toggle button to make using the reported gateway location
optional for geolocation assistance.

#### Rust SDK features

This improves the SDK `features`, to make it easier to exclude certain
dependencies by disabling features in `Cargo.toml` configuration.

#### Internal dependencies

Internal dependencies have been updated to the latest versions.

### Bugfixes

* Fix typo in `application.proto`. ([#143](https://github.com/chirpstack/chirpstack/pull/143))
* Earlier db initialization + fix unwrap errors. ([#147](https://github.com/chirpstack/chirpstack/issues/147))

## v4.3.0

### Features

### C# SDK

This adds support for generating C# API code. ([#100](https://github.com/chirpstack/chirpstack/pull/100))

### Multicast improvements

#### Add gateways to multicast-group

This makes it possible to explicitly define which gateways will be used for
a given multicast-group. In the gateways overview it is possible to select
one or multiple gateways and then add these gateways to a multicast-group.

#### Multicast scheduling option

This moves the multicast scheduling configuration to the multicast-group
configuration from the `chirpstack.toml` configuration file. Scheduling
options include scheduling based on GPS time or delay based.

#### Split private gateways in private uplink and downlink

This makes it possible to define per tenant if and how gateways can be used
by other tenants. For example in case downlink is set to private, other
tenants will benefit from uplinks received by these gateways, but they will
not be able to send downlinks (ChirpStack will filter these gateways out when
selecting the gateway for scheduling the downlink).

### Improvements

* Implement `get_downlink_data_delay` setting.
* Update internal dependencies.
* Add missing `fPort` validation to avoid enqueue on `fPort=0`.
* Do not overwrite `RxInfo` `location` field with gateway location if it is already set.
* Do not log stats handling `NotFound` errors if `allow_unknown_gateways` is configured.
* Decode FRMPayload mac-commands in device LoRaWAN frames log.
* Show `FCnt` in device event log.

### Bugfixes

* Fix `/api/multcast-groups/...` > `/api/multicast-groups/...` typo in enqueue API URL.
* Fix `gateway_id` is missing errors (in case the uplink was also received by an unknown gateway).
* Fix disabling mac-commands.
* Fix region configuration defaults + use region `id` as fallback for `description` if the latter is missing. ([#120](https://github.com/chirpstack/chirpstack/issues/120))
* Fix API authorization for listing ADR algorithms. ([#112](https://github.com/chirpstack/chirpstack/issues/112))
* Fix US915 downlink channel `min_dr` configuration. ([#115](https://github.com/chirpstack/chirpstack/pull/115))
* Fix incorrect rendering of `integration.mqtt.client.ca_key` and `.ca_cert` in config template. ([#124](https://github.com/chirpstack/chirpstack/pull/124))
* Remove `tls_enabled` config option as it is not actually implemented. ([#128](https://github.com/chirpstack/chirpstack/issues/128))

## v4.2.0

### Features

#### Fix devices to specific region configuration (optional)

This adds a _Region configuration_ option to the device-profile, which lists
the region configurations for the selected _Region_. If a region configuration
is selected, then the device will only be able to work under the selected
configuration. If no region configuration is selected, then the device will
be able to operate under all available region configurations for the selected
region.

#### Java API SDKs

This adds support for Java and Kotlin API SDK code generation. ([#64](https://github.com/chirpstack/chirpstack/pull/64))

### Improvements

* Add `description` configuration option per region configuration.
* Change `name` to `id` within the region configuration (`name` will be used as fallback option).
* Make gateway state in UI consistent and make expected stats interval configurable. ([#76](https://github.com/chirpstack/chirpstack/issues/76))
* Add Python type information to Python API SDK code. ([#68](https://github.com/chirpstack/chirpstack/pull/68))
* Add back `crc_status` field to `UplinkRxInfo` message (the status will be reported as no CRC until the ChirpStack Gateway Bridge, ChirpStack Concentratord and / or ChirpStack MQTT Forwarder have been updated).
* Add back Class-B ping-slot parameters to the device-profile.
* Update Class-B ping-slot data-rate configuration in examples.
* Remove separate gateway topic config and move it into single `topic_prefix` configuration.
* Reset internally stored channels to default on `ADRACKReq` uplink to avoid out-of-sync channel configuration on device.
* Update internal dependencies.

### Bugfixes

* Fix hiding _Delete device_ option if the user has no permissions to perform this action. ([#71](https://github.com/chirpstack/chirpstack/issues/71))
* Fix not recording device metrics if auto-detect of measurements is disabled. ([#94](https://github.com/chirpstack/chirpstack/issues/94))

## v4.1.3

### Bugfixes

* Fix Redis `key_prefix` configuration. While this value could be configured, it was not applied to the generated keys.
* Fix header `z-index` issue in UI. This was causing the dropdowns to render partly behind the header.

## v4.1.2

### Improvement / bugfix

* Do not wait for integrations to finish before sending downlink.

## v4.1.1

### Improvements

* Update JS API dependencies to latest versions.
* Replace relative paths in Rust API build to absolute. ([#69](https://github.com/chirpstack/chirpstack/pull/69))

### Bugfixes

* Fix setting the full frame-counter to the uplink frame after resolving the device-session (this can affect payload decryption).

## v4.1.0

### Features

#### API request logging

This feature logs API requests to Redis Streams. This enables external services
to monitor for example device create, update and deletes by reading from the
Redis Streams. A code-example can be found [here](https://github.com/chirpstack/chirpstack/blob/master/examples/request_log/go/main.go).

#### Uplink logging for unknown devices

While the feature to log frames was already present, it was not possible
to only read uplink frames of devices that are unknown. This extends the frame
logging feature to also log uplinks for unknown devices, in which case DevEUI
`0000000000000000` is used. A code-example for reading the frame log can be
found [here](https://github.com/chirpstack/chirpstack/blob/master/examples/frame_log/go/main.go).

#### Event logging

A code-example to read event logs from the Redis Streams was added. It can
be found [here](https://github.com/chirpstack/chirpstack/blob/master/examples/event_log/go/main.go).

### Improvements

* Make `metadata` fields in gateway messages consistent.
* Emit all fields for JSON integration messages, even if they are their default values. ([#63](https://github.com/chirpstack/chirpstack/pull/63))

### Bugfixes

* Fix Redis pipelined commands in case Redis Cluster is configured.
* Fix UI notifications z-index.

### Other changes

* The `import-ttn-lorawan-devices` sub-command has been renamed to `import-legacy-lorawan-devices-repository`.

## v4.0.5

### Improvements

* Add `keep_alive_interval` for MQTT configuration.

### Bugfixes

* Fix incorrect splitting of multiple URLs in HTTP integration. ([#62](https://github.com/chirpstack/chirpstack/issues/62))
* Fix missing ThingsBoard location and status telemetry.
* Send ThingsBoard telemetry (fPort, fCnt, ...) even in case there is no decoded payload.
* Fix LeafletJS controls floating over header (UI).

## v4.0.4

### Bugfixes

* Fix coding-rate for LoRaWAN 2.4GHz. ([#51](https://github.com/chirpstack/chirpstack/issues/51))
* Fix not removing of queue-item after it was sent, if payload did not fit RX2.

## v4.0.3

### Improvements

* Add code example for reading frame-logs from Redis Streams.
* Add missing metadata logging and add code example for reading metadata from Redis Streams.
* Add `dev_addr_prefix` configuration. ([#49](https://github.com/chirpstack/chirpstack/issues/49))
* Make it possible to enable / disable auto-detection of metrics in device-profile. ([#42](https://github.com/chirpstack/chirpstack/issues/42))
* Add Redis config examples for username and password configuration. ([#54](https://github.com/chirpstack/chirpstack/issues/54))
* Add metadata tab to gateway configuration in UI.

### Bugfixes

* Use `trust_store` instead of `ca_path` for MQTT integration. ([#47](https://github.com/chirpstack/chirpstack/pull/47))
* Fix _Cannot serialize NaN as google.protobuf.Value.number_value_ error.
* Fix logout URL in case OIDC is enabled.
* Fix `java_outer_classname` in `tenant.proto`. ([#55](https://github.com/chirpstack/chirpstack/pull/55))
* Fix closing and detecting if eventlog channel is closed.
* Fix metrics interval calculation on daily aggregate in case of DST to non-DST timezone change.

## v4.0.2

### Features

* A new menu option Regions was added to the web-interface, exposing per region
  information.

### Improvements

* Internal dependencies were updated.

### Bugfixes

* Fix Wi-Fi geolocation issue in LoRa Cloud integration.
* Fix missing per data-rate stats in gateway dashboard.
* Fix terminating stream loop (and releasing of Redis connection) on client disconnect. ([#40](https://github.com/chirpstack/chirpstack/issues/40))
* Fix rendering `client_cert_lifetime` rendering in configuration template (mqtt integration).

## v4.0.1

Not released.

## v4.0.0

After many months of development and testing, we are really excited to share
ChirpStack v4.

The aim of ChirpStack v4 is to make it significantly easier to
setup and use ChirpStack, compared to the previous version.
One of the major changes that you will notice is that the ChirpStack Network
Server and ChirpStack Application Server have been merged into a single
component. Over the years we have seen many issues reported
on the forum and GitHub, related to setting up and connecting both services.
ChirpStack v4 also provides multi-region support out-of-the-box, including
region configuration. No longer it is needed to define your own configuration
file or setup multiple ChirpStack Network Server instances to serve multiple
regions simultaneously.

A big thank you to the ChirpStack community for supporting and contributing
to the ChirpStack project! Please find below a breakdown of all the new features
and changes that v4 brings.

### Main features and changes

#### Multi-region support

ChirpStack v4 adds multi-region support, removing the need to setup multiple
ChirpStack Network Server instances. Configuration files are included for the
common regions (as defined by the LoRa Alliance), which should help getting
started with ChirpStack.

Each enabled region has its own gateway backend, making it possible to use one
or multiple MQTT brokers for the different gateway pools. In case a single
MQTT broker is used, each backend must be configured with its own MQTT topic prefix
(e.g. `eu868/gateway/...`).

ChirpStack v4 also supports multiple configurations of the same region, e.g.
to configure a US915 for 8 channels and to configure a US915 band for 16
channels.

#### TTN device repository support

ChirpStack v4 adds support for importing the TTN [LoRaWAN Devices](https://github.com/TheThingsNetwork/lorawan-devices)
repository as device-profile templates, including codec functions if these are
defined in this repository.

#### Device metrics dashboard

In the device-profile template and / or device-profile, it is possible to define
the measurements that are exposed by the device in the decoded payload. Once
defined, ChirpStack will automatically aggregate and store this data. These
metrics can be viewed in the web-interface on the device dashboard.

#### Configuration

##### Directory

Instead of using a single configuration file (e.g. `chirpstack-network-server.toml`),
ChirpStack makes use of a configuration directory such that the configuration
can be split in multiple files. By default you will find a single `chirpstack.toml`
configuration file, and many `region_...toml` configuration files, split
by region.

##### Environment variables

ChirpStack v4 removes the TOML hierarchy to environment variable mapping.
Instead it allows you to define the variables like `$MY_CONFIGURATION`, which
will get automatically substituted when an environment variable is found with
the name `MY_CONFIGURATION`.

#### API

The REST interface that was present in ChirpStack Application Server v3 has
been removed, in favor of the gRPC API interface (please see the `api/` folder
of the repository for the API definitions). However, a gRPC to REST interface
bridge component will be provided as a separate service. Please note that
in v3, this bridge component was embedded and REST interface calls were
internally translated to gRPC calls. Therefore, gRPC was always recommended
interface to use.

The ChirpStack v4 gRPC API does support server reflection, making it possible
to use for example [gRPC UI](https://github.com/fullstorydev/grpcui) or
[BloomRPC](https://github.com/bloomrpc/bloomrpc) as development interface.

#### ChirpStack Gateway Bridge v3 compatibility

ChirpStack v4 is fully compatibility with the latest version of ChirpStack
Gateway Bridge v3. This should help migrating from v3 to v4. Please note that
the ChirpStack Gateway Bridge **must** be configured with the `protobuf`
marshaler.

#### UI rewrite

ChirpStack v4 contains a rewrite of the ChirpStack Application Server v3 UI.
The new UI aims to be more user-friendly. Under the hood the API interface
has been ported to gRPC-web and all code has been ported to Typescript.

### Other changes

#### Passive Roaming improvements

The implementation of Passive Roaming has been improved, adding support for
appending `/sns` and `/fns` server endpoint suffixes. The usage of this
suffix is not specified in the Backend Interfaces specification, but is
required by some other network-server implementations.

#### UUID identifiers

All identifiers that are exposed have been changed to UUID. Previously most
identifiers (e.g. users, applications...) were incremental integers. In case
ChirpStack is setup as multi-tenant instance, this could expose some information
about the number of clients on the network. The migration script (see below)
will migrate these integers by converting these as strings, prefixed with
zeros in the UUID format. E.g. ID `123` would be converted to the UUID string
`00000000-0000-0000-0000-000000000123`.

#### String identifiers in API

All binary identifiers have been changed to string type in the API. While
binary fields are more efficient, these were confusing when encoded as JSON
as the Protobuf to JSON mapping uses base64 encoding for binary fields.
For example, a Gateway ID `0102030405060708` was encoded as `AQIDBAUGBwg=`
in JSON.

#### JavaScript codec engine

The JavaScript codec engine is based on QuickJS, which is an embeddable
JavaScript engine which supports the ES2020 specification. 

#### API interface changes

While the structure of API messages is roughly the same as the ChirpStack
Application Server API interface, some small changes have been made.

#### Integration events

The integration event messages have been restructured for better consistency.
Each event message has a `deviceInfo` field which holds device-related
information (tenant id & name, application id & name, device-profile id &
name, device EUI & name and tags).

### Development changes

#### Single repository

ChirpStack v4 will make it a lot easier to make customizations, especially
when API changes are involved, as API definitions are no longer separated
from the code. In v3 these definitions were moved to an external repository
to avoid cross dependencies.

#### Rust

For ChirpStack v4, it was decided to use Rust rather than Go. This was not
an easy choice and the arguments for this decision are debatable. However, 
as most code was touched during the ChirpStack Application Server and ChirpStack
Network Server merge, it was the only moment to re-consider this. The Rust 
memory management prevents many memory related pitfalls and helps catching
bugs at compile time rather than runtime.

### Migrating from v3 to v4

The recommended way to migrate from v3 to ChirpStack v4 is to create a new
PostgreSQL and Redis database and to use the [ChirpStack v3 to v4](https://github.com/chirpstack/chirpstack-v3-to-v4)
migration script. This script will **copy** all the data from the "old"
into the "new" database. While the script does not make any modifications
to the old database, it is always recommended to make a backup first.
