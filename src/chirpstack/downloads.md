# Downloads

## Binaries

| File name | Type | DB | OS | Arch |
| --------- | ---- | -- | -- | ---- |
| [chirpstack_%CHIRPSTACK_VERSION_linux_amd64.deb](https://artifacts.chirpstack.io/downloads/chirpstack/chirpstack_%CHIRPSTACK_VERSION_linux_amd64.deb) | .deb | PostgreSQL | Linux | amd64 |
| [chirpstack_%CHIRPSTACK_VERSION_linux_armhf.deb](https://artifacts.chirpstack.io/downloads/chirpstack/chirpstack_%CHIRPSTACK_VERSION_linux_armhf.deb) | .deb | PostgreSQL | Linux | armv7 |
| [chirpstack_%CHIRPSTACK_VERSION_linux_arm64.deb](https://artifacts.chirpstack.io/downloads/chirpstack/chirpstack_%CHIRPSTACK_VERSION_linux_arm64.deb) | .deb | PostgreSQL | Linux | arm64 |
| [chirpstack-sqlite_%CHIRPSTACK_VERSION_linux_amd64.deb](https://artifacts.chirpstack.io/downloads/chirpstack/chirpstack-sqlite_%CHIRPSTACK_VERSION_linux_amd64.deb) | .deb | SQLite | Linux | amd64 |
| [chirpstack-sqlite_%CHIRPSTACK_VERSION_linux_armhf.deb](https://artifacts.chirpstack.io/downloads/chirpstack/chirpstack-sqlite_%CHIRPSTACK_VERSION_linux_armhf.deb) | .deb | SQLite | Linux | armv7 |
| [chirpstack-sqlite_%CHIRPSTACK_VERSION_linux_arm64.deb](https://artifacts.chirpstack.io/downloads/chirpstack/chirpstack-sqlite_%CHIRPSTACK_VERSION_linux_arm64.deb) | .deb | SQLite | Linux | arm64 |
| [chirpstack_%CHIRPSTACK_VERSION_postgres_linux_amd64.tar.gz](https://artifacts.chirpstack.io/downloads/chirpstack/chirpstack_%CHIRPSTACK_VERSION_postgres_linux_amd64.tar.gz) | .tar.gz | PostgreSQL | Linux | amd64 |
| [chirpstack_%CHIRPSTACK_VERSION_postgres_linux_armv7hf.tar.gz](https://artifacts.chirpstack.io/downloads/chirpstack/chirpstack_%CHIRPSTACK_VERSION_postgres_linux_armv7hf.tar.gz) | .tar.gz | PostgreSQL | Linux | armv7 |
| [chirpstack_%CHIRPSTACK_VERSION_postgres_linux_arm64.tar.gz](https://artifacts.chirpstack.io/downloads/chirpstack/chirpstack_%CHIRPSTACK_VERSION_postgres_linux_arm64.tar.gz) | .tar.gz | PostgreSQL | Linux | arm64 |
| [chirpstack_%CHIRPSTACK_VERSION_sqlite_linux_amd64.tar.gz](https://artifacts.chirpstack.io/downloads/chirpstack/chirpstack_%CHIRPSTACK_VERSION_sqlite_linux_amd64.tar.gz) | .tar.gz | SQLite | Linux | amd64 |
| [chirpstack_%CHIRPSTACK_VERSION_sqlite_linux_armv7hf.tar.gz](https://artifacts.chirpstack.io/downloads/chirpstack/chirpstack_%CHIRPSTACK_VERSION_sqlite_linux_armv7hf.tar.gz) | .tar.gz | SQLite | Linux | armv7 |
| [chirpstack_%CHIRPSTACK_VERSION_sqlite_linux_arm64.tar.gz](https://artifacts.chirpstack.io/downloads/chirpstack/chirpstack_%CHIRPSTACK_VERSION_sqlite_linux_arm64.tar.gz) | .tar.gz | SQLite | Linux | arm64 |
| [chirpstack_%CHIRPSTACK_VERSION_linux_amd64.rpm](https://artifacts.chirpstack.io/downloads/chirpstack/chirpstack_%CHIRPSTACK_VERSION_linux_amd64.rpm) | .rpm | PostgreSQL | Linux | amd64 |
| [chirpstack_%CHIRPSTACK_VERSION_linux_armv7.rpm](https://artifacts.chirpstack.io/downloads/chirpstack/chirpstack_%CHIRPSTACK_VERSION_linux_armv7.rpm) | .rpm | PostgreSQL | Linux | armv7 |
| [chirpstack_%CHIRPSTACK_VERSION_linux_arm64.rpm](https://artifacts.chirpstack.io/downloads/chirpstack/chirpstack_%CHIRPSTACK_VERSION_linux_arm64.rpm) | .rpm | PostgreSQL | Linux | arm64 |
| [chirpstack-sqlite_%CHIRPSTACK_VERSION_linux_amd64.rpm](https://artifacts.chirpstack.io/downloads/chirpstack/chirpstack-sqlite_%CHIRPSTACK_VERSION_linux_amd64.rpm) | .rpm | SQLite | Linux | amd64 |
| [chirpstack-sqlite_%CHIRPSTACK_VERSION_linux_armv7.rpm](https://artifacts.chirpstack.io/downloads/chirpstack/chirpstack-sqlite_%CHIRPSTACK_VERSION_linux_armv7.rpm) | .rpm | SQLite | Linux | armv7 |
| [chirpstack-sqlite_%CHIRPSTACK_VERSION_linux_arm64.rpm](https://artifacts.chirpstack.io/downloads/chirpstack/chirpstack-sqlite_%CHIRPSTACK_VERSION_linux_arm64.rpm) | .rpm | SQLite | Linux | arm64 |


## Debian / Ubuntu repository

ChirpStack provides pre-compiled binaries as Debian (`.deb`) packages. To
activate this repository, execute the following commands:

```bash
sudo apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 1CE2AFD36DBCCA00

sudo echo "deb https://artifacts.chirpstack.io/packages/4.x/deb stable main" | sudo tee /etc/apt/sources.list.d/chirpstack_4.list
sudo apt update
```

ChirpStack also provides a repository for test-builds. To activate the test
channel use the following commands (notice that `stable` has changed into `testing`):

```bash
sudo apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 1CE2AFD36DBCCA00

sudo echo "deb https://artifacts.chirpstack.io/packages/4.x/deb testing main" | sudo tee /etc/apt/sources.list.d/chirpstack_4.list
sudo apt update
```

## Docker

For Docker images, please refer to:

* [https://hub.docker.com/r/chirpstack/chirpstack](https://hub.docker.com/r/chirpstack/chirpstack/tags) (PostgreSQL)
* [https://hub.docker.com/r/chirpstack/chirpstack-sqlite](https://hub.docker.com/r/chirpstack/chirpstack-sqlite/tags) (SQLite)
