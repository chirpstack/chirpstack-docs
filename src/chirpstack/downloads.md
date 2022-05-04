# Downloads

## Binaries

| File name | Type | OS | Arch |
| --------- | ---- | -- | ---- |
| [chirpstack_{{chirpstack_version}}_linux_amd64.deb](https://artifacts.chirpstack.io/downloads/chirpstack/chirpstack_{{chirpstack_version}}_linux_amd64.deb) | .deb | Linux | amd64 |
| [chirpstack_{{chirpstack_version}}_linux_arm64.deb](https://artifacts.chirpstack.io/downloads/chirpstack/chirpstack_{{chirpstack_version}}_linux_arm64.deb) | .deb | Linux | arm64 |
| [chirpstack_{{chirpstack_version}}_linux_amd64.tar.gz](https://artifacts.chirpstack.io/downloads/chirpstack/chirpstack_{{chirpstack_version}}_linux_amd64.tar.gz) | .tar.gz | Linux | amd64 |
| [chirpstack_{{chirpstack_version}}_linux_arm64.tar.gz](https://artifacts.chirpstack.io/downloads/chirpstack/chirpstack_{{chirpstack_version}}_linux_arm64.tar.gz) | .tar.gz | Linux | arm64 |
| [chirpstack_{{chirpstack_version}}_linux_amd64.rpm](https://artifacts.chirpstack.io/downloads/chirpstack/chirpstack_{{chirpstack_version}}_linux_amd64.rpm) | .rpm | Linux | amd64 |
| [chirpstack_{{chirpstack_version}}_linux_arm64.rpm](https://artifacts.chirpstack.io/downloads/chirpstack/chirpstack_{{chirpstack_version}}_linux_arm64.rpm) | .rpm | Linux | arm64 |


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

For Docker images, please refer to
[https://hub.docker.com/r/chirpstack/chirpstack](https://hub.docker.com/r/chirpstack/chirpstack/tags)
