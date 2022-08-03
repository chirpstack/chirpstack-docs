# Quickstart Docker Compose

[Docker Compose](https://docs.docker.com/compose/) (part of Docker) makes
it possible to orchestrate the configuration of multiple Docker containers
at once using a `docker-compose.yml` file.

## Requirements

Before you continue, please make sure that you have Docker and Compose
installed. Please refer to [https://docs.docker.com/get-docker/](https://docs.docker.com/get-docker/)
for documentation on how to install Docker.

## Clone example repository

ChirpStack provides a repository with example Docker Compose configuration
to help you getting started with ChirpStack. This repository can be found
at [https://github.com/chirpstack/chirpstack-docker](https://github.com/chirpstack/chirpstack-docker).

To clone this repository, you can use the following command:

```bash
git clone https://github.com/chirpstack/chirpstack-docker.git
cd chirpstack-docker
```

## Configuration

ChirpStack configuration can be found in the `configuration/chirpstack`
directory. It contains the global `chirpstack.toml` configuration, as well
the configuration for all regions. The example configuration works out of the
box.

## Start

You can execute the following command to start ChirpStack and its dependencies:

```bash
docker-compose up
```

Please note that the first time you execute this command, there might be
some errors logged as the database needs to be initialized.

## Connect gateway

### UDP Packet Forwarder

The example Compose environment contains an instance of the ChirpStack Gateway
Bridge, pre-configured for the EU868 region and listening for UDP data on port
`1700`. Please consult the documentation of your gateway how you should
configure the UDP Packet Forwarder on your gateway.

#### Changing the region

To change the gateway region, open the `configuration/chirpstack-gateway-bridge/chirpstack-gateway-bridge.toml`
configuration file, and change the topic prefix to the prefix matching the
ChirpStack region configuration. After making configuration changes, make sure
to restart the Compose environment.

```bash
docker-compose stop
docker-compose up
```

### ChirpStack Gateway Bridge on gateway

Alternatively, you could install the [ChirpStack Gateway Bridge](../chirpstack-gateway-bridge/install/index.md)
on your gateway, and connect it to the Mosquitto MQTT Broker (part of this
`docker-compose.yml`). Please make sure that the `chirpstack-gateway-bridge.toml`
topic configuration matches the desired ChirpStack region.

## Import TTN device repository

To import the TTN [lorawan-devices](https://github.com/TheThingsNetwork/lorawan-devices)
repository (optional step), run the following command:

```bash
make import-lorawan-devices
```

This will clone the `lorawan-devices` repository and execute the `import-ttn-lorawan-devices`
command of ChirpStack. Please note that for this step, you need to have the `git` and `make`
commands installed.

## Login

Navigate to [http://localhost:8080](http://localhost:8080) and login with
`admin` / `admin`. In case the Compose environment is running on a different
host, you need to change the address in your browser.
