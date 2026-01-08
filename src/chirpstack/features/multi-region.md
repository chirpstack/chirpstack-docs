# Multi-region

ChirpStack supports multiple regions simultaneously, without the need to start
multiple ChirpStack instances (like it was needed in ChirpStack v3). It is also
possible to configure the same region more than once, to support for example a
mix of 8-channel and 16-channel gateways.

There are two terms that are important:

- **Region name**, which is a user-defined name for a region configuration, this
  can be anything.
- **Region common name**, which is the region common name as defined by the LoRa
  Alliance.

In the ChirpStack configuration examples, you will find for example region name
`us915_0` for the channels 0 - 7 configuration, `us915_1` for the channels 8 -
15 configuration etc...

Each region configuration contains its own gateway backend configuration. Either
a single MQTT broker can be used for all different regions, or multiple MQTT
brokers can be used (e.g. for each region a separate MQTT broker).

ChirpStack will automatically assign the _region name_ and _region common name_
to each gateway, based on the receiving gateway backend.

## Single MQTT broker

The advantage of a single MQTT broker is the ease of setup. Adding a new region
is changing the ChirpStack configuration and there is no need to deploy new MQTT
brokers.

However, it is important that each region uses its own MQTT topic prefix, to
separate the MQTT message flows between regions. It is a good practice to use
the _region name_ as topic prefix for the gateway.

## Multiple MQTT brokers

The advantage of multiple MQTT brokers can be scalability. Each region
configuration can have its own MQTT broker, and therefore the load of one region
does not impact the load of an other region (MQTT broker).

As message flows for the different region configurations is separated by MQTT
broker instance, there is no need to add a region specific topic prefix.
