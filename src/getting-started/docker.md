# Quickstart Docker Compose

[Docker Compose](https://docs.docker.com/compose/) (part of Docker) makes it
possible to orchestrate the configuration of multiple Docker containers at once
using a `docker-compose.yml` file.

## Requirements

Before you continue, please make sure that you have Docker and Compose
installed. Please refer to
[https://docs.docker.com/get-docker/](https://docs.docker.com/get-docker/) for
documentation on how to install Docker.

## Clone example repository

ChirpStack provides a repository with example Docker Compose configuration to
help you getting started with ChirpStack. This repository can be found at
[https://github.com/chirpstack/chirpstack-docker](https://github.com/chirpstack/chirpstack-docker).

To clone this repository, you can use the following command:

```bash
git clone https://github.com/chirpstack/chirpstack-docker.git
cd chirpstack-docker
```

## Configuration

Please refer to the [README.md](https://github.com/chirpstack/chirpstack-docker)
for configuration instructions. By default the EU868 region is configured but
this can be changed to any supported region.

## Start

Execute the following command to start ChirpStack and its dependencies:

```bash
docker-compose up
```

Please note that the first time you execute this command, there might be some
errors logged as the database needs to be initialized.

## Connect gateway

### UDP Packet Forwarder

The example Compose environment provides an instance of the ChirpStack Gateway
Bridge, by default configured for the EU868 region and listening for UDP data on
port `1700`. Please consult the documentation of your gateway for configuration
instructions.

### ChirpStack MQTT

#### ChirpStack MQTT Forwarder

If you have installed the
[ChirpStack MQTT Forwarder](../chirpstack-mqtt-forwarder/install/index.md) on
your gateway or are planning to do so, then you can connect it to the MQTT
broker (part of the `docker-compose.yml`). For this you must update the
ChirpStack MQTT Forwarder configuration (on your gateway) which can be found in
`chirpstack-mqtt-forwarder.toml`. Please make sure that the
`chirpstack-mqtt-forwarder.toml` topic-prefix configuration matches the desired
ChirpStack region.

#### ChirpStack Gateway Bridge

If you have installed the
[ChirpStack Gateway Bridge](../chirpstack-gateway-bridge/install/index.md) on
your gateway or are planning to do so, then you can connect it to the MQTT
broker (part of the `docker-compose.yml`). For this you must update the
ChirpStack Gateway Bridge configuration (on your gateway) which can be found in
`chirpstack-gateway-bridge.toml`. Please make sure that the
`chirpstack-gateway-bridge.toml` topic configuration matches the desired
ChirpStack region.

### Semtech Basics Station

The example Compose environment provides an instance of the ChirpStack Gateway
Bridge with Basics Station backend listening on port `3001`. Please consult the
documentation of your gateway for configuration instructions.

## Import LoRaWAN device repository

To import the TTN
[lorawan-devices](https://github.com/TheThingsNetwork/lorawan-devices)
repository (optional step), run the following command:

```bash
make import-lorawan-devices
```

This will clone the `lorawan-devices` repository and execute the import command
of ChirpStack. Please note that for this step, you need to have the `git` and
`make` commands installed.

## Login

Navigate to [http://localhost:8080](http://localhost:8080) and login with
`admin` / `admin`. In case the Compose environment is running on a different
host, you need to change the address in your browser.
