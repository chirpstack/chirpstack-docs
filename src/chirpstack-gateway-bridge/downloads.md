# Downloads

## Binaries

| File name                                                                                                                                                                                                        | Type    | OS    | Arch  |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------- | ----- | ----- |
| [chirpstack-gateway-bridge_%GATEWAY_BRIDGE_VERSION_linux_amd64.deb](https://artifacts.chirpstack.io/downloads/chirpstack-gateway-bridge/chirpstack-gateway-bridge_%GATEWAY_BRIDGE_VERSION_linux_amd64.deb)       | .deb    | Linux | amd64 |
| [chirpstack-gateway-bridge_%GATEWAY_BRIDGE_VERSION_linux_armhf.deb](https://artifacts.chirpstack.io/downloads/chirpstack-gateway-bridge/chirpstack-gateway-bridge_%GATEWAY_BRIDGE_VERSION_linux_armhf.deb)       | .deb    | Linux | armv7 |
| [chirpstack-gateway-bridge_%GATEWAY_BRIDGE_VERSION_linux_arm64.deb](https://artifacts.chirpstack.io/downloads/chirpstack-gateway-bridge/chirpstack-gateway-bridge_%GATEWAY_BRIDGE_VERSION_linux_arm64.deb)       | .deb    | Linux | arm64 |
| [chirpstack-gateway-bridge_%GATEWAY_BRIDGE_VERSION_linux_amd64.tar.gz](https://artifacts.chirpstack.io/downloads/chirpstack-gateway-bridge/chirpstack-gateway-bridge_%GATEWAY_BRIDGE_VERSION_linux_amd64.tar.gz) | .tar.gz | Linux | amd64 |
| [chirpstack-gateway-bridge_%GATEWAY_BRIDGE_VERSION_linux_armv7.tar.gz](https://artifacts.chirpstack.io/downloads/chirpstack-gateway-bridge/chirpstack-gateway-bridge_%GATEWAY_BRIDGE_VERSION_linux_armv7.tar.gz) | .tar.gz | Linux | armv7 |
| [chirpstack-gateway-bridge_%GATEWAY_BRIDGE_VERSION_linux_arm64.tar.gz](https://artifacts.chirpstack.io/downloads/chirpstack-gateway-bridge/chirpstack-gateway-bridge_%GATEWAY_BRIDGE_VERSION_linux_arm64.tar.gz) | .tar.gz | Linux | arm64 |
| [chirpstack-gateway-bridge_%GATEWAY_BRIDGE_VERSION_linux_mips.tar.gz](https://artifacts.chirpstack.io/downloads/chirpstack-gateway-bridge/chirpstack-gateway-bridge_%GATEWAY_BRIDGE_VERSION_linux_mips.tar.gz)   | .tar.gz | Linux | mips  |
| [chirpstack-gateway-bridge_%GATEWAY_BRIDGE_VERSION_linux_amd64.rpm](https://artifacts.chirpstack.io/downloads/chirpstack-gateway-bridge/chirpstack-gateway-bridge_%GATEWAY_BRIDGE_VERSION_linux_amd64.rpm)       | .rpm    | Linux | amd64 |
| [chirpstack-gateway-bridge_%GATEWAY_BRIDGE_VERSION_linux_armv7.rpm](https://artifacts.chirpstack.io/downloads/chirpstack-gateway-bridge/chirpstack-gateway-bridge_%GATEWAY_BRIDGE_VERSION_linux_armv7.rpm)       | .rpm    | Linux | armv7 |
| [chirpstack-gateway-bridge_%GATEWAY_BRIDGE_VERSION_linux_arm64.rpm](https://artifacts.chirpstack.io/downloads/chirpstack-gateway-bridge/chirpstack-gateway-bridge_%GATEWAY_BRIDGE_VERSION_linux_arm64.rpm)       | .rpm    | Linux | arm64 |

## Debian / Ubuntu repository

ChirpStack provides pre-compiled binaries as Debian (`.deb`) packages. To
activate this repository, execute the following commands:

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

## Docker

For Docker images, please refer to
[https://hub.docker.com/r/chirpstack/chirpstack-gateway-bridge](https://hub.docker.com/r/chirpstack/chirpstack-gateway-bridge/tags)
