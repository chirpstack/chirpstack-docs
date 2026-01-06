# Building new integrations

ChirpStack provides three kinds of integrations.

* Per application integrations
* Global integrations
* Standalone integrations

## Per application integrations

These integrations can be configured per application through the ChirpStack
web-interface. As these are user-configurable, this kind of integration
is stateless thus it does not maintain a connection to the integration.
Adding such integration requires implementing the integration, adding API
methods and modifying the ChirpStack UI.

## Global integrations

These integrations are provided by ChirpStack and once configured, will be
globally enabled. An example is the [MQTT](./mqtt.md) integration. Adding
such integration requires implementing the integration. No API or UI changes
are required.

## Standalone integrations

Standalone integrations operate independent from ChirpStack, that means no
ChirpStack changes will be required to add a new integration of this kind.
These integrations will subscribe to a [Redis Stream](https://redis.io/docs/data-types/streams/)
to which ChirpStack publishes integration events. Please see the
`device_frame_log_max_history` setting in the [ChirpStack Configuration](../configuration.md)
to enable this feature.

```d2
vars: {
	d2-config: {
		layout-engine: elk
	}
}

ChirpStack -> Redis Stream
Redis Stream -> Integration 1
Redis Stream -> Integration 2
Redis Stream -> Integration n
```

ChirpStack provides a [`chirpstack-intergration`](https://crates.io/crates/chirpstack-integration)
crate which provides the Redis Stream interface and a framework to easily build
integrations. The [Apache Pulsar](./pulsar.md) integration can be used as an
example implementation.
