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
at [https://github.com/brocaar/chirpstack-docker](https://github.com/brocaar/chirpstack-docker).

To clone this repository, you can use the following command:

```bash
git clone https://github.com/brocaar/chirpstack-docker.git
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

The example Compose environment contains an instance of the ChirpStack Gateway
Bridge, pre-configured for the EU868 region and listening for UDP data on port
`1700`. Please consult the documentation of your gateway how you should
configure your gateway.

### Changing the region

To change the gateway region, open the `configuration/chirpstack-gateway-bridge/chirpstack-gateway-bridge.toml`
configuration file, and change the topic prefix to the prefix matching the
ChirpStack region configuration. After making configuration changes, make sure
to restart the Compose environment.

```bash
docker-compose stop
docker-compose up
```

## Login

Navigate to [http://localhost:8080](http://localhost:8080) and login with
`admin` / `admin`. In case the Compose environment is running on a different
host, you need to change the address in your browser.
