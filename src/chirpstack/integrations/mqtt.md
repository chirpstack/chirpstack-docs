# MQTT

The MQTT integration publishes all the data it receives from the devices
as JSON over MQTT. To receive data from your device, you therefore 
need to subscribe to its MQTT topic. For debugging, you could use a 
(command-line) tool like `mosquitto_sub` which is part of the 
[Mosquitto](http://mosquitto.org/) MQTT broker.

### MQTT quickstart

Use `+` for a single-level wildcard, `#` for a multi-level wildcard.
Examples:

```bash
mosquitto_sub -t "application/APPLICATION_ID/#" -v                  # display everything for the given APPLICATION_ID
mosquitto_sub -t "application/APPLICATION_ID/device/+/event/up" -v  # display only the uplink payloads for the given APPLICATION_ID
```

Note:

* MQTT topics are case-sensitive
* The `APPLICATION_ID` can be retrieved using the API or using the web-interface,
  this is not the same as the `AppEUI` / `JoinEUI`!

## Events

The MQTT integration exposes all events as documented by [Event types](events.md).
The default event topic is: `application/APPLICATION_ID/device/DEV_EUI/event/EVENT`.

## Scheduling a downlink

The default topic for scheduling downlink payloads is: `application/APPLICATION_ID/device/DEV_EUI/command/down`.

The Application ID and DevEUI of the device will be taken from the topic.

Example payload:

```json
{
    "confirmed": true,                        // whether the payload must be sent as confirmed data down or not
    "fPort": 10,                              // FPort to use (must be > 0)
    "data": "...."                            // base64 encoded data (plaintext, will be encrypted by ChirpStack)
    "object": {                               // decoded object (when application coded has been configured)
        "temperatureSensor": {"1": 25},       // when providing the 'object', you can omit 'data'
        "humiditySensor": {"1": 32}
    }
}
```
