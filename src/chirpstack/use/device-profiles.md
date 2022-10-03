# Device profiles

A Device Profile defines the device capabilities and boot parameters
that are needed by ChirpStack for setting the LoRaWAN<sup>&reg;</sup> radio
access service. These information elements shall be provided by the
end-device manufacturer.

**Note**: RX and TX related parameters configured in the device-profile are
the parameters that ChirpStack will use on initial communication. However,
ChirpStack will re-configure the parameters using mac-commands to the
parameters set in the [Configuration](../configuration.md) file.

## Device-profile templates

On create or update it is possible to use a device-profile template to fill in
most device-profile values. This also includes payload codecs. Please refer to
[Device-profile templates](device-profile-templates.md) for more information.

## Payload codecs

**Note:** the raw payload will always be available, even when a codec has been
configured.

### Cayenne LPP

When selecting the Cayenne LPP codec, ChirpStack will decode and encode
following the [Cayenne Low Power Payload](https://mydevices.com/cayenne/docs/lora/)
specification.

### Custom JavaScript codec functions

When selecting the Custom JavaScript codec functions option, you can write your
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

## Measurements

If a codec has been configured, then ChirpStack will automatically detect the
measurements in the decoded uplink payload. Detected measurements can be found
under the **Measurements** tab of the device-profile.

Once a measurement kind has been selected, ChirpStack will start aggregating
these measurements and will visualize these under **Device metrics** on the
device dashboard. Aggregation is per hour for the last 24 hours, per day for
the last 31 days and per month for the last year.
