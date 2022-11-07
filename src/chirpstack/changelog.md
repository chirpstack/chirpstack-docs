# Changelog

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
