# Device profiles

A Device Profile defines the device capabilities and boot parameters
that are needed by ChirpStack for setting the LoRaWAN<sup>&reg;</sup> radio
access service. These information elements shall be provided by the
end-device manufacturer.

**Note**: RX and TX related parameters configured in the device-profile are
the parameters that ChirpStack will use on initial communication. However,
ChirpStack will re-configure the parameters using mac-commands to the
parameters set in the [Configuration](../configuration.md) file.

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
