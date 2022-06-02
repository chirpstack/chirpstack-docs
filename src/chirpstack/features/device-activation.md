# Device activation

## Over The Air Activation (OTAA)

For OTAA devices, it is possible to configure the root-keys within ChirpStack
or use an external join-server. In the latter case, you must configure the
join-server(s) in the ChirpStack [Configuration](../configuration.md) file.

If ChirpStack receives an OTAA join-request with a JoinEUI that is configured
in the configuration file, it will forward the join-request to this join-server.
In the other case ChirpStack will use the root-keys that are configured within
ChirpStack, or raise an error if no keys have been configured.

The Join Server API is specified by the [LoRaWAN Backend Interfaces](https://www.lora-alliance.org/lorawan-for-developers)

## Activation By Personalization (ABP)

ChirpStack also support ABP devices. In this case devices must be activated
through the web-interface or API by entering the session-keys. Once activated,
ChirpStack will handle these devices the same as OTAA activated devices.
