# ThingsBoard

If configured, the ThingsBoard integration will send device attributes
and telemetry to the configured [ThingsBoard](https://thingsboard.io/) instance.

* [ThingsBoard guides](https://thingsboard.io/docs/guides/)

## Requirements

Before this integration is able to write data to ThingsBoard, the uplink
payloads must be decoded. The payload codec can be configured per
[Device Profile](../use/device-profiles.md). To validate that the uplink
payloads are decoded, you can use the [live device event-log](../use/event-logging.md)
feature. Decoded payload data will be available under the `object` key in
the JSON object.

ThingsBoard will generate a _Access Token_ per device. This token must be
configured as a [device variable](../use/devices.md) in ChirpStack. 
The variable must be named **ThingsBoardAccessToken**.

## Attributes

For each event, ChirpStack will update the ThingsBoard device with the
following attributes:

* `application_id`
* `application_name`
* `dev_eui`
* `device_name`

In case any [device tags](../use/devices.md) are configured for the
device in ChirpStack, these will also be added to the attributes.

## Telemetry

### Uplink

The following metrics are recorded for every uplink:

* `dr`
* `fcnt`
* `fport`
* `rssi`
* `snr`

Decoded uplink data is prefixed with the **data_** prefix. Make sure to
configure a coded in the [Device Profile](../use/device-profiles.md).

### Device-status

Device-status is prefixed with the **status_** prefix. The interval of the
device-status requests can be configured through the [Device Profile](../use/device-profiles.md).

### Location

Location data is prefixed with the **location_** prefix. Please note that this
is **only** available in case geolocation capable gateways are being used and
ChirpStack is configured with geolocation support.
