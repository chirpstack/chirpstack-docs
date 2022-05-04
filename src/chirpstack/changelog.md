# Changelog

## v4.0.0 (in development)

After many months of development, We are really excited to share ChirpStack v4. The
idea of ChirpStack v4 was born based on experience using ChirpStack for
various clients, as well as many discussion with community members
and recurring issues that were reported on the forum.

ChirpStack v4 merges the (v3) ChirpStack Application and ChirpStack Network Server,
adding support for multi-region without the need to setup multiple ChirpStack
instances. ChirpStack is a new component, which replaces the ChirpStack Application
Server and ChirpStack Network Server.

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
JavaScript engine which supports the ES2020 specification. In the near future
support for the NodeJS [Buffer class](https://nodejs.org/api/buffer.html)
will be added.

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

### Still missing

There are a few features still missing, that were present in v3. These will be
added shortly. These are:

* Global integrations (only MQTT is currently implemented)
* Passive roaming
* REST API bridge
* Support for GCP Cloud IoT Core and Azure IoT Hub gateway interfaces

#### Rust

For ChirpStack v4, it was decided to use Rust rather than Go. This was not
an easy choice and the arguments for this decision are debatable. However, 
as most code was touched during the ChirpStack Application Server and ChirpStack
Network Server merge, it was the only moment to re-consider this. The Rust 
memory management prevents many memory related pitfalls and with that potential
bugs, as these can be catched at compile-time rather than during runtime.

### Migrating from v3 to v4

The recommended way to migrate from v3 to ChirpStack v4 is to create a new
PostgreSQL and Redis database and to use the [ChirpStack v3 to v4](https://github.com/chirpstack/chirpstack-v3-to-v4)
migration script. This script will **copy** all the data from the "old"
into the "new" database. While the script does not make any modifications
to the old database, it is always recommended to make a backup first.

The script is called with at least three configuration files; the
`chirpstack-network-server.toml` (v3), `chirpstack-application-server.toml` (v3)
and the `chirpstack.toml` (v4) configuration files. In case you have setup
multiple ChirpStack Network Server regions, you can repeat the ChirpStack
Network Server configuration file argument for each region.
