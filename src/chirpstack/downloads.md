# Downloads

## Binaries

| File name                                                                                                                                                                         | Type    | DB         | OS    | Arch  |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------- | ---------- | ----- | ----- |
| [chirpstack_%CHIRPSTACK_VERSION_linux_amd64.deb](https://artifacts.chirpstack.io/downloads/chirpstack/chirpstack_%CHIRPSTACK_VERSION_linux_amd64.deb)                             | .deb    | PostgreSQL | Linux | amd64 |
| [chirpstack_%CHIRPSTACK_VERSION_linux_armhf.deb](https://artifacts.chirpstack.io/downloads/chirpstack/chirpstack_%CHIRPSTACK_VERSION_linux_armhf.deb)                             | .deb    | PostgreSQL | Linux | armv7 |
| [chirpstack_%CHIRPSTACK_VERSION_linux_arm64.deb](https://artifacts.chirpstack.io/downloads/chirpstack/chirpstack_%CHIRPSTACK_VERSION_linux_arm64.deb)                             | .deb    | PostgreSQL | Linux | arm64 |
| [chirpstack-sqlite_%CHIRPSTACK_VERSION_linux_amd64.deb](https://artifacts.chirpstack.io/downloads/chirpstack/chirpstack-sqlite_%CHIRPSTACK_VERSION_linux_amd64.deb)               | .deb    | SQLite     | Linux | amd64 |
| [chirpstack-sqlite_%CHIRPSTACK_VERSION_linux_armhf.deb](https://artifacts.chirpstack.io/downloads/chirpstack/chirpstack-sqlite_%CHIRPSTACK_VERSION_linux_armhf.deb)               | .deb    | SQLite     | Linux | armv7 |
| [chirpstack-sqlite_%CHIRPSTACK_VERSION_linux_arm64.deb](https://artifacts.chirpstack.io/downloads/chirpstack/chirpstack-sqlite_%CHIRPSTACK_VERSION_linux_arm64.deb)               | .deb    | SQLite     | Linux | arm64 |
| [chirpstack_%CHIRPSTACK_VERSION_postgres_linux_amd64.tar.gz](https://artifacts.chirpstack.io/downloads/chirpstack/chirpstack_%CHIRPSTACK_VERSION_postgres_linux_amd64.tar.gz)     | .tar.gz | PostgreSQL | Linux | amd64 |
| [chirpstack_%CHIRPSTACK_VERSION_postgres_linux_armv7hf.tar.gz](https://artifacts.chirpstack.io/downloads/chirpstack/chirpstack_%CHIRPSTACK_VERSION_postgres_linux_armv7hf.tar.gz) | .tar.gz | PostgreSQL | Linux | armv7 |
| [chirpstack_%CHIRPSTACK_VERSION_postgres_linux_arm64.tar.gz](https://artifacts.chirpstack.io/downloads/chirpstack/chirpstack_%CHIRPSTACK_VERSION_postgres_linux_arm64.tar.gz)     | .tar.gz | PostgreSQL | Linux | arm64 |
| [chirpstack_%CHIRPSTACK_VERSION_sqlite_linux_amd64.tar.gz](https://artifacts.chirpstack.io/downloads/chirpstack/chirpstack_%CHIRPSTACK_VERSION_sqlite_linux_amd64.tar.gz)         | .tar.gz | SQLite     | Linux | amd64 |
| [chirpstack_%CHIRPSTACK_VERSION_sqlite_linux_armv7hf.tar.gz](https://artifacts.chirpstack.io/downloads/chirpstack/chirpstack_%CHIRPSTACK_VERSION_sqlite_linux_armv7hf.tar.gz)     | .tar.gz | SQLite     | Linux | armv7 |
| [chirpstack_%CHIRPSTACK_VERSION_sqlite_linux_arm64.tar.gz](https://artifacts.chirpstack.io/downloads/chirpstack/chirpstack_%CHIRPSTACK_VERSION_sqlite_linux_arm64.tar.gz)         | .tar.gz | SQLite     | Linux | arm64 |
| [chirpstack_%CHIRPSTACK_VERSION_linux_amd64.rpm](https://artifacts.chirpstack.io/downloads/chirpstack/chirpstack_%CHIRPSTACK_VERSION_linux_amd64.rpm)                             | .rpm    | PostgreSQL | Linux | amd64 |
| [chirpstack_%CHIRPSTACK_VERSION_linux_armv7.rpm](https://artifacts.chirpstack.io/downloads/chirpstack/chirpstack_%CHIRPSTACK_VERSION_linux_armv7.rpm)                             | .rpm    | PostgreSQL | Linux | armv7 |
| [chirpstack_%CHIRPSTACK_VERSION_linux_arm64.rpm](https://artifacts.chirpstack.io/downloads/chirpstack/chirpstack_%CHIRPSTACK_VERSION_linux_arm64.rpm)                             | .rpm    | PostgreSQL | Linux | arm64 |
| [chirpstack-sqlite_%CHIRPSTACK_VERSION_linux_amd64.rpm](https://artifacts.chirpstack.io/downloads/chirpstack/chirpstack-sqlite_%CHIRPSTACK_VERSION_linux_amd64.rpm)               | .rpm    | SQLite     | Linux | amd64 |
| [chirpstack-sqlite_%CHIRPSTACK_VERSION_linux_armv7.rpm](https://artifacts.chirpstack.io/downloads/chirpstack/chirpstack-sqlite_%CHIRPSTACK_VERSION_linux_armv7.rpm)               | .rpm    | SQLite     | Linux | armv7 |
| [chirpstack-sqlite_%CHIRPSTACK_VERSION_linux_arm64.rpm](https://artifacts.chirpstack.io/downloads/chirpstack/chirpstack-sqlite_%CHIRPSTACK_VERSION_linux_arm64.rpm)               | .rpm    | SQLite     | Linux | arm64 |

Previous versions are available at:
[https://artifacts.chirpstack.io/downloads/chirpstack/](https://artifacts.chirpstack.io/downloads/chirpstack/).

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

## Docker

For Docker images, please refer to:

- [https://hub.docker.com/r/chirpstack/chirpstack](https://hub.docker.com/r/chirpstack/chirpstack/tags)
  (PostgreSQL)
- [https://hub.docker.com/r/chirpstack/chirpstack-sqlite](https://hub.docker.com/r/chirpstack/chirpstack-sqlite/tags)
  (SQLite)
