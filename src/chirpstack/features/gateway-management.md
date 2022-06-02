# Gateway management

ChirpStack has support for managing gateways. Gateways can be created in the web-interface or
using the API.

## Gateway location

The (last known) location of the gateway will be stored in the database. If
the gateway is equipped with a GPS, its location will be automatically updated
after every stats update in case it has changed. Else, it can be manually set
when creating or updating the gateway.

## Gateway re-configuration

If the [Concentratord](../../concentratord/index.md) is installed on the gateway
as packet-forwarder, ChirpStack will automatically configure the channel-plan
of the gateway and will keep it up-to-date with the gateway channels as configured
in the ChirpStack region configuration (please refer to the `[[regions.gateways.channels]]`
configuration option). See [region_eu868.toml](https://github.com/chirpstack/chirpstack/blob/master/chirpstack/configuration/region_eu868.toml)
as example.
