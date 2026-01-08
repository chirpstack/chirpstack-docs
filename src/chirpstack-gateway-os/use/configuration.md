# Configuration

## Web-interface

ChirpStack Gateway OS can be configured using the provided web-interface. To
access the web-interface, enter `http://GATEWAY-IP-ADDRESS/` in your browser.
For example if the IP address of your gateway is `192.168.0.1`, then you need to
enter `http://192.168.0.1`.

The default credentials are:

- Username: `root`
- Password: _(not set)_

## SSH

It is also possible to configure the ChirpStack Gateway OS using a CLI over SSH.
Use the following command to SSH into the gateway:

```
ssh root@GATEWAY-IP-ADDRESS
```

ChirpStack Gateway OS uses the OpenWrt
[UCI system](https://openwrt.org/docs/guide-user/base-system/uci) for handling
configuration.

To show the current ChirpStack configurations:

```bash
uci show chirpstack-concentratord
uci show chirpstack-mqtt-forwarder
uci show chirpstack-udp-forwarder

# Only for Full images
uci show chirpstack
```

Please refer to the
[UCI system](https://openwrt.org/docs/guide-user/base-system/uci) guide for
information about using the UCI configuration system. After modifying ChirpStack
configuration and a `uci commit`, the updated services will be automatically
restarted.
