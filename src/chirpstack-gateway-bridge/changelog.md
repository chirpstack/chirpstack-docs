# Changelog

## v4.1.1

### Bugfixes

* Fix gateway re-connect issue caused by [#238](https://github.com/chirpstack/chirpstack-gateway-bridge/pull/238).

## v4.1.0

### Removal

#### Concentratord backend

This release removes the ChirpStack Concentratord backend. For a long time, the
ChirpStack MQTT Forwarder has been a better alternative to the ChirpStack
Gateway Bridge for on-gateway installation. For more information, please see
[ChirpStack MQTT Forwarder](https://www.chirpstack.io/docs/chirpstack-mqtt-forwarder/).

### Improvements

* Make UDP connection timeout configurable. ([#238](https://github.com/chirpstack/chirpstack-gateway-bridge/pull/238))
* Add TLS support by Reverse-Proxy for Basics Station. ([#241](https://github.com/chirpstack/chirpstack-gateway-bridge/pull/241))
* Make Semtech UDP cache expiration and cleanup interval configurable. ([#244](https://github.com/chirpstack/chirpstack-gateway-bridge/pull/244))

### Bugfixes

* Add missing multiplier to convert to nanoseconds. ([#234](https://github.com/chirpstack/chirpstack-gateway-bridge/pull/234))
* Fix typo in Prometheus metric name. ([#247](https://github.com/chirpstack/chirpstack-gateway-bridge/pull/247))

## v4.0.11

### Improvements

* Update internal dependencies.

### Bugfixes

* Fix firewall security issue on Kerlink gateways.

## v4.0.10

### Improvements

* Add configuration option to print log in JSON format.

## v4.0.9

### Bugfixes

* Prevent duplicate MQTT connections. ([#229](https://github.com/chirpstack/chirpstack-gateway-bridge/pull/229))

## v4.0.8

### Bugfixes

* Fix cached RouterConfig response (Basics Station).

## v4.0.7

### Improvements

* Add `MuxTime` to downlink messages (Basics Station) and remove `timesync_interval` default.
* Update internal dependencies.

## v4.0.6

### Improvements

* Include location in `UplinkRxInfo` if Semtech UDP payload contains both `rxpk` and `stat` (with location).

## v4.0.5

### Bugfixes

* Fix loading config from default paths after merging [#199](https://github.com/chirpstack/chirpstack-gateway-bridge/pull/199).

## v4.0.4

### Improvements

* Support for multiple configuration files. ([#199](https://github.com/chirpstack/chirpstack-gateway-bridge/pull/199))
* Add filters to Concentratord backend. ([#218](https://github.com/chirpstack/chirpstack-gateway-bridge/pull/218))
* Add CRC status back to UplinkRxInfo.

## v4.0.3

### Bugfixes

* Add missing coding-rates for TXPK CodR. ([#217](https://github.com/chirpstack/chirpstack-gateway-bridge/issues/217))

## v4.0.2

### Features

#### `meta` in rxpk and stat

If using the Semtech UDP Packet Forwarder backend, it is possible to expose
additional key / value metadata using an optional `meta` key in either the 
rxpk or stat package. The value of `meta` must be an object with key (string)
/ value (string) data if present.

### Improvements

* Update Go compiler to v1.19.3.
* Update Alpine base container to v3.17.0.
* Update prometheus/client_golang to v1.14.0.
* Make `metadata` field in `chirpstack_api` consistent.

## v4.0.1

### Improvements

* Expose ack-rate (Semtech UDP protocol) through Prometheus metrics.
* Update Go compiler toolchain to latest version.

### Bugfixes

* Update Multitech Conduit packaging to work Monit for process monitoring (included in latest mPower).

## v4.0.0

### Features

This release implements the ChirpStack v4 API interface. Please note that this
version is not compatible with ChirpStack v3.

## v3.14.3

### Bugfixes

* Fix incorrect handling of MQTT token timeout, causing re-connect to not work in some cases. ([#210](https://github.com/chirpstack/chirpstack-gateway-bridge/issues/210))

## v3.14.2

### Bugfixes

* Basics Station: Prefer `gpstime` over `rxtime` in case both are available.

## v3.14.1

### Improvements

* Update dependencies.
* Update to Go 1.18.

## v3.14.0

### Features

* Add support for periodic time-sync requests (Basics Station).

### Improvements

* Make it possible to use 32bit downlink tokens instead of 16bit only (will be used by ChirpStack v4).

## v3.13.3

### Improvements

* Update MQTT client to latest version.
* Add timeout to MQTT token wait.
* Basics Station backend:
	* Add `WORKAROUND_IGNORE_RX_TIME` flag (environment variable) to disable the parsing of `rxtime`
	  in cases where the gateway does not provide a correct `rxtime` value.

## v3.13.2

### Improvements

* Expose SX1302 / SX1303 `ftime` as plain fine-timestamp if available.
* Consolidate Tektelic `.ipk` packages to single `kona` architecture.

## v3.13.1

### Bugfixes

* Fix stats Hz to kHz conversion for Concentratord backend.

## v3.13.0

### Features

#### Gateway metrics

This implements the aggregation of additional gateway metrics:

* Uplink & downlink count split by frequency
* Uplink & downlink count split by modulation parameters
* Downlink (attempts) split by status

### Improvements

* Docker images added for ARM v6, v7 and ARM64. ([#198](https://github.com/brocaar/chirpstack-gateway-bridge/pull/198))

### Bugfixes

* Correct the command_url logged. ([#197](https://github.com/brocaar/chirpstack-gateway-bridge/pull/197))

## v3.12.0

### Features

* Implement LR-FHSS modulation type.

### Improvements

* Make Kerlink `.ipk` packaging generic for all KerOS models.

### Bugfixes

* Fix `null` time when GPS is not available in gateway (Basics Station). ([#196](https://github.com/brocaar/chirpstack-gateway-bridge/pull/196))

## v3.11.0

### Features

#### Connection state messages

This publishes a `conn` state message, indicating if a gateway is online or
offline. State messages are published as retained so that a MQTT client will
immediately receive the latest state of the gateway(s) after subscribing to
the `conn` state MQTT topic. The default MQTT topic for state messages is
`gateway/ID/state/STATE` where `ID` is the ID of the gateway and `STATE` is
the state (in this case `conn`). This topic is configurable in the
[Configuration](install/config.md).

### Improvements

* When `terminate_on_connect_error` is enabled, terminate on connection lost. ([#184](https://github.com/brocaar/chirpstack-gateway-bridge/pull/184))
* Do not run as root (Docker). ([#185](https://github.com/brocaar/chirpstack-gateway-bridge/pull/185))
* Add ns precision to log output timestamp.

### Bugfixes

* Add websocket connection mutex to avoid concurrent writes. ([#191](https://github.com/brocaar/chirpstack-gateway-bridge/pull/191/))

## v3.10.0

### Features

* Basic Station timesync request/response support. ([#179](https://github.com/brocaar/chirpstack-gateway-bridge/issues/179))

### Improvements

* Refactor backend <> integration communication code.
* Refactor the way how MQTT (un)subscribe is implemented.
* Make MQTT keepalive interval configurable. ([#176](https://github.com/brocaar/chirpstack-gateway-bridge/pull/176))
* Add new Azure CA certificates. ([#178](https://github.com/brocaar/chirpstack-gateway-bridge/issues/178))
* Implement paho mqtt client logging. ([#182](https://github.com/brocaar/chirpstack-gateway-bridge/pull/182))

## v3.9.2

### Improvements

* Provide uplink / downlink stats when using the BasicStation backend.

## v3.9.1

### Bugfixes

* Downlinks are now visible in the web-interface when using in combination with the BasicStation.
* RX2 parameters are now set for Class-A downlinks when using in combination with the BasicStation.

## v3.9.0

### Features

#### Multi-downlink commands and ACKs

This multi-downlink feature allows a Network Server to send multiple downlink
opportunities at once, e.g. one item for RX1 and one for RX2. The gateway will
send at most one downlink item from the list. The ChirpStack Gateway Bridge will
respond with a list of ACK items, indicating which item item was emitted, which
item(s) failed and which item(s) were ignored (because the previous item was
emitted).

#### Channel re-configuration

This implements channel re-configuration for the [ChirpStack Concentratord](https://github.com/brocaar/chirpstack-concentratord)
backend over MQTT. This also removes the channel re-configuration for the Semtech
UDP Packet Forwarder. While this feature had been removed from the documentation
code references were still present. As there are different implementations of
the Semtech UDP Packet Forwarder, there are also different flavors of
configuration files.

The ChirpStack Concentratord provides a unified API to different types of
gateways which makes it possible to implement channel re-configuration for
this backend.

#### Kerlink iStation

This release provides IPK packages that can be installed on a Kerlink iStation
gateway.

### Improvements

* Split multi-line meta-data returned by a command only once on delimiter.
* Forward meta-data provided by a Packet Forwarder backend.

## v3.8.0

### Features

#### Multi-line meta-data commands

This makes it possible to let a meta-data command return multiple lines of
values. In case of multi-line output, ChirpStack Gateway Bridge will split
each line by a configured delimiter to obtain the key / value pairs.

### Bugfixes

* Fix race condition in Semtech UDP Packet Forwarder backend. ([#162](https://github.com/brocaar/chirpstack-gateway-bridge))

## v3.7.1

### Bugfixes

* Fix duplicated configuration in the Tektelic config example.
* Fix / improve ZMQ socket re-connect for Concentratord backend.

## v3.7.0

### Features

#### Concentratord backend

The [Concentratord](https://www.chirpstack.io/chirpstack-gateway-bridge/backends/chirpstack-concentratord/) backend is compatible with the
[ChirpStack Concentratord](https://github.com/brocaar/chirpstack-concentratord)
daemon. Please note that this backend is experimental.

#### Syslog output

When `log_to_syslog` is enabled in the configuration file, the log output will
be written to syslog.

### Improvements

* Change MQTT max. reconnect interval default value to 1 minute.
* Refactor MQTT subscribe / unsubscribe logic.

### Bugfixes

* Use DevEUI object length instead of JoinEUI. ([#154](https://github.com/brocaar/chirpstack-gateway-bridge/pull/154))
* Remove default values for BasicStation filters. ([#157](https://github.com/brocaar/chirpstack-gateway-bridge/issues/157))
* Fix BasicStation websocket ping loop. ([#156](https://github.com/brocaar/chirpstack-gateway-bridge/issues/156))

## v3.6.0

### Features

#### Raw packet-forwarder events / commands

Using the raw packet-forwarder events and commands, it is possible to use the
[BasicStation Remote Command](https://doc.sm.tc/station/tcproto.html#remote-command)
and [BasicStation Remote Shell](https://doc.sm.tc/station/tcproto.html#remote-shell)
features.

#### Terminate on connect error

The `terminate_on_connect_error` will terminate the ChirpStack Gateway Bridge
process when it can't connect to the MQTT broker. In some cases this is
preferred over the retry loop.

### Improvements

#### FSK FDev

Before the frequency deviation was always set to the FSK `datarate / 2`. It
is now possible to set it manually (falling back on `datarate / 2` when unset).

## v3.5.0

### Features

#### MIPS builds

This is the first release providing binaries compiled for MIPS. These binaries
are compiled with `GOMIPS=softfloat` and are compressed with `upx` as MIPS
based gateways have usually limited storage space available. ([#65](https://github.com/brocaar/chirpstack-gateway-bridge/issues/65))

#### RPM packaging

This is the first release providing .rpm packages for CentOS and RedHat. ([#145](https://github.com/brocaar/chirpstack-gateway-bridge/pull/145))

### Improvements

#### Environment variable configuration

The usage of `.` in environment variables for configuration has been deprecated
(it will continue to work, but will log a warning). Instead of `.`, use a
double underscore (`__`). ([#144](https://github.com/brocaar/chirpstack-gateway-bridge/pull/144))

#### Multiple MQTT servers

This release adds support for the configuration of multiple MQTT servers. ([#141](https://github.com/brocaar/chirpstack-gateway-bridge/pull/141))

## v3.4.1

### Bugfixes

* Fixes init stop script which could cause the ChirpStack Gateway Bridge to not properly stop or restart.

## v3.4.0

This release renames LoRa Gateway Bridge to ChirpStack Gateway Bridge.
See the [Rename Announcement](https://www.chirpstack.io/r/rename-announcement) for more information.

### Bugfixes

* Fix nested read lock. ([#139](https://github.com/brocaar/lora-gateway-bridge/issues/138))

## v3.3.0

### Features

#### Message IDs

This release implements unique message IDs (both for events and commands) that
can be used for correlation or logging purposes.

#### BasicStation channel-plan configuration

To make it easier to work with the [BasicStation](https://doc.sm.tc/station/)
Packet Forwarder, this release makes it possible to configure the channel-plan
inside the LoRa Gateway Bridge configuration file. It is no longer needed to
configure a Gateway Profile within LoRa App Server. This means that it is also
possible to use the LoRa Gateway Bridge with BasicStation gateways, without
the need of installing LoRa Server and LoRa App Server.

## v3.2.1

### Bugfixes

* Fix NetID 3 & 4 filter according to the [errata](https://lora-alliance.org/resource-hub/nwkid-length-fix-type-3-and-type-4-netids-errata-lorawan-backend-10-specification) published by the LoRa Alliance.
* Fix Basic Station bandwidth for LoRa Std channel (from kHz to Hz). ([#130](https://github.com/brocaar/chirpstack-gateway-bridge/pull/130))

## v3.2.0

### Features

#### NetID and JoinEUI filters

Configuration options have been implemented to filter uplink messages on netID
and JoinEUI. This makes it possible to ignore messages from neighboring
to save on bandwidth usage (e.g. when the gateway is using a cellular backhaul).

#### Execute commands

This feature makes it possible to execute commands on the gateway (when the
LoRa Gateway Bridge is running on the gateway). Note: commands must be
pre-configured in the LoRa Gateway Bridge configuration file.

### Improvements

#### Basic Station backend

* Verify Common Name when using client certificates. ([#129](https://github.com/brocaar/chirpstack-gateway-bridge/pull/129))

## v3.1.0

### Features

#### Azure IoT Hub X.509

X.509 certificate authentication is added to the `azure_iot_hub` MQTT
authentication option.

### Improvements

#### Prometheus metrics

The Prometheus metrics have been improved for consistency.
Documentation has been updated to document the metrics that are available.

#### Max reconnect interval

The MQTT max. reconnect interval is now a global MQTT configuration and can
be used regardless of the MQTT authentication type.

#### Lat / lon / alt = 0

When the latitude, longitude and altitude are all three `0`, then LoRa Gateway
Bridge assumes there is no GPS module / location available and it will not
expose a location in the stats message.

### Bugfixes

* The CRC check has been fixed for the Semtech UDP backend.
* Fix message buffering for QoS > 0 (on manual reconnect, the Paho library discards messages when offline).

### Deprecated

#### Managed gateway configuration

Documentation (and configuration) references to the managed
packet-forwarder configuration have been removed. The feature itself is still
available and will stay available until the next major release.
Please refer to [this forum discussion](https://forum.chirpstack.io/t/who-is-using-the-gateway-profile-how-are-you-using-it/5091)
for more background.

## v3.0.1

### Bugfixes

* Fix acquiring double read-lock (which could result in a deadlock). [#119](https://github.com/brocaar/chirpstack-gateway-bridge/pull/119)

## v3.0.0

### Features

#### Basic Station packet-forwarder support

The LoRa Gateway Bridge has been refactored to support multiple packet-forwarder
backends. Next to the Semtech UDP packet-forwarder, support has been added to
support the Basic Station packet-forwarder. This backend implements the 
[Basic Station LNS protocol](https://doc.sm.tc/station/tcproto.html).

#### Updated payload formats

The uplink payload contains a `context` field, used to store gateway
specific context data (like the internal counter).

The downlink frame contains a `timing` field which can be either
`IMMEDIATELY`, `DELAY` or `GPS_EPOCH`. Based on the `timing` value, an
additional object must be given with the additional timing information.
Refer to [Commands](https://www.chirpstack.io/chirpstack-gateway-bridge/payloads/commands/)
for more details.

#### Azure IoT Hub integration

LoRa Gateway Bridge now has support to connect the [Azure IoT Hub](https://azure.microsoft.com/en-us/services/iot-hub/)
MQTT interface. Like with the Google Cloud Platform IoT Core integration this
makes it easier to scale up and manage gateway credentials using the IoT Hub
console (or API).

#### Fake RX time

In case the gateway does not have a GPS module, the RX time would would be
left blank. This feature makes it possible to use the system time as a fallback.
([#109](https://github.com/brocaar/chirpstack-gateway-bridge/pull/109))

#### Custom gateway meta-data

This feature makes it possible to expose additional meta-data in the gateway
stats. Meta-data can either be static or dynamic (executing external commands).
The latter option allows to for example read and expose the gateway temperature,
humidity, ...

### Upgrading

LoRa Gateway Bridge v3.0.0 include a couple of changes that are not backwards
compatible. You need to re-generate the configuration file and update it where
needed. LoRa Gateway Bridge v3 is compatible with LoRa Server v2.7+.
Below a summary:

#### MQTT topics

The MQTT topic configuration has been updated from:

```toml
{% raw %}
uplink_topic_template="gateway/{{ .MAC }}/rx"
downlink_topic_template="gateway/{{ .MAC }}/tx"
stats_topic_template="gateway/{{ .MAC }}/stats"
ack_topic_template="gateway/{{ .MAC }}/ack"
config_topic_template="gateway/{{ .MAC }}/config"
{% endraw %}
```

To:

```toml
{% raw %}
event_topic_template="gateway/{{ .GatewayID }}/event/{{ .EventType }}"
command_topic_template="gateway/{{ .GatewayID }}/command/#"
{% endraw %}
```

Event types are: `up`, `stats` and `ack`.
Commands are: `down` and `config`.

Please note that LoRa Gateway Bridge v3 is compatible with LoRa Server v2.7+,
but you will need to update the MQTT topics in your `loraserver.toml` file.
Example snippet:

```toml
{% raw %}
uplink_topic_template="gateway/+/event/up"
stats_topic_template="gateway/+/event/stats"
ack_topic_template="gateway/+/event/ack"
downlink_topic_template="gateway/{{ .MAC }}/command/down"
config_topic_template="gateway/{{ .MAC }}/command/config"
{% endraw %}
```

#### Backends

With LoRa Gateway Bridge v2 you would configure the MQTT backend under the
`[backend...]` section. This has changed and the `[backend...]` section is
now used for selecting and configuring the packet-forwarder backends.

#### Integrations

The MQTT integration configuration has moved to the new `[integration...]`
section. This allows for adding new integrations in the future besides MQTT.

#### Prometheus metrics

The Prometheus metrics have been updated / cleaned up.

## v2.7.1

### Bugfixes

* Fix `rxpk.brd` parsing for some packet-forwarders. ([#104](https://github.com/brocaar/chirpstack-gateway-bridge/issues/104))

## v2.7.0

### Features

* Environment variable based [configuration](https://www.chirpstack.io/chirpstack-gateway-bridge/install/config/) has been re-implemented.

### Improvements

* Remove QoS > 0 bottleneck with non-blocking MQTT publish ([#101](https://github.com/brocaar/chirpstack-gateway-bridge/pull/101))

### Bugfixes

* Fix potential deadlock on MQTT re-connect ([#103](https://github.com/brocaar/chirpstack-gateway-bridge/issues/103))
* Fix logrotate issue (init based systems) ([#98](https://github.com/brocaar/chirpstack-gateway-bridge/pull/98))

## v2.6.2

### Bugfixes

* Fix QoS backend setting regression. [#95](https://github.com/brocaar/chirpstack-gateway-bridge/pull/95)

## v2.6.1

### Bugfixes

* Fix corrupted gateway statistics for downlink packets. [#94](https://github.com/brocaar/chirpstack-gateway-bridge/pull/94)

## v2.6.0

### Features

#### Google Cloud Platform integration

LoRa Gateway Bridge is able to integrate with [Cloud IoT Core](https://cloud.google.com/iot-core/).
This makes it easier to scale up, but also makes it possible to manage gateway
authentication using the web-based Cloud IoT console (or APIs).

#### Gateway IP address

The gateway IP address is exposed through the gateway stats.

## v2.5.1

### Bugfixes

Fix configuration parse error (some keys were not loaded because of configuration alias).

## v2.5.0

### Upgrade notes

When using the `v2_json` marshaler (default), this version is fully compatible
with previous versions of LoRa Server. When changing the marshaler to `protobuf`
or `json`, you first need to upgrade to LoRa Server v2.1.0. When
LoRa Server v2.1.0 is installed, it is recommended to use either the `protobuf`
or `json` marshaler as it provides better compatibility (e.g. with the iBTS gateway).

### Features

#### Protocol Buffer data serialization

To save on bandwith between the gateway and the MQTT broker (e.g. when the
gateway uses a cellular connection), this update makes it possible to
configure the `marshaler` for encoding / decoding the data (in
`chirpstack-gateway-bridge.toml`) using [Protocol Buffers](https://developers.google.com/protocol-buffers/).

This will become the default option in LoRa Gateway Bridge v3.

#### New JSON format

The new JSON structure re-uses the messages defined for
[Protocol Buffers](https://developers.google.com/protocol-buffers/docs/proto3#json)
based serialization. For backwards compatibility, the default `marshaler` for
LoRa Gateway Bridge v2.x will be `v2_json`. 

#### Kerlink iBTS fine-timestamp

When using the `protobuf` or `json` `marshaler`, LoRa Gateway Bridge will
expose the fine-timestamp fields when using Kerlink iBTS gateways.

#### Prometheus metrics

For monitoring the health and throughput of each LoRa Gateway Bridge instance,
LoRa Gateway Bridge is able to expose various metrics that can be scraped by
[Prometheus](https://prometheus.io/).

### Changes

#### Configuration structure of MQTT backend

The structure of the `[backend.mqtt]` configuration section has been updated.
These changes are fully backwards compatible.

### Improvements

All vendored dependencies have been updated.

## 2.4.1

**Bugfixes:**

* Only log in case of UDP read / write errors (instead of breaking out of loop).

## 2.4.0

**Features:**

* LoRa Gateway Bridge can now manage the packet-forwarder configuration (channels).
  See [Send / receive data](https://www.chirpstack.io/chirpstack-gateway-bridge/use/data/) for more information.

## 2.3.2

**Improvements:**

* Expose the following MQTT options for the MQTT gateway backend:
  * QoS (quality of service)
  * Client ID
  * Clean session on connect

**Bugfixes:**

* Use topic from configuration file on re-connect (this was still hardcoded).

## 2.3.1

**Improvements:**

* MQTT topics are now configurable through the configuration file.
  See [Configuration](https://docs.chirpstack.io/chirpstack-gateway-bridge/install/config/).

## 2.3.0

**Features:**

* LoRa Gateway Bridge uses a new configuration file format.
  See [configuration](https://docs.chirpstack.io/chirpstack-gateway-bridge/install/config/) for more information.

* Support MQTT client certificate authentication ([#74](https://github.com/brocaar/chirpstack-gateway-bridge/pull/74)).

**Upgrade notes:**

When upgrading using the `.deb` package / using `apt` or `apt-get`, your
configuration will be automatically migrated for you. In any other case,
please see [configuration](https://docs.chirpstack.io/chirpstack-gateway-bridge/install/config/).

## 2.2.0

**Features:**

* LoRa Gateway Bridge now publishes TX acknowledgement messages over MQTT.
  See [MQTT topics](https://docs.chirpstack.io/chirpstack-gateway-bridge/use/data/).

* TX (GPS) `timeSinceGPSEpoch` field is now exposed to transmit at given
  time since GPS epoch (1980-01-06, only possible when the gateway
  has a GPS time source).

* RX (GPS) `timeSinceGPSEpoch` field is now exposed, containing the time
  since GPS epoch (1980-01-06, only available when the gateway has a GPS
  time source).

**Bugfixes:**

* Without GPS time source, the gateway would use `0001-01-01T00:00:00Z`
  as RX `time`. The `time` field is now omitted when unavailable.


## 2.1.6

**Features:**

* Add Kerlink iBTS compatibility (multi antenna and multi board).

## 2.1.5

**Improvements:**

* `--mqtt-ca-cert` / `MQTT_CA_CERT` configuration flag was added to
  specify an optional CA certificate
  (thanks [@minggi](https://github.com/minggi)).

**Bugfixes:**

* MQTT client library update which fixes an issue where during a failed
  re-connect the protocol version would be downgraded
  ([paho.mqtt.golang#116](https://github.com/eclipse/paho.mqtt.golang/issues/116)).

## 2.1.4

* Retry connecting to MQTT broker on startup (thanks @jcampanell-cablelabs).
* Make latitude, longitude and altitude optional as this data is not always
  provided by the packet_forwarder and would else incorrectly be set to `0`.

## 2.1.3

* Provide `TXPacketsReceived` and `TXPacketsEmitted` in stats.
* Implement `--skip-crc-check` / `SKIP_CRC_CHECK` config flag to disable CRC
  checks by LoRa Gateway Bridge.

## 2.1.2

* Add optional `iPol` field to `txInfo` struct in JSON to override the default
  behaviour (which is `iPol=true` when using LoRa modulation)

## 2.1.1

* Do not unmarshal and marshal PHYPayload on receiving / sending
* Set FDev field when using FSK modulation ([#16](https://github.com/brocaar/chirpstack-gateway-bridge/issues/16))
* Omit empty time field ([#15](https://github.com/brocaar/chirpstack-gateway-bridge/issues/16))

## 2.1.0

* Support protocol v1 & v2 simultaneously.

## 2.0.2

* Rename from `lora-semtech-bridge` to `chirpstack-gateway-bridge`

## 2.0.1

* Update `lorawan` vendor to fix a mac command related marshaling issue.

## 2.0.0

* Update to Semtech UDP protocol v2. This is the protocol version used
  since `packet_forwarder` 3.0.0 (which implements just-in-time scheduling).

## 1.1.3

* Minor log related improvements.

## 1.1.2

* Provide binaries for multiple platforms.

## 1.1.1

* Rename `DataRate` to `BitRate` (FSK modulation).

## 1.1.0

* Change from [GOB](https://golang.org/pkg/encoding/gob/) to JSON.

## 1.0.1

* Update MQTT vendor to fix various connection issues.

## 1.0.0

Initial release.
