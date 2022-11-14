# AMQP / RabbitMQ

The [AMQP](https://en.wikipedia.org/wiki/Advanced_Message_Queuing_Protocol) /
[RabbitMQ](https://www.rabbitmq.com/) integration publishes events to an AMQP
routing-key. By creating one or multiple bindings to one or multiple queues,
it is possible to subscribe to all data, or just a sub-set.

## Activating integration

This integration must be configured in the [Configuration](../configuration.md) file.

### Enable integration

In the configuration file:

```toml
[integration]
  enabled = [
    "mqtt",
  ]
```

Your enabled line may look slightly different. Add `"amqp"`` to the list. In
this case, the modified line should appear as enabled=["mqtt", "amqp"].

### Configuration

You must also add / update the `[integration.amqp]` section with the URL of
the AMQP broker URL and the routing-key. Example:

```toml
# AMQP / RabbitMQ integration configuration.
{%raw %}[integration.amqp]
  url="amqp://guest:guest@localhost:5672"
  event_routing_key="application.{{application_id}}.device.{{dev_eui}}.event.{{event}}"
  json=true{% endraw %}
```
