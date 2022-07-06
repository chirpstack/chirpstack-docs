# IFTTT

If configured, the IFTTT integration will forward device measurements to
the [IFTTT](https://ifttt.com/) [Webhooks](https://ifttt.com/maker_webhooks).
For this ChirpStack uses the event trigger with 3 JSON values format. The first
value contains the device DevEUI, the other two values can be freely assigned.

## Preparation

### Key

Before you can setup the IFTTT integration, you must obtain the key to
authenticate with the IFTTT API. Please refer to the [Webhooks](https://ifttt.com/maker_webhooks)
documentation. On top of the documentation page, you will find _Your key is:_ ....

### Codec

You must configure a payload codec in the device-profile in order to use the
IFTTT integration. See also [Device profiles](../use/device-profiles.md).


## Configuration

* **Key**: Please use the key obtained in the previous step.
* **Value 1 & 2 keys**: This must map to the measurement keys in the decoded
  payload. Please note that ChirpStack will automatically detect these when
  an uplink is received, and will store these under **Measurements** in the
  device-profile.
