# Contribute & source

## Links

- Community forum: [https://forum.chirpstack.io/](https://forum.chirpstack.io/)
- Source repository: [https://github.com/chirpstack/chirpstack](https://github.com/chirpstack/chirpstack)
  - Issue and bug reports: [https://github.com/chirpstack/chirpstack/issues](https://github.com/chirpstack/chirpstack/issues)

## Building from source

The easiest way to get started is by using the provided 
[docker-compose](https://docs.docker.com/compose/) environment. This environment
comes with all dependencies pre-installed. To start a bash shell within the
docker-compose environment, execute the following command from the root of 
this project:

```bash
make devshell
```

For a list of available make commands that can be executed, please consult
the `Makefile` that can be found in the root of each component.
