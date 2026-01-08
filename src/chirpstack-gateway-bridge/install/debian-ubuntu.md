# Debian / Ubuntu

## ChirpStack Debian / Ubuntu repository

Please see the [Downloads](../downloads.md) page for information on how to
activate the Debian / Ubuntu package repository.

## Install ChirpStack Gateway Bridge

In order to install ChirpStack Gateway Bridge, execute the following command:

```bash
sudo apt install chirpstack-gateway-bridge
```

This will setup an user and group, create start scripts for systemd or init.d
(this depends on your version of Debian / Ubuntu). The configuration file is
located at `/etc/chirpstack-gateway-bridge/chirpstack-gateway-bridge.toml`.

## Starting ChirpStack Gateway Bridge

```bash
sudo systemctl [start|stop|restart|status] chirpstack-gateway-bridge
```

## ChirpStack Gateway Bridge log output

Now you've setup ChirpStack Gateway Bridge and your gateway is configured to
forward it's data to it, it is a good time to verify that data is being
received. To view the ChirpStack Gateway Bridge logs, execute the following
command:

```bash
journalctl -u chirpstack-gateway-bridge -f -n 50
```
