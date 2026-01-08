# Subscribe to MQTT

The **Full** image comes with [Mosquitto](https://mosquitto.org/) as MQTT broker
pre-installed.

## Subscribing

You can use the following command to receive data:

```bash
mosquitto_sub -h localhost -t "application/#" -v
```

If you are connecting from a different host, then you must change `localhost` to
the IP address of your gateway.
