# v4 breaking changes

ChirpStack v4 is a new major version that brings many improvements over the
ChirpStack v3. Before you migrate your ChirpStack v3 instance to ChirpStack
v4, please be aware that there are changes that are not backwards compatible!

## Gateway to ChirpStack MQTT

The ChirpStack Gateway Bridge v3.14.0+ version **is** compatible with both
ChirpStack v3 and ChirpStack v4, as long as the `protobuf` marshaler is
configured (default value). This to eases the migration, as the published
MQTT data can be consumed by both ChirpStack v3 and ChirpStack v4.

Please note that the default ChirpStack v4 configuration contains the region
as MQTT prefix in the configuration. This means that you need to either:

* Update the ChirpStack Gateway Bridge topic configurations to include the
  region in the templates, or
* Update the ChirpStack region configuration, removing the topic prefix from
  the topic templates, or
* Implement a service which re-publishes MQTT messages, resulting in messages
  being published both with and without the topic prefix.

Support for using AMQP, GCP Pub/Sub or Azure IoT Hub as gateway backend is not
supported by the initial ChirpStack v4 release, but will be added in a future
version.

## gRPC API

If using the gRPC API, you must upgrade the ChirpStack gRPC API client-libraries
to v4. In general, the structure of the gRPC is mostly identical to the v3 gRPC
API definition, but there are minor changes and renames. The v3 client-code is
not compatible with the v4 API.

## REST API

ChirpStack v3 provided a REST API on top of the gRPC API. This means that REST
request were internally translated to gRPC calls and responses back to REST.
For this the [gRPC Gateway](https://github.com/grpc-ecosystem/grpc-gateway)
project was used internally. Historically, this component was embedded because
of the web-interface. At the time of this implementation, gRPC could not be used
within the browser. ChirpStack v4 uses [gRPC Web](https://github.com/grpc/grpc-web)
to interface with the gRPC API.

With the recommendation to use the gRPC API directly (to avoid the REST to gRPC
translation for each API request) and the web-interface using gRPC Web,
ChirpStack v4 no longer provides an embedded REST API interface. However, the
ChirpStack project does provide a component to provide a REST interface on top
of the ChirpStack gRPC API. Please refer to
[https://github.com/chirpstack/chirpstack-rest-api](https://github.com/chirpstack/chirpstack-rest-api).

## Integration event messages

The integration event messages have been refactored to provide the same base
information for all events. This includes the tenant, application, device-profile
and device information (names and ids). Decoded uplink events are directly
represented as object, before this was encoded as JSON string. Some fields
have been renamed.

## Identifiers

ChirpStack v3 contained numeric identifiers for organizations, users,
applications etc. On public instances this could expose some information (e.g. number of users).
In ChirpStack v4 identifiers have been changed to UUIDs. The v3 to v4 migration
utility will migrate numeric identifiers to UUIDs by starting with an empty
UUID (`00000000-0000-0000-0000-000000000000`) and setting the last characters
to the numeric value.

For example, the numeric id `1234` will be migrated to 
`00000000-0000-0000-0000-000000001234`.

## Network controller interface

ChirpStack v3 provided a `NetworkControllerService` which could be implemented
to handle uplink and downlink meta-data. This interface has been removed. This
meta-data can now be logged into a [Redis Stream](https://redis.io/docs/manual/data-types/streams/)
by setting the `meta_log_max_history` to a non-zero value. This log can then be
consumed by an external application.

## Custom ADR algorithms

Custom ADR algorithms must now be implemented in JavaScript. In ChirpStack v3
algorithms needed to be implemented in Go and then compiled using the Go
toolchain.

## Gateway discovery

The gateway discovery feature has been removed. The problem is that this is not
supported by the Basics Station packet-forwarder (see [#26](https://github.com/lorabasics/basicstation/issues/26)).
As well, some gateways have RF filters on the RX and TX path to filter out noise.
In regions where the uplink frequencies are not equal to the downlink frequencies,
these filters would prevent proper gateway to gateway communication.
