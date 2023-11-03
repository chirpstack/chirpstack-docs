# ChirpStack v3 to v4 migration

The purpose of this document is to describe the steps needed to migrate
data from ChirpStack v3 to ChirpStack v4.

<!-- toc -->

## Preparation

### ChirpStack Application Server

Please make sure to upgrade to the latest v3 version. It is recommended to
make a database backup.

### ChirpStack Network Server

Please make sure to upgrade to the latest v3 version. It is recommended to
make a database backup.

### ChirpStack Gateway Bridge

Please make sure to upgrade to the latest v3 version. ChirpStack Gateway Bridge
v3.14.0+ is compatible with both ChirpStack v3 and ChirpStack v4.

ChirpStack v4 supports the latest ChirpStack Gateway Bridge v3, but it does
require the `marshaler` option to be set to `protobuf` (this is the default).

### ChirpStack (v4)

Please make sure that the latest ChirpStack v4 version has been installed
and is up and running. You must create a new PostgreSQL database, while not
strictly required, it is also recommended to create a new Redis server.

**Note:** Do not create any gateways and devices yet, as the migration script
will delete all objects from the ChirpStack v4 database, before copying over
the data from the v3 databases.

## Data migration

The ChirpStack project provides a migration script which will copy data from
the ChirpStack Application Server and ChirpStack Network Server databases:
[https://github.com/chirpstack/chirpstack-v3-to-v4](https://github.com/chirpstack/chirpstack-v3-to-v4).
Please see the [Releases](https://github.com/chirpstack/chirpstack-v3-to-v4/releases)
for pre-compiled binaries.

This script requires at least three configuration files (see below):

1. `chirpstack-application-server.toml` (ChirpStack Application Server v3)
1. `chirpstack-network-server.toml` (ChirpStack Network Server v3)
1. `chirpstack.toml` (ChirpStack v4)

Example:

```
./chirpstack-v3-to-v4 \
	--as-config-file /etc/chirpstack-application-server/chirpstack-application-server.toml \
	--ns-config-file /etc/chirpstack-network-server/chirpstack-network-server.toml \
	--cs-config-file /etc/chirpstack/chirpstack.toml
```

It will iterate over all the items in the v3 databases (ChirpStack Application
Server and ChirpStack Network Server), and will write these into the
ChirpStack v4 database.

**Important:** In case you have multiple ChirpStack Network Server instances setup for
multi-region support, you must repeat the `--ns-config-file` option for each
instance.

Example:

```
./chirpstack-v3-to-v4 \
	--as-config-file /etc/chirpstack-application-server/chirpstack-application-server.toml \
	--ns-config-file /etc/chirpstack-network-server/chirpstack-network-server-eu868.toml \
	--ns-config-file /etc/chirpstack-network-server/chirpstack-network-server-us915.toml \
	--cs-config-file /etc/chirpstack/chirpstack.toml
```

## Configuration

You must confirm that the ChirpStack Gateway Bridge and MQTT topics configured
in the ChirpStack region configuration are aligned.

The default configuration that is provided by ChirpStack v4 assumes a region
name as MQTT topic prefix (e.g. `eu868/gateway/...` or `us915_0/gateway/...` for
the US915 region with channels 0 - 7). Please refer to the [Multi region](./chirpstack/features/multi-region.md)
documentation for more information.

## Validation

To make validate that the migration was successful, validate the following
steps in the ChirpStack v4 web-interface:

- [ ] Gateways last seen status is periodically updated
- [ ] Gateway frames are being reported
- [ ] Device frames are appearing in the web-interface
- [ ] Device events are appearing in the web-interface
