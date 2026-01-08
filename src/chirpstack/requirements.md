# Requirements

ChirpStack uses several external components for storage of data and receiving
gateway events and sending gateway commands.

## MQTT broker

By default, ChirpStack uses MQTT for subscribing to gateway events and sending
commands to the gateway. [Mosquitto](http://mosquitto.org) is a popular
open-source MQTT broker. Any other MQTT broker can be used, as long as it
implements the MQTT v5.0 specification and supports shared-subscriptions.

### Install

#### Debian / Ubuntu

To install Mosquitto:

```bash
sudo apt install mosquitto
```

#### Other platforms

Please refer to the [Mosquitto download](https://mosquitto.org/download/) page
for information about how to setup Mosquitto for your platform.

## PostgreSQL database

ChirpStack uses [PostgreSQL](https://www.postgresql.org) as persistent
data-storage. PostgreSQL v13+ is required.

### Install

#### Debian / Ubuntu

To install the PostgreSQL:

```bash
sudo apt install postgresql
```

#### Other platforms

Please refer to the [PostgreSQL download](https://www.postgresql.org/download/)
page for information how to setup PostgreSQL on your platform.

### Create database

Once PostgreSQL has been installed, you can use the following command to create
a `chirpstack` role with `chirpstack` as password and a `chirpstack` database.
**Note:** the `pg_trgm` extension must be installed (see example below).

From the CLI:

```bash
sudo -u postgres psql
```

Within the SQL console:

```sql
-- create role for authentication
create role chirpstack with login password 'chirpstack';

-- create database
create database chirpstack with owner chirpstack;

-- change to chirpstack database
\c chirpstack

-- create pg_trgm extension
create extension pg_trgm;

-- exit psql
\q
```

## SQLite database

As an alternative to PostgreSQL, ChirpStack can use
[SQLite](https://www.sqlite.org) as persistent data-storage.

There is no extra software to install or configure to use SQLite database.

Make sure to run the `chirpstack-sqlite` binary to use SQLite database. It can
be found in
[Downloads](https://www.chirpstack.io/docs/chirpstack/downloads.html).

## Redis database

ChirpStack stores device-sessions, metrics and cache into a
[Redis](http://redis.io) database. Note that at least Redis 6.2.0 is required.

### Install

#### Debian / Ubuntu

To Install Redis:

```bash
sudo apt install redis-server
```

#### Other platforms

Please refer to the [Redis](https://redis.io/) documentation for information
about how to setup Redis for your platform.
