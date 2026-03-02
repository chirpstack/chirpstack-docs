# Changelog

## v4.17.0

### Features

#### CLI reset-password command

This release adds a new `reset-password` command to the ChirpStack CLI. This
command enables resetting user passwords without API access, which is useful for:

- Initial setup of fresh installations
- Automated deployment scripts
- Recovery from forgotten passwords

**Usage:**

```bash
# Interactive password reset
chirpstack --config /etc/chirpstack reset-password -e admin@example.com

# Password from file (recommended for scripts)
chirpstack --config /etc/chirpstack reset-password -e admin@example.com -p /tmp/pw.txt

# Password from stdin (e.g. piped from a secrets manager)
my-secrets-tool get admin-password | chirpstack --config /etc/chirpstack reset-password -e admin@example.com --stdin
```

See [CLI commands](./use/cli-commands.md) for more information.

## v4.16.2

### Bugfixes

#### SQLite migration fix

This release fixes a critical issue in the SQLite migration, that causes
deleting devices from the device table.
([#825](https://github.com/chirpstack/chirpstack/issues/825)).

Context: SQLite has limited support for `ALTER TABLE` queries, therefore it is
often needed to re-create the table with the updated schema, copy over the data
and remove the old table. In this specific migration, the `device_profile` table
is re-created to drop the `NOT NULL` constraint on the `tenant_id` column to
make it possible to store both "global" and tenant-owned device profiles. To
avoid that this triggers the deletion of data with a foreign-key constraint to
this table, the re-create, copy, delete and rename process was between:

```sql
pragma foreign_keys = 0;
-- create, copy, delete, rename queries
pragma foreign_keys = 1;
```

This works fine when executed outside the context of an SQLite transaction
(which was tested), but fails when executed within a transaction (when executed
through `diesel`), as per
[SQLite documentation](https://www.sqlite.org/pragma.html#pragma_foreign_keys):
_This pragma is a no-op within a transaction_.

This is fixed by disabling the `foreign_keys` `PRAGMA` before `diesel` starts
the transaction for executing migrations.

## v4.16.1

**Note:** Do not use this version if you are using the SQLite build!

### Bugfixes

- Fix missing RP002-1.0.5 in reg. params. revision dropdown.
- Fix incorrect `google.api.http` option in `DeviceProfileService` Protobuf
  file.

## v4.16.0

### Features

#### RP002-1.0.5 support

This adds support for RP002-1.0.5 support, which includes support for LoRa SF5
and SF6. This also refactors how the min / max data-rate range works internally.
This is required as with RP002-1.0.5 the semantics of min / max data-rates has
changed and is no longer guaranteed to be a range. Please refer to the
RP002-1.0.5 specs for more information.

As devices supporting RP002-1.0.5 are not required to implement all data-rates,
ChirpStack assumes the device only supports the mandatory data-rates, unless
otherwise configured in the device-profile.

#### DNS-based join-server resolving

This release adds DNS-based join-server resolving, based on the JoinEUI from the
join-request.

#### Device-profile refactor

This change refactors the device-profile templates into the device profiles.
Device profiles can now be global (and used by multiple tenants) and / or
managed per-tenant (when the device profiles can only be used by the given
tenant). This way, it is no longer needed to copy a device-profile template into
a device profile. It also means that changes to a global device profile
immediately propagate to the devices using the change device profile.

This change replaces the `import-legacy-lorawan-devices-repository` device
profile import command by `import-device-profiles`, which import device profiles
from
[https://github.com/chirpstack/chirpstack-device-profiles](https://github.com/chirpstack/chirpstack-device-profiles).
Vendors and ChirpStack-users are encouraged to add their devices / devices they
are using to this repository.

**Note:** Migrating existing device-profile templates to global device profiles
can be performed using the `migrate-device-profile-templates` sub-command, e.g.:

```
chirpstack -c /path/to/config migrate-device-profile-templates
```

### Improvements

- Add configurable PostgreSQL connection recycling method.
  ([#794](https://github.com/chirpstack/chirpstack/pull/794))
- Remove unused "integration" event code-traces.
- Replace `rand` by `getrandom` crate to reduce number of dependencies.
- Add option to filter out non-LoRaWAN frames in `lrwn_filters` crate.
- Move shared meta-data to workspace `Cargo.toml`.
- Update internal dependencies.

### Bugfixes

- Fix typo in `region_eu433.toml` description.
  ([#775](https://github.com/chirpstack/chirpstack/pull/775))
- Fix saving measurements + auto-detect options in device profiles.
  ([#769](https://github.com/chirpstack/chirpstack/issues/769))

## v4.15.0

### Features

#### Blynk integration

This release adds a new integration with the [Blynk](https://www.blynk.io/) IoT
platform.

#### InfluxDB v3

This release updates the InfluxDB integration to also support InfluxDB v3.

#### Move FCntUp out of device-session

This moves the `f_cnt_up` field out of the device-session blob and stores it as
a separate column in the `device` table. By doing this, it is no longer needed
to wrap the uplink validation into a transaction and lock the device records
matching the uplink DevAddr, as we can update the record with a
`where f_cnt_up = OLD_F_CNT_UP` and check that the updated record-count != 0.
Removing the transaction and locking the device-records matching the uplink
DevAddr avoids a potential deadlock situation that might occur when multiple
devices are sharing the same DevAddr.

The `f_cnt_up` will be automatically migrated from the device-session into the
`f_cnt_up` device column at runtime.

### Improvements

- Add `flush_queue` option to `Enqueue` API method.
- Add configuration option to flatten json log output.
  ([#759](https://github.com/chirpstack/chirpstack/pull/759))
- Update Docker Compose dependency versions (Mosquitto, PostgreSQL, Redis,
  RabbitMQ and Kafka versions).
- Update internal dependencies.
- Cleanup dead code + fix clippy warnings.
- Move dependency versions into `[workspace.dependencies]` of `Cargo.toml` (such
  that versions can be managed at a single place).

### Bugfixes

- UI: Fix setting gateway last-seen to undefined.
  ([#75](https://github.com/chirpstack/chirpstack/issues/750))
- Fix setting TX Info for OTAA join-accept such that RX2 can be prioritized over
  TX1.

## v4.14.1

### Improvements

- Implement using `fPort` returned by JS codec.

### Bugfixes

- Fix dependency mismatch in AMQP integration.

## v4.14.0

### LoRa Cloud integration removal

This release removes the [LoRa Cloud](https://www.semtech.com/) integration from
ChirpStack as it has been discontinued by Semtech. This will also remove the
LoRa Cloud integration settings (if any) from the database.

### Improvements

- Update internal dependencies.
- Update device-profile importer.
- Ensure `jq` is available in the development shell.
  ([#705](https://github.com/chirpstack/chirpstack/pull/705))
- Add support for 2.4GHz frequencies in FUOTA multicast-setup.
  ([#700](https://github.com/chirpstack/chirpstack/issues/700))
- Improve reported errors by ADR plugin system.
  ([#712](https://github.com/chirpstack/chirpstack/pull/712))
- Add `expires_at` field to MQTT downlink command.
  ([#608](https://github.com/chirpstack/chirpstack/issues/608))

### Bugfixes

- Fix missing FUOTA API methods in gRPC client packages.
  ([#709](https://github.com/chirpstack/chirpstack/issues/709))
- Add distance type to Cayenne LPP codec.
  ([#704](https://github.com/chirpstack/chirpstack/pull/704))
- Fix FUOTA `block_ack_delay` and `descriptor` fields.
  ([#717](https://github.com/chirpstack/chirpstack/issues/717))

## v4.13.0

### Features

#### Refactor Mesh heartbeat

This refactors the Gateway Mesh heartbeat payload into a more generic Mesh Event
payload, which can be used by both "known" and "proprietary" event types. The
first implemented "known" type is the heartbeat. While this refactor is not
backwards compatible and requires
[ChirpStack MQTT Forwarder](https://www.chirpstack.io/docs/chirpstack-mqtt-forwarder/)
v4.4+, the backwards incompatibility does not affect the functioning of the
Gateway Mesh itself.

### Improvements

- Store mac-command stats in device-session (instead of Redis key).
- Update `scheduler_run_after` directly at devicve-session get (avoiding
  race-conditions between Class-A & Class-B & -C).
- Rename `m_type` to `f_type` to align with LoRaWAN 1.0.4 specification naming.
- Make it more clear if FUOTA job is running or if it is schedule.
- Make AMQP integration exchange configurable.
  ([#677](https://github.com/chirpstack/chirpstack/pull/677))
- Replace Redis `GETDEL` by pipelined `GET` and `DEL` commands.
  ([#682](https://github.com/chirpstack/chirpstack/pull/682))
- Sort `NULL` values as smallest value (device and gateway last-seen sorting).
  ([#683](https://github.com/chirpstack/chirpstack/issues/683))

### Bugfixes

- Fix custom root CA certificates not taken into account.
  ([#684](https://github.com/chirpstack/chirpstack/pull/684))
- Fix inconsistency in Class-B periodicity between backend and UI and align
  naming with LoRaWAN 1.0.4 specification naming.
  ([#670](https://github.com/chirpstack/chirpstack/issues/670))

## v4.12.1

### Improvements

- Update internal dependencies.

### Bugfixes

- ui: Fix JSON enqueue error.
  (#[658](https://github.com/chirpstack/chirpstack/issues/658))
- ui: Fix error after deleting tenant.
  ([#635](https://github.com/chirpstack/chirpstack/issues/635))
- ui: Fix potential code-editor rendering issues (by replacing codemirror by
  ace).

## v4.12.0

### Features

#### FUOTA Support

This release brings FUOTA support to ChirpStack, deprecating the need to use the
[ChirpStack FUOTA Server](https://github.com/chirpstack/chirpstack-fuota-server).
Currently it supports both the v1 and v2 versions of TS003, TS004 and TS005
specifications.

**Note:**

This has been tested against the
[LoRa Basics Modem](https://github.com/lora-net/SWL2001) stack. While testing, a
small implementation difference has been found between the LoRa Basics Modem and
the TS005 (v2) specification
([#106](https://github.com/Lora-net/SWL2001/issues/106)). At this moment, the
ChirpStack implementation has been aligned with the LoRa Basics Modem.

### Improvements

- Add sorting to device and gateway tables.
  ([#579](https://github.com/chirpstack/chirpstack/pull/579))
- Add Yandex ID OAuth provider support.
  ([#622](https://github.com/chirpstack/chirpstack/pull/622))
- Generalize auto-conversion of SEC1 EC Keys to PKCS#8.
  ([#618](https://github.com/chirpstack/chirpstack/pull/618))
- Replace static UUIDs by `gen_random_uuid()` in PostgreSQL migration.
  ([#634](https://github.com/chirpstack/chirpstack/issues/634))
- Show origin in case of bind parse error.
  ([#632](https://github.com/chirpstack/chirpstack/issues/632))
- Refactor device-profile fields in database to reduce amount of used columns.
- Add filtering to device table.
- Update internal dependencies.
- Add JavaScript example for default ADR algorithm.
- Add option to configure join-server authorization header.

## v4.11.1

### Improvements

- Return `LinkCheckAns` with `margin=0` for non-LoRa modulations.
- Re-use http clients in integrations to reduce overhead.

### Bugfixes

- Fix missing `ns_time` after multi-region refactor.

## v4.11.0

### Improvements

#### Refactor (multi-region) de-duplication

This release refactors the uplink de-duplication, such that uplinks received by
gateways in different region configurations (but with overlapping channels) are
de-duplicated separately. After de-duplication and after resolving the
device-session of the device, ChirpStack will only handle the de-duplicated
uplink that is within the same region-configuration assigned to the
device-session.

This makes it possible to create region configurations of which channels and
gateway-coverage are overlapping. Before this release such scenario could result
in constantly changing settings (through mac-commands) of the end-device or
errors related to unknown channels.

This does include two integration changes:

- Region information is no longer exposed through the `rxInfo` elements.
- A new `region_config_id` field has been added to join and uplink events.

## v4.10.2

### Improvements

- Store DevNonces per JoinUI.
- Update base Docker base to Alpine 3.21.0.
- Internal dependency updates.

### Bugfixes

- Delete expired multicast queue-items in case the gateway is offline
  (previously these would be delete once the gateway would come online).

## v4.10.1

### Features

#### SQLite support

This add support for using SQLite instead of PostgreSQL as database. Please note
that the database backend must be selected at compile time (default PostgreSQL).
SQLite is ideal for embedded ChirpStack deployments (like the ChirpStack Gateway
OS, which will migrate to SQLite in the next version). Pre-compiled binaries are
provided for both PostgreSQL and SQLite in the
[Downloads](https://www.chirpstack.io/docs/chirpstack/downloads.html).
([#418](https://github.com/chirpstack/chirpstack/pull/418),
[#540](https://github.com/chirpstack/chirpstack/pull/540))

#### Queue-item expiration

This makes it possible to set an expires-at timestamp for multicast and unicast
queue-items, in case they must be automatically removed from the queue after a
given timestamp.

### Improvements

- ui: Update page-title according to page.
  ([#539](https://github.com/chirpstack/chirpstack/pull/539))
- api: Update gRPC dependencies in Java API.
  ([#494](https://github.com/chirpstack/chirpstack/pull/494))
- api: Re-export `pbjson_types` and `tonic` in Rust API.
  ([#503](https://github.com/chirpstack/chirpstack/pull/504))
- api: Add support for C.
  ([#549](https://github.com/chirpstack/chirpstack/pull/549))

### Bugfixes

- ui: Fix tooltip data formatting.
  ([#503](https://github.com/chirpstack/chirpstack/issues/503))
- Remove `OFF` from log levels (this is an invalid level).
- Fix escaping when generating config template.
- Fix `validate_mic` clause (Backend Interfaces).

## v4.10.0

Unreleased because of issue in Debian packaging.

## v4.9.0

### Features

#### Gateway Mesh

This is a new feature that makes it possible to extend network coverage by
deploying 'Relay Gateways' that operate without internet access. While the
Gatway Mesh only requires changes to the LoRa Gateways, this implements an
overview of detected Relay Gateways and their last heartbeat. Please see
[ChirpStack Gateway Mesh](../chirpstack-gateway-mesh/index.md) for for
information. Support for this feature will be implemented in the next
[ChirpStack Gateway OS](../chirpstack-gateway-os/index.md) version.

#### Other features

- Add support for generating PHP gRPC code.
  ([#431](https://github.com/chirpstack/chirpstack/pull/431))

### Improvements

- Update internal dependencies.
- Add `passive_roaming_validate_mic` configuration option to enable / disable
  MIC validation for passive-roaming.
- Update sslmode options in PostgreSQL config template.
  ([#433](https://github.com/chirpstack/chirpstack/pull/433))
- Fallback to empty DevAddr, EUI, NetID or AES-key in case of empty string.
  ([#453](https://github.com/chirpstack/chirpstack/issues/453))
- Make OIDC and OAuth2 scopes configurable.
  ([#443](https://github.com/chirpstack/chirpstack/pull/445))
- Update GitHub workflow action versions.
  ([#461](https://github.com/chirpstack/chirpstack/pull/461))
- Document `user_info` configuration in `region_xxxx.toml` configuration files.
  ([#462](https://github.com/chirpstack/chirpstack/issues/462))
- Replace `warp` dependency by `axum` for API request handling.
- UI: Make it possible to configure the tileserver.
  ([#451](https://github.com/chirpstack/chirpstack/pull/451))
- UI: Replace `moment` with `date-fns`.
  ([#460](https://github.com/chirpstack/chirpstack/pull/460))
- UI: Migrate from `create-react-app` to `vite`.
  ([#459](https://github.com/chirpstack/chirpstack/pull/459))
- UI: Update codec template with JSDoc.
  ([#473](https://github.com/chirpstack/chirpstack/pull/473))
- API: Improve device activation in documentation.
  ([#454](https://github.com/chirpstack/chirpstack/pull/454))

### Bugfixes

- API: Remove `frame_log` from `__init__.py` in Python SDK.
  ([#422](https://github.com/chirpstack/chirpstack/issues/422))
- Fix sending `HomeNSAns` message-type to `HomeNSReq` request.
- Fix empty string in Regional Paramaters revision column.
  ([#432](https://github.com/chirpstack/chirpstack/pull/432))
- Fix exposing full JS codec errors.
  ([#440](https://github.com/chirpstack/chirpstack/issues/440))
- Fix setting gateway altitude in stats handler.
  ([#458](https://github.com/chirpstack/chirpstack/pull/458))
- Add signal handler to handle `SIGTERM` correctly.
  ([#480](https://github.com/chirpstack/chirpstack/issues/480))

## v4.8.1

### Features

#### Duty-cycle metrics

This release implements duty-cycle metrics for each gateway. Please note that
this feature is (currently) only available for EU868 region and if the
ChirpStack Concentratord is used (v4.4.0+).

#### Increase RX1 Delay in device-profile

This makes it possible to increase the RX1 delay for one or multiple devices
through the device-profile. This can be useful when the latency between
ChirpStack and the integration is too high to be able to enqueue a downlink for
the same RX1 / RX2 receive-windows. Note that it does not allow setting a RX1
delay lower than the `rx1_delay` value configured in the ChirpStack
configuration files.

### Improvements

- Update internal dependencies.
- Document missing `recvTime` codec option.
  ([#385](https://github.com/chirpstack/chirpstack/issues/385))
- Show error in UI if clipboard API is not available.
  ([231](https://github.com/chirpstack/chirpstack/issues/231))
- Show ChirpStack version in web-interface (if user is admin).
  ([#73](https://github.com/chirpstack/chirpstack/issues/73))
- Expose more verbose JavaScript codec error output.
  ([#391](https://github.com/chirpstack/chirpstack/issues/391))
- Align multicast Class-B ping-slot configuration.
  ([#255](https://github.com/chirpstack/chirpstack/issues/255))
- Remove generated API code from repository to reduce amount of changes in case
  of API changes.
- Add support for responding to `HomeNSReq` API requests (Backend Interfaces).

### Bugfixes

- Do not update gateway location when updated lat / lon / alt are all set to 0.
- Fix typo in auth error string.
  ([#367](https://github.com/chirpstack/chirpstack/pull/367))
- Add auto-coversion from SEC1 EC keys to PKCS#8.
  ([#386](https://github.com/chirpstack/chirpstack/issues/386))
- Fix drawer header z-index issue in web-interface.
  ([#393](https://github.com/chirpstack/chirpstack/pull/393))
- Auto-detect if MQTT (v5) broker supports shared-subscriptions.
  ([#413](https://github.com/chirpstack/chirpstack/pull/413))
- Fix stats interval calculations in case of DST changes (in which case some
  timestamps don't exist).
  ([#415](https://github.com/chirpstack/chirpstack/issues/415))
- Fix loading auto-complete options in web-interface.
  ([#334](https://github.com/chirpstack/chirpstack/issues/334))
- Do not schedule Class-B / Class-C downlinks for disabled devices.

## v4.8.0

Release was skipped because an issue was found with the generated gRPC-web and
JavaScript API packages.

## v4.7.0

### Notes before you upgrade

#### Device-session migration

This release moves the device-session storage from Redis to PostgreSQL. After
upgrading, you must execute the following command (adapted to your environment):

```bash
chirpstack -c /etc/chirpstack migrate-device-sessions-to-postgres
```

This command will iterate over the devices in the PostgreSQL database of which
the device-session column is empty and will migrate the device-session from
Redis if it exists. This will not overwrite existing device-sessions in
PostgreSQL thus it is safe to re-execute this command in case needed.

#### OpenID Connect / OAuth2

A new authentication backend has been added for OAuth2 based providers (see
below). If you are using the OpenID Connect authentication backend, you must
update your configuration from:

```toml
[user_authentication.openid_connect]
enabled=true
```

To:

```toml
[user_authentication]
enabled="openid_connect"
```

#### PostgreSQL CA certificate

If the PostgreSQL server uses TLS, please read the note below with regards to
the `ca_cert` configuration option.

### Features

#### MQTT shared subscription

This by defaults connects to the MQTT broker using a shared subscription name
`chirpstack`. Using a shared subscription, the MQTT broker will send a received
uplink only to one subscriber. In case ChirpStack is deployed as a cluster, this
removes the overhead caused by all instances receiving the same uplink (which
would be correctly handled by the deduplication logic).

In case you have multiple ChirpStack environments (e.g. production, staging and
testing) connected to the same MQTT broker, then make sure that each environment
has a correct `share_name` configured in the `region_XXXXX.toml` configuration.

**Note:** Shared subscriptions is a MQTTv5 feature and therefore a MQTTv5
capable MQTT broker is required (Mosquitto supports MQTTv5 since v1.6).

#### OAuth2 / Clerk integration

This adds support for integrating with the [Clerk](https://clerk.com/)
authentication backend (OAuth2 interface).

### Improvements

#### Store device-sessions in PostgreSQL

This moves the device-session storage from Redis to PostgreSQL. In case of a
high DevAddr re-usage (where multiple DevEUIs share the same DevAddr), the old
architecture had a significant overhead, because it would perform a significant
amount of Redis queries to retrieve all the potential device-sessions. It would
retrieve the DevAddr -> DevEUIs mapping, and then retrieve the device-session
for each DevEUI. Because device-session data can be sharded (Redis Cluster), a
separate query per device-session was required.

With this improvement, a device-session column has been added to the device
table in the PostgreSQL database, which also contains a DevAddr column. The big
advantage is that all device-sessions for a given DevAddr can be retrieved using
a single query.

This improvements also means that:

- Device-sessions will no longer expire after device inactivity
- Device-sessions can be restored from a PostgreSQL backup

#### Replace OpenSSL with Rustls

This is an internal improvement and removes the OpenSSL dependency in favor of
Rustls, which is a pure Rust TLS implementation. This makes the build process
easier as it is no longer needed to build a (static) OpenSSL version to link
against.

#### Use async PostgreSQL

This is an internal improvement, and migrates away from `pq-sys` in favor of
`tokio-postgres`, which is a pure Rust client implementation which works with
`async`. This removes all `task::spawn_blocking(...)` blocks around SQL queries.
As well, we no longer need to static link against libpq (C library, with
dependency on OpenSSL).

**Important note:** this also adds a new `ca_cert` option to the `[postgresql]`
configuration section where you can configure the CA certificate which must be
used for validating the PostgreSQL server certificate (if not already provided
by the host system).

#### Use async Redis

This is an internal improvement and updates all Redis queries to use `async` /
`await` and removes all `task::spawn_blocking(...)` blocks around Redis queries.

#### Paho MQTT > rumqttc client

This is an internal improvement and replaces `paho-mqtt` with `rumqttc`. The
latter crate is a pure Rust client which uses `rustls` for TLS instead of
OpenSSL.

#### Other

- Update dependencies.
- Return Redis connection to the pool immediately after query completion.
- Return PostgreSQL client to the pool immediately after query completion.

## v4.6.0

**Important note before you upgrade:**

- If you are migrating from ChirpStack v3 to ChirpStack v4 and are still using
  the ChirpStack Gateway Bridge v3.14.x, then you must add `v4_migrate=true` to
  your configuration as described by the
  [v3 to v4 migration documentation](https://www.chirpstack.io/docs/v3-v4-migration.html#enable-v3-compatibility-to-chirpstack).

### Features

#### End-to-end encryption

This feature makes it possible to implement end-to-end encryption between the
end-device and end-application. On OTAA join, the join-server will provide
Chirpstack with the encrypted `AppSKey`, which will be forwarded on every uplink
to the end-application (integration events). The end-application then first
decrypts the `AppSKey` with the KEK key shared between the JS and
end-application, and then uses the decrypted `AppSKey` to decrypt the
application payload.

On enqueue downlink, the end-application encrypts the application payload before
enqueue. As well, it must set the `f_cnt_down` and `is_encrypted` fields such
that ChirpStack knows that the payload is already encrypted and which downlink
frame-counter was used during the encryption of the payload.

**Note:** This feature requires an external join-server.

#### Add `chirpstack_integration` crate

This `chirpstack_integration` crate can be used to build external integrations
using the Redis Streams that are exposed by ChirpStack. An example
implementation is the
[ChirpStack Pulsar Integration](https://github.com/chirpstack/chirpstack-pulsar-integration/).

#### Tenant and application tags

This adds tags (like already can be found on device-profiles and devices) to
tenants and applications. Note that the integration events will contain the
aggregation of application + device-profile + device tags. Integration events
will not contain the tenant tags.

#### Allow JoinEUI prefix configuration

This makes it possible to configure a JoinEUI prefix when configuring a
join-server, to forward a range of JoinEUI to a single join-server without the
need to add multiple join-server configuration blocks. As well, this makes it
possible to configure a 'catch-all' join-server, using a JoinEUI prefix that
would match all JoinEUIs.

#### Refactor streams API + expose Backend Interfaces requests

This moves some of the API:

- `meta/meta.proto` -> `streams/meta.proto`
- `api/frame_log.proto` -> `streams/frames.proto`
- `api/request_log.proto` -> `streams/api_requests.proto`

If you are using these messages in your application, then you might need to
update the import paths when updating the API SDK.

As well, this adds a new Redis Stream exposing the Backend Interfaces requests
and responses (Passive Roaming + Join Server).

#### Add allow roaming option to Device Profile

This makes it possible to select which devices are allowed to use roaming and
which devices not. This option can be configured in the device-profile. On
migration this value will be set to `true` (as all devices could use roaming
before this version).

#### Add `assume_email_verified` option for OIDC

Some OpenID Connect providers do not provide an `email_verified` value. By
setting the `assume_email_verified` to `true`, ChirpStack will assume that the
e-mail address has been verified.
([#302](https://github.com/chirpstack/chirpstack/issues/302))

### Improvements

- Expose `skip_f_cnt` and device variables to ADR plugins.
- Reset uplink ADR history table in case of DR / TxPower / NbTrans change.
- Add `secondary_net_ids` configuration option.
- Do not fail in case of corrupted mac-commands.
- Use region default RX2 frequency if device-session RX2 frequency == 0.
- Make it explicit that TX Power is in EIRP + update region configuration from
  ERP to EIRP.
- Refactor device-lock / `scheduler_run_after` setting.
- Ignore unknown JSON fields when decoding JSON to API structures in Rust.
- Rename `time` to `gw_time` and add `ns_time` to the gateway rx-info struct.
- Speed up API authorization validation queries (SQL).
- Improve log output (better log messages + adding better correlation
  identifiers to each message).
- Add `preamble` and `no_crc` fields to `LoraModulationInfo` (this is not used
  by ChirpStack, but it can be used by applications directly interacting with
  the gateway).
- Omit `null` fields in Backend Interfaces JSON output.
  ([#316](https://github.com/chirpstack/chirpstack/pull/316))
- Reduce dependencies for AWS SNS integration by replacing `aws-sdk-sns` crate
  with `aws-sign-v4` + REST call.
- Make device metric name optional.
  ([#313](https://github.com/chirpstack/chirpstack/issues/313))
- Get all device-data in a single query to improve performance.
- Change `v4_migrate` default to `false` (please
  [v3 to v4 migration guide](https://www.chirpstack.io/docs/v3-v4-migration.html)).

## Bugfixes

- Debian package: Fix `postinst` to only run on install.
  ([#295](https://github.com/chirpstack/chirpstack/issues/295))
- Fix setting initial tags in tenant form (UI).
- Use unbounded MQTT client channels / fix dropping MQTT messages under high
  load.
- Add misspelled `UnkownReceiver` as valid `ResultCode` in Backend Interfaces
  (this is a typo in the specifications).
  ([#317](https://github.com/chirpstack/chirpstack/pull/317))
- Reload device on change event.
  ([#319](https://github.com/chirpstack/chirpstack/issues/319))
- Fix sending empty downlink to Relay in case uplink `ADRACKReq` bit was set.
- Return error in ThingsBoard integration if `ThingsBoardAccessToken` is not
  set. ([#277](https://github.com/chirpstack/chirpstack/issues/277))

## v4.5.1

### Improvements

- Show notification in UI for form validation errors.
  ([#282](https://github.com/chirpstack/chirpstack/issues/282))
- Log OTAA join-requests of unknown devices as warning instead of error.
- Re-export `prost` dependency in `chirpstack_api` crate.
  ([#285](https://github.com/chirpstack/chirpstack/pull/285))

### Bugfixes

- Fix `enabled_uplink_channels` error in `region_au915_2.toml`.
  ([#274](https://github.com/chirpstack/chirpstack/pull/274))
- Fix AS923 max-payload size table.
  ([#283](https://github.com/chirpstack/chirpstack/pull/283))
- Fix doughnut chart resizing in UI.
  ([https://github.com/chirpstack/chirpstack/issues/284])

## v4.5.0

### Features

#### UI improvements

This updates the ChirpStack UI to [Ant Design](https://ant.design/) v5 and
[React](https://react.dev/) v18.2.0. This also refactors the UI components code
from classes to functions.

While this refactor does not add additional features to the UI, it does
eliminate some unnecessary reloading of data due to better route handling. This
should result in less flickering when navigating between some of the UI pages.

### Improvements

- Make it possible to use `DeviceModeInd` for LoRaWAN 1.0.x devices (this is not
  according the LoRaWAN 1.0.x specs).
- Implement removing device from Relay filter list.
- Add `rediss://` connection string example to config template (for TLS).
  ([#219](https://github.com/chirpstack/chirpstack/issues/219))
- Update internal dependencies.

### Bugfixes

- Fix showing initial InfluxDB integration form.
  ([#254](https://github.com/chirpstack/chirpstack/issues/254))
- Fix ADR plugin variable mapping typo (`maxRd` > `maxDr`).
  ([#256](https://github.com/chirpstack/chirpstack/pull/256))
- Fix setting empty username / password in MQTT.
  ([#257](https://github.com/chirpstack/chirpstack/issues/257))
- Fix RPM install error caused by auto dependencies.
  ([#267](https://github.com/chirpstack/chirpstack/issues/267))

## v4.4.3

### Bugfixes

- Wait for first uplink before scheduling Class-C downlink.
- Do not disable device activation fields in UI (making it impossible to read &
  copy session-keys).

## v4.4.2

### Features

- Add `create-api-key` sub-command to create API keys using the CLI.

### Improvements

- Update internal dependencies.

### Bugfixes

- Fix sending multiple `FilterListReq` mac-commands (Relay).

## v4.4.1

### Improvements

- Update internal dependencies.
- Add support for JSON log output.

### Bugfixes

- Fix sending channel-mask twice (CFList + LinkADRReq) for US915-like regions.

## v4.4.0

### Features

#### Relay support (TS011)

This adds support for the Relay specification (TS011). In the Device Profile it
is possible to configure the Relay and Relay capable end-device parameters.
Under the Application view it is possible to assign devices to a Relay (required
for filtering and exchanging the device list with the Relay).

**Note:** Please note that this requires a Relay and Relay capable end-device.

### Improvements

#### Build changes

The build configuration has been updated to generate fully static binaries based
on _musl libc_. This solves the issue where in some cases ChirpStack would not
connect over TLS to a PostgreSQL database.
([#156](https://github.com/chirpstack/chirpstack/issues/156)).

This also changes the Docker base image to `alpine`, reducing the Docker image
size by ~ 50% compared to `debian:buster-slim`. Within the `Dockerfile` we now
`COPY` the already compiled binaries, which also reduces the build time on
release.

If you are compiling ChirpStack from source, please refer to the `README.md` in
the source repository as some commands have changed.

#### IFTTT integration

Configuration has been added to configure the prefix of the event name and to
send arbitrary JSON payloads instead of the 3 value payload.

#### Other improvements

- Add `lrwn_filters` crate for filtering LoRaWAN PHYPayloads.
- Dependencies have been updated.
- Expose enabled device-class in API and integration event output.
  ([#58](https://github.com/chirpstack/chirpstack/issues/58))

### Bugfixes

- Fix `netid_type` method `panic` in case of invalid DevAddr prefix type.
- Fix missing device `search` filter (API).
- Fix incorrect config template key for `bind` ([backend_interface]).
- Fix default `CN470` Class-B ping-slot frequencies.
- Fix using system CA certificates for TLS.
  ([#204](https://github.com/chirpstack/chirpstack/pull/204))
- Fix OTAA join-request MIC check and DevNonce validation order.

## v4.3.2

### Improvements

#### Disable v3 compatibility option

This add a configuration option, to enable / disable compatibility with the
latest ChirpStack Gateway Bridge v3 version. If compatibility is not needed,
then add `v4_migrate=false` to the `[regions.gateway.backend.mqtt]` section (in
`region_...toml`). This will save some bandwidth in the GW <> NS communication.
The current default is `true`, in ChirpStack v4.4+, the default will change to
`false` (and thus it must be explicitly enabled).

#### Other improvements

- Add back web-interface option to download events and frames as JSON.
- Update internal dependencies.

### Bugfixes

- Fix sending to multiple URLs in case one endpoint fails.
- Enable new `tls-rustls` feature of `redis` dependency (fixes Redis TLS
  issues). ([#170](https://github.com/chirpstack/chirpstack/issues/170))

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

- Fix typo in `application.proto`.
  ([#143](https://github.com/chirpstack/chirpstack/pull/143))
- Earlier db initialization + fix unwrap errors.
  ([#147](https://github.com/chirpstack/chirpstack/issues/147))

## v4.3.0

### Features

### C# SDK

This adds support for generating C# API code.
([#100](https://github.com/chirpstack/chirpstack/pull/100))

### Multicast improvements

#### Add gateways to multicast-group

This makes it possible to explicitly define which gateways will be used for a
given multicast-group. In the gateways overview it is possible to select one or
multiple gateways and then add these gateways to a multicast-group.

#### Multicast scheduling option

This moves the multicast scheduling configuration to the multicast-group
configuration from the `chirpstack.toml` configuration file. Scheduling options
include scheduling based on GPS time or delay based.

#### Split private gateways in private uplink and downlink

This makes it possible to define per tenant if and how gateways can be used by
other tenants. For example in case downlink is set to private, other tenants
will benefit from uplinks received by these gateways, but they will not be able
to send downlinks (ChirpStack will filter these gateways out when selecting the
gateway for scheduling the downlink).

### Improvements

- Implement `get_downlink_data_delay` setting.
- Update internal dependencies.
- Add missing `fPort` validation to avoid enqueue on `fPort=0`.
- Do not overwrite `RxInfo` `location` field with gateway location if it is
  already set.
- Do not log stats handling `NotFound` errors if `allow_unknown_gateways` is
  configured.
- Decode FRMPayload mac-commands in device LoRaWAN frames log.
- Show `FCnt` in device event log.

### Bugfixes

- Fix `/api/multcast-groups/...` > `/api/multicast-groups/...` typo in enqueue
  API URL.
- Fix `gateway_id` is missing errors (in case the uplink was also received by an
  unknown gateway).
- Fix disabling mac-commands.
- Fix region configuration defaults + use region `id` as fallback for
  `description` if the latter is missing.
  ([#120](https://github.com/chirpstack/chirpstack/issues/120))
- Fix API authorization for listing ADR algorithms.
  ([#112](https://github.com/chirpstack/chirpstack/issues/112))
- Fix US915 downlink channel `min_dr` configuration.
  ([#115](https://github.com/chirpstack/chirpstack/pull/115))
- Fix incorrect rendering of `integration.mqtt.client.ca_key` and `.ca_cert` in
  config template. ([#124](https://github.com/chirpstack/chirpstack/pull/124))
- Remove `tls_enabled` config option as it is not actually implemented.
  ([#128](https://github.com/chirpstack/chirpstack/issues/128))

## v4.2.0

### Features

#### Fix devices to specific region configuration (optional)

This adds a _Region configuration_ option to the device-profile, which lists the
region configurations for the selected _Region_. If a region configuration is
selected, then the device will only be able to work under the selected
configuration. If no region configuration is selected, then the device will be
able to operate under all available region configurations for the selected
region.

#### Java API SDKs

This adds support for Java and Kotlin API SDK code generation.
([#64](https://github.com/chirpstack/chirpstack/pull/64))

### Improvements

- Add `description` configuration option per region configuration.
- Change `name` to `id` within the region configuration (`name` will be used as
  fallback option).
- Make gateway state in UI consistent and make expected stats interval
  configurable. ([#76](https://github.com/chirpstack/chirpstack/issues/76))
- Add Python type information to Python API SDK code.
  ([#68](https://github.com/chirpstack/chirpstack/pull/68))
- Add back `crc_status` field to `UplinkRxInfo` message (the status will be
  reported as no CRC until the ChirpStack Gateway Bridge, ChirpStack
  Concentratord and / or ChirpStack MQTT Forwarder have been updated).
- Add back Class-B ping-slot parameters to the device-profile.
- Update Class-B ping-slot data-rate configuration in examples.
- Remove separate gateway topic config and move it into single `topic_prefix`
  configuration.
- Reset internally stored channels to default on `ADRACKReq` uplink to avoid
  out-of-sync channel configuration on device.
- Update internal dependencies.

### Bugfixes

- Fix hiding _Delete device_ option if the user has no permissions to perform
  this action. ([#71](https://github.com/chirpstack/chirpstack/issues/71))
- Fix not recording device metrics if auto-detect of measurements is disabled.
  ([#94](https://github.com/chirpstack/chirpstack/issues/94))

## v4.1.3

### Bugfixes

- Fix Redis `key_prefix` configuration. While this value could be configured, it
  was not applied to the generated keys.
- Fix header `z-index` issue in UI. This was causing the dropdowns to render
  partly behind the header.

## v4.1.2

### Improvement / bugfix

- Do not wait for integrations to finish before sending downlink.

## v4.1.1

### Improvements

- Update JS API dependencies to latest versions.
- Replace relative paths in Rust API build to absolute.
  ([#69](https://github.com/chirpstack/chirpstack/pull/69))

### Bugfixes

- Fix setting the full frame-counter to the uplink frame after resolving the
  device-session (this can affect payload decryption).

## v4.1.0

### Features

#### API request logging

This feature logs API requests to Redis Streams. This enables external services
to monitor for example device create, update and deletes by reading from the
Redis Streams. A code-example can be found
[here](https://github.com/chirpstack/chirpstack/blob/master/examples/request_log/go/main.go).

#### Uplink logging for unknown devices

While the feature to log frames was already present, it was not possible to only
read uplink frames of devices that are unknown. This extends the frame logging
feature to also log uplinks for unknown devices, in which case DevEUI
`0000000000000000` is used. A code-example for reading the frame log can be
found
[here](https://github.com/chirpstack/chirpstack/blob/master/examples/frame_log/go/main.go).

#### Event logging

A code-example to read event logs from the Redis Streams was added. It can be
found
[here](https://github.com/chirpstack/chirpstack/blob/master/examples/event_log/go/main.go).

### Improvements

- Make `metadata` fields in gateway messages consistent.
- Emit all fields for JSON integration messages, even if they are their default
  values. ([#63](https://github.com/chirpstack/chirpstack/pull/63))

### Bugfixes

- Fix Redis pipelined commands in case Redis Cluster is configured.
- Fix UI notifications z-index.

### Other changes

- The `import-ttn-lorawan-devices` sub-command has been renamed to
  `import-legacy-lorawan-devices-repository`.

## v4.0.5

### Improvements

- Add `keep_alive_interval` for MQTT configuration.

### Bugfixes

- Fix incorrect splitting of multiple URLs in HTTP integration.
  ([#62](https://github.com/chirpstack/chirpstack/issues/62))
- Fix missing ThingsBoard location and status telemetry.
- Send ThingsBoard telemetry (fPort, fCnt, ...) even in case there is no decoded
  payload.
- Fix LeafletJS controls floating over header (UI).

## v4.0.4

### Bugfixes

- Fix coding-rate for LoRaWAN 2.4GHz.
  ([#51](https://github.com/chirpstack/chirpstack/issues/51))
- Fix not removing of queue-item after it was sent, if payload did not fit RX2.

## v4.0.3

### Improvements

- Add code example for reading frame-logs from Redis Streams.
- Add missing metadata logging and add code example for reading metadata from
  Redis Streams.
- Add `dev_addr_prefix` configuration.
  ([#49](https://github.com/chirpstack/chirpstack/issues/49))
- Make it possible to enable / disable auto-detection of metrics in
  device-profile. ([#42](https://github.com/chirpstack/chirpstack/issues/42))
- Add Redis config examples for username and password configuration.
  ([#54](https://github.com/chirpstack/chirpstack/issues/54))
- Add metadata tab to gateway configuration in UI.

### Bugfixes

- Use `trust_store` instead of `ca_path` for MQTT integration.
  ([#47](https://github.com/chirpstack/chirpstack/pull/47))
- Fix _Cannot serialize NaN as google.protobuf.Value.number_value_ error.
- Fix logout URL in case OIDC is enabled.
- Fix `java_outer_classname` in `tenant.proto`.
  ([#55](https://github.com/chirpstack/chirpstack/pull/55))
- Fix closing and detecting if eventlog channel is closed.
- Fix metrics interval calculation on daily aggregate in case of DST to non-DST
  timezone change.

## v4.0.2

### Features

- A new menu option Regions was added to the web-interface, exposing per region
  information.

### Improvements

- Internal dependencies were updated.

### Bugfixes

- Fix Wi-Fi geolocation issue in LoRa Cloud integration.
- Fix missing per data-rate stats in gateway dashboard.
- Fix terminating stream loop (and releasing of Redis connection) on client
  disconnect. ([#40](https://github.com/chirpstack/chirpstack/issues/40))
- Fix rendering `client_cert_lifetime` rendering in configuration template (mqtt
  integration).

## v4.0.1

Not released.

## v4.0.0

After many months of development and testing, we are really excited to share
ChirpStack v4.

The aim of ChirpStack v4 is to make it significantly easier to setup and use
ChirpStack, compared to the previous version. One of the major changes that you
will notice is that the ChirpStack Network Server and ChirpStack Application
Server have been merged into a single component. Over the years we have seen
many issues reported on the forum and GitHub, related to setting up and
connecting both services. ChirpStack v4 also provides multi-region support
out-of-the-box, including region configuration. No longer it is needed to define
your own configuration file or setup multiple ChirpStack Network Server
instances to serve multiple regions simultaneously.

A big thank you to the ChirpStack community for supporting and contributing to
the ChirpStack project! Please find below a breakdown of all the new features
and changes that v4 brings.

### Main features and changes

#### Multi-region support

ChirpStack v4 adds multi-region support, removing the need to setup multiple
ChirpStack Network Server instances. Configuration files are included for the
common regions (as defined by the LoRa Alliance), which should help getting
started with ChirpStack.

Each enabled region has its own gateway backend, making it possible to use one
or multiple MQTT brokers for the different gateway pools. In case a single MQTT
broker is used, each backend must be configured with its own MQTT topic prefix
(e.g. `eu868/gateway/...`).

ChirpStack v4 also supports multiple configurations of the same region, e.g. to
configure a US915 for 8 channels and to configure a US915 band for 16 channels.

#### TTN device repository support

ChirpStack v4 adds support for importing the TTN
[LoRaWAN Devices](https://github.com/TheThingsNetwork/lorawan-devices)
repository as device-profile templates, including codec functions if these are
defined in this repository.

#### Device metrics dashboard

In the device-profile template and / or device-profile, it is possible to define
the measurements that are exposed by the device in the decoded payload. Once
defined, ChirpStack will automatically aggregate and store this data. These
metrics can be viewed in the web-interface on the device dashboard.

#### Configuration

##### Directory

Instead of using a single configuration file (e.g.
`chirpstack-network-server.toml`), ChirpStack makes use of a configuration
directory such that the configuration can be split in multiple files. By default
you will find a single `chirpstack.toml` configuration file, and many
`region_...toml` configuration files, split by region.

##### Environment variables

ChirpStack v4 removes the TOML hierarchy to environment variable mapping.
Instead it allows you to define the variables like `$MY_CONFIGURATION`, which
will get automatically substituted when an environment variable is found with
the name `MY_CONFIGURATION`.

#### API

The REST interface that was present in ChirpStack Application Server v3 has been
removed, in favor of the gRPC API interface (please see the `api/` folder of the
repository for the API definitions). However, a gRPC to REST interface bridge
component will be provided as a separate service. Please note that in v3, this
bridge component was embedded and REST interface calls were internally
translated to gRPC calls. Therefore, gRPC was always recommended interface to
use.

The ChirpStack v4 gRPC API does support server reflection, making it possible to
use for example [gRPC UI](https://github.com/fullstorydev/grpcui) or
[BloomRPC](https://github.com/bloomrpc/bloomrpc) as development interface.

#### ChirpStack Gateway Bridge v3 compatibility

ChirpStack v4 is fully compatibility with the latest version of ChirpStack
Gateway Bridge v3. This should help migrating from v3 to v4. Please note that
the ChirpStack Gateway Bridge **must** be configured with the `protobuf`
marshaler.

#### UI rewrite

ChirpStack v4 contains a rewrite of the ChirpStack Application Server v3 UI. The
new UI aims to be more user-friendly. Under the hood the API interface has been
ported to gRPC-web and all code has been ported to Typescript.

### Other changes

#### Passive Roaming improvements

The implementation of Passive Roaming has been improved, adding support for
appending `/sns` and `/fns` server endpoint suffixes. The usage of this suffix
is not specified in the Backend Interfaces specification, but is required by
some other network-server implementations.

#### UUID identifiers

All identifiers that are exposed have been changed to UUID. Previously most
identifiers (e.g. users, applications...) were incremental integers. In case
ChirpStack is setup as multi-tenant instance, this could expose some information
about the number of clients on the network. The migration script (see below)
will migrate these integers by converting these as strings, prefixed with zeros
in the UUID format. E.g. ID `123` would be converted to the UUID string
`00000000-0000-0000-0000-000000000123`.

#### String identifiers in API

All binary identifiers have been changed to string type in the API. While binary
fields are more efficient, these were confusing when encoded as JSON as the
Protobuf to JSON mapping uses base64 encoding for binary fields. For example, a
Gateway ID `0102030405060708` was encoded as `AQIDBAUGBwg=` in JSON.

#### JavaScript codec engine

The JavaScript codec engine is based on QuickJS, which is an embeddable
JavaScript engine which supports the ES2020 specification.

#### API interface changes

While the structure of API messages is roughly the same as the ChirpStack
Application Server API interface, some small changes have been made.

#### Integration events

The integration event messages have been restructured for better consistency.
Each event message has a `deviceInfo` field which holds device-related
information (tenant id & name, application id & name, device-profile id & name,
device EUI & name and tags).

### Development changes

#### Single repository

ChirpStack v4 will make it a lot easier to make customizations, especially when
API changes are involved, as API definitions are no longer separated from the
code. In v3 these definitions were moved to an external repository to avoid
cross dependencies.

#### Rust

For ChirpStack v4, it was decided to use Rust rather than Go. This was not an
easy choice and the arguments for this decision are debatable. However, as most
code was touched during the ChirpStack Application Server and ChirpStack Network
Server merge, it was the only moment to re-consider this. The Rust memory
management prevents many memory related pitfalls and helps catching bugs at
compile time rather than runtime.

### Migrating from v3 to v4

The recommended way to migrate from v3 to ChirpStack v4 is to create a new
PostgreSQL and Redis database and to use the
[ChirpStack v3 to v4](https://github.com/chirpstack/chirpstack-v3-to-v4)
migration script. This script will **copy** all the data from the "old" into the
"new" database. While the script does not make any modifications to the old
database, it is always recommended to make a backup first.
