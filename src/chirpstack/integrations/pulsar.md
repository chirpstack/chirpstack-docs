# Apache Pulsar

ChirpStack Pulsar integration publishes integration events to an
[Apache Pulsar](https://pulsar.apache.org/) instance. It reads integration
events by directly subscribing to the Redis event
[stream](https://redis.io/docs/data-types/streams/).

**Note:** this integration must be installed as an additional component.

<!-- toc -->

## Installation

### Binaries

| File name                                                                                                                                                                                                                                | Type    | OS    | Arch  |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------- | ----- | ----- |
| [chirpstack-pulsar-integration_%PULSAR_INTEGRATION_VERSION_linux_amd64.deb](https://artifacts.chirpstack.io/downloads/chirpstack-pulsar-integration/chirpstack-pulsar-integration_%PULSAR_INTEGRATION_VERSION_linux_amd64.deb)           | .deb    | Linux | amd64 |
| [chirpstack-pulsar-integration_%PULSAR_INTEGRATION_VERSION_linux_armhf.deb](https://artifacts.chirpstack.io/downloads/chirpstack-pulsar-integration/chirpstack-pulsar-integration_%PULSAR_INTEGRATION_VERSION_linux_armhf.deb)           | .deb    | Linux | armv7 |
| [chirpstack-pulsar-integration_%PULSAR_INTEGRATION_VERSION_linux_arm64.deb](https://artifacts.chirpstack.io/downloads/chirpstack-pulsar-integration/chirpstack-pulsar-integration_%PULSAR_INTEGRATION_VERSION_linux_arm64.deb)           | .deb    | Linux | arm64 |
| [chirpstack-pulsar-integration_%PULSAR_INTEGRATION_VERSION_linux_amd64.tar.gz](https://artifacts.chirpstack.io/downloads/chirpstack-pulsar-integration/chirpstack-pulsar-integration_%PULSAR_INTEGRATION_VERSION_linux_amd64.tar.gz)     | .tar.gz | Linux | amd64 |
| [chirpstack-pulsar-integration_%PULSAR_INTEGRATION_VERSION_linux_armv7hf.tar.gz](https://artifacts.chirpstack.io/downloads/chirpstack-pulsar-integration/chirpstack-pulsar-integration_%PULSAR_INTEGRATION_VERSION_linux_armv7hf.tar.gz) | .tar.gz | Linux | armv7 |
| [chirpstack-pulsar-integration_%PULSAR_INTEGRATION_VERSION_linux_arm64.tar.gz](https://artifacts.chirpstack.io/downloads/chirpstack-pulsar-integration/chirpstack-pulsar-integration_%PULSAR_INTEGRATION_VERSION_linux_arm64.tar.gz)     | .tar.gz | Linux | arm64 |
| [chirpstack-pulsar-integration_%PULSAR_INTEGRATION_VERSION_linux_amd64.rpm](https://artifacts.chirpstack.io/downloads/chirpstack-pulsar-integration/chirpstack-pulsar-integration_%PULSAR_INTEGRATION_VERSION_linux_amd64.rpm)           | .rpm    | Linux | amd64 |
| [chirpstack-pulsar-integration_%PULSAR_INTEGRATION_VERSION_linux_armv7.rpm](https://artifacts.chirpstack.io/downloads/chirpstack-pulsar-integration/chirpstack-pulsar-integration_%PULSAR_INTEGRATION_VERSION_linux_armv7.rpm)           | .rpm    | Linux | armv7 |
| [chirpstack-pulsar-integration_%PULSAR_INTEGRATION_VERSION_linux_arm64.rpm](https://artifacts.chirpstack.io/downloads/chirpstack-pulsar-integration/chirpstack-pulsar-integration_%PULSAR_INTEGRATION_VERSION_linux_arm64.rpm)           | .rpm    | Linux | arm64 |

### Setup software repository

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

### Docker

For Docker images, please refer to
[https://hub.docker.com/r/chirpstack/chirpstack-pulsar-integration](https://hub.docker.com/r/chirpstack/chirpstack-pulsar-integration/tags)

## Configuration

### ChirpStack Pulsar Integration

To list all CLI options, start `chirpstack-pulsar-integration` with the `--help`
flag. This will print:

```text
Apache Pulsar integration for ChirpStack

Usage: chirpstack-pulsar-integration --config <FILE>

Options:
  -c, --config <FILE>  
  -h, --help           Print help
  -V, --version        Print version
```

Example configuration file:

```toml
{{#include chirpstack-pulsar-integration.toml}}
```

### ChirpStack configuration

The Redis Stream storing the integration events retains a limited amount of
payloads. This number can be configured by setting the
`device_event_log_max_history` option in the
[ChirpStack Configuration](../configuration.md).

## Source

Please refer to
[https://github.com/chirpstack/chirpstack-pulsar-integration/](https://github.com/chirpstack/chirpstack-pulsar-integration/)
for source-code.
