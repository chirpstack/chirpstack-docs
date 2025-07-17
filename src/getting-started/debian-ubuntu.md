# Quickstart Debian / Ubuntu

This tutorial describes the steps needed to setup ChirpStack
on a single machine. Recommended distributions are:

* Ubuntu (latest LTS)
* Debian (latest stable)

## Note before you start

The purpose of this tutorial is to help you getting started. Topics like
HA setup, database backups, setting up firewall rules are beyond the scope
of this tutorial.

## Install requirements

The following command installs the ChirpStack requirements using the `apt`
package manager:

```bash
sudo apt install \
	mosquitto \
	mosquitto-clients \
	redis-server \
	redis-tools \
	postgresql
```

## PostgreSQL setup

To enter the command line utility for PostgreSQL:

```bash
sudo -u postgres psql
```

Inside this prompt, execute the following queries to setup the ChirpStack
database. It is recommended to change the usernames and passwords.

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

## Setup software repository

ChirpStack provides a Debian / Ubuntu repository which can be used to install
the latest ChirpStack version. First make sure that `gpg` is installed:

```bash
sudo apt install gpg
```

Set up the key for the ChirpStack repository:

```bash
sudo mkdir -p /etc/apt/keyrings/
sudo sh -c 'wget -q -O - https://artifacts.chirpstack.io/packages/chirpstack.key | gpg --dearmor > /etc/apt/keyrings/chirpstack.gpg'
```

Add the repository to the repository list:

```bash
echo "deb [signed-by=/etc/apt/keyrings/chirpstack.gpg] https://artifacts.chirpstack.io/packages/4.x/deb stable main" | sudo tee /etc/apt/sources.list.d/chirpstack.list
```

Update the apt package cache:

```bash
sudo apt update
```

## Install ChirpStack Gateway Bridge

**Note:** If you intend to install the ChirpStack Gateway Bridge only on
gateway(s) themselves, you can skip this step.

Install the package using `apt`:

```bash
sudo apt install chirpstack-gateway-bridge
```

### Configuration

The configuration file is located at `/etc/chirpstack-gateway-bridge/chirpstack-gateway-bridge.toml`.
Please update the `[integration.mqtt]` section to match the region prefix for
the region that applies to this ChirpStack Gateway Bridge instance.

Example for EU868:

```toml
{%raw%}[integration.mqtt]
event_topic_template="eu868/gateway/{{ .GatewayID }}/event/{{ .EventType }}"
command_topic_template="eu868/gateway/{{ .GatewayID }}/command/#"{%endraw%}
```

You can find the region configurations that are included by default here:
[https://github.com/chirpstack/chirpstack/tree/master/chirpstack/configuration](https://github.com/chirpstack/chirpstack/tree/master/chirpstack/configuration).

### Starting

Start the ChirpStack Gateway Bridge service:

```bash
# start chirpstack-gateway-bridge
sudo systemctl start chirpstack-gateway-bridge

# start chirpstack-gateway-bridge on boot
sudo systemctl enable chirpstack-gateway-bridge
```

## Install ChirpStack

Install the package using apt:

```bash
apt install chirpstack
```

### Configuration

The configuration files are located at `/etc/chirpstack`. You will find one
global configuration file called `chirpstack.toml` and various region
configuration files.

### Starting

Start the ChirpStack service:

```bash
# start chirpstack
sudo systemctl start chirpstack

# start chirpstack on boot
sudo systemctl enable chirpstack
```

Print the ChirpStack log-output:

```bash
sudo journalctl -f -n 100 -u chirpstack
```
