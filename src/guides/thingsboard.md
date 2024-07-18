# ThingsBoard getting started

[ThingsBoard](https://www.thingsboard.io) is an Open-core IoT Platform for
Device management, data collection, processing and visualization for your IoT
solution.

This guide describes how to setup ChirpStack Professional Edition so that it forwards device-data
to [ThingsBoard](https://www.thingsboard.io) for processing and visualization.

## Installing ChirpStack components

The installation of the ChirpStack components is not covered in this guide. The easiest
option to get started is the [Docker Compose](../getting-started/docker.md), on which this
guide continues.

Before you continue, make sure that you have a working ChirpStack environment with activated Professional Licence.

## Install ThingsBoard

ThingsBoard can be installed in many different ways. Please refer to the
[ThingsBoard Installation Documentation](https://thingsboard.io/docs/installation/)
for more information.

### Docker Compose setup

When you have installed ChirpStack using the Docker Compose guide, the easiest
option to integrate with ThingsBoard is by adding the Thingsboard Docker
container to the ChirpStack `docker-compose.yml` file.

After making the modifications mentioned below, ThingsBoard will be started
together with all ChirpStack components when running a `docker-compose up`.

**Note:** The ThingsBoard server used in one of the next steps will be
`http://thingsboard:9090/`.

#### Service

Add the following snippet under `services`:

```yaml
  thingsboard:
    image: thingsboard/tb-postgres
    volumes:
      - thingsboarddata:/data
    ports:
      - 9090:9090
```

Your `docker-compose.yml` file now looks like:

```yaml
version: "3"

services:
  chirpstack:
    image: chirpstack/chirpstack:4
    volumes:
      - ./configuration/chirpstack:/etc/chirpstack

  [...]

  thingsboard:
    image: thingsboard/tb-postgres
    volumes:
      - thingsboarddata:/data
    ports:
      - 9090:9090

  [...]
```

#### Volume

Add the following under `volumes`:

```yaml
  thingsboarddata:
```

Your `docker-compose.yml` file now looks like:

```yaml
[...]

volumes:
  postgresqldata:
  redisdata:
  thingsboarddata:
```

## Setup ThingsBoard

For getting started with ThingsBoard, please refer to the
[ThingsBoard Getting Started](https://thingsboard.io/docs/getting-started-guides/helloworld/)
guide. The important thing is that you have created Device within ThingsBoard.

## ThingsBoard integration

### Get Device Auth Token

In order to let ChirpStack push data to your ThingsBoard device, you need
to obtain the ThingsBoard Device _Access Token_. Within ThingsBoard, open your
Device and click the **Copy Access Token** button. This will copy the
_Access Token_ to your clipboard.

### Set Device Auth Token in ChirpStack

Now open the ChirpStack web-interface and navigate to the Device. Click
**Configuration**, then click **Variables**.

Add a variable named **ThingsBoardAccessToken** and with as value the content
from your clipboard (containing the ThingsBoard Device _Access Token_).

### Setup ChirpStack ThingsBoard integration

Inside the ChirpStack web-interface, navigate to the Application to
which the Device belongs. Click **Integrations**, then under **ThingsBoard**
click **+**.

* ThingsBoard.io server: Usually this is **http://host:9090** (where **host**
  is replaced by the hostname of the server serving ThingsBoard). When you are
  using the Docker Compose instructions, set this to **http://thingsboard:9090**.

## Validate integration

When you have completed all the steps, then ThingsBoard is ready to receive uplink
data (or telemetry) and ChirpStack is setup to forward data for your
Device, using the _Access Token_ for authentication.

The last step is to let your device send some data and validate that this data
is received by Thingsboard. You will find this data under the _Latest Telemetry_
tab when navigating to the Device within Thingsboard.
