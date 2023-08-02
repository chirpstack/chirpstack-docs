# Device profiles


A Device Profile defines the device capabilities and boot parameters
that are needed by ChirpStack for setting the LoRaWAN<sup>&reg;</sup> radio
access service. These information elements shall be provided by the
end-device manufacturer.

**Note**: RX and TX related parameters configured in the device-profile are
the parameters that ChirpStack will use on initial communication. However,
ChirpStack will re-configure the parameters using mac-commands to the
parameters set in the [Configuration](../configuration.md) file.

<!-- toc -->

## Device-profile templates

On create or update it is possible to use a device-profile template to fill in
most device-profile values. This also includes payload codecs. Please refer to
[Device-profile templates](device-profile-templates.md) for more information.

## Configuration

### General

The general tab configures the main properties of the device like the:

* Name
* Description
* MAC version
* Region
* Regional parameters revision
* ...

### Join (OTAA / ABP)

The Join (OTAA / ABP) tab defines the activation method of the device. In case
of ABP, you must configure the initial RX configuration.

### Class-B

For Class-B devices you must configure the Class-B confirmed downlink timeout
in (seconds), which is the time that ChirpStack will wait for a confirmation
on a confirmed downlink. If this time has passed, it will considered as not
acknowledged and the item will be discarded from the queue.

As well you will need to configure the (initial) Class-B ping-slot parameters
as configured by the device.

### Class-C

For Class-C devices you must configure the Class-C confirmed downlink timeout
in (seconds), which is the time that ChirpStack will wait for a confirmation
on a confirmed downlink.

### Codec

In this tab you can configure a payload codec. A codec will handle the encoding
and decoding of raw (binary) payloads.

**Note:** the raw payload will always be available, even when a codec has been
configured.

#### Cayenne LPP

If selecting the Cayenne LPP codec, ChirpStack will decode and encode
following the [Cayenne Low Power Payload](https://mydevices.com/cayenne/docs/lora/)
specification.

#### JavaScript functions

If selecting the Custom JavaScript codec functions option, you can write your
own (JavaScript) functions to decode an array of bytes to a JavaScript object
and encode a JavaScript object to an array of bytes. The JavaScript environment
in which the codecs are executed is based on [QuickJS](https://bellard.org/quickjs/)
which supports the ES2020 specification. The Node.js `Buffer` class is also
available.

Example template:

```js
// Decode uplink function.
//
// Input is an object with the following fields:
// - bytes = Byte array containing the uplink payload, e.g. [255, 230, 255, 0]
// - fPort = Uplink fPort.
// - variables = Object containing the configured device variables.
//
// Output must be an object with the following fields:
// - data = Object representing the decoded payload.
function decodeUplink(input) {
  return {
    data: {
      temp: 22.5
    }
  };
}

// Encode downlink function.
//
// Input is an object with the following fields:
// - data = Object representing the payload that must be encoded.
// - variables = Object containing the configured device variables.
//
// Output must be an object with the following fields:
// - bytes = Byte array containing the downlink payload.
function encodeDownlink(input) {
  return {
    bytes: [225, 230, 255, 0]
  };
}
```

### Relay

The Relay tab lets you configure Relay and Relay end-device configuration.
Please refer to [TS011](https://resources.lora-alliance.org/) for more
information about the Relay protocol.

#### Device is a Relay

A Relay device acts as an repeater between Relay capable end-devices and nearby
LoRa gateways.

If this option is enabled, you must configure the channel-configuration as well
as the bucket size and refresh rate. Please refer to _TS011_ for more
information about these properties.

#### Device is a Relay capable end-device

A Relay capable end-device is a device which is able to operate under a Relay.
For this it needs to implement the _TS011_ specification.

If this option is enabled, you must configure the channel-configuration, the
bucket size as well as the activation mode.

**Note:** To ignore uplinks from Relay capable end-devices that are directly
received by nearby LoRa gateways, you can enable **Only use Relay**.

### Tags

In this tab you can assign additional tags to the Device Profile. These tags
will be exposed on device events and can include additional metadata for
example:

* Vendor name
* Device model
* ...

### Measurements

If a codec has been configured, then ChirpStack will automatically detect the
measurements in the decoded uplink payload. Detected measurements can be found
under the **Measurements** tab of the device-profile.

Once a measurement kind has been selected, ChirpStack will start aggregating
these measurements and will visualize these under **Device metrics** on the
device dashboard. Aggregation is per hour for the last 24 hours, per day for
the last 31 days and per month for the last year.
