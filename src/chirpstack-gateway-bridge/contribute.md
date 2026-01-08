# Contribute & source

## Links

- Community support:
  [https://support.chirpstack.io/](https://support.chirpstack.io/)
- Source repository:
  [https://github.com/chirpstack/chirpstack-gateway-bridge](https://github.com/chirpstack/chirpstack-gateway-bridge)
  - Issue and bug reports:
    [https://github.com/chirpstack/chirpstack-gateway-bridge/issues](https://github.com/chirpstack/chirpstack-gateway-bridge/issues)

## Building from source

### With Docker

The easiest way to get started is by using the provided
[docker-compose](https://docs.docker.com/compose/) environment. To start a bash
shell within the docker-compose environment, execute the following command from
the root of this project:

```bash
docker-compose run --rm chirpstack-gateway-bridge bash
```

### Without Docker

It is possible to build ChirpStack Gateway Bridge without Docker. However this
requires to install a couple of dependencies (depending your platform, there
might be pre-compiled packages available):

#### Go

Make sure you have [Go](https://golang.org/) installed (1.11+). As ChirpStack
Gateway Bridge, the repository must be cloned outside the `$GOPATH`.

### Example commands

A few example commands that you can run:

```bash
# install development requirements
make dev-requirements

# run the tests
make test

# compile
make build

# compile snapshot for supported architectures (using goreleaser)
make snapshot
```
