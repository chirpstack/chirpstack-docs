# Contribute & source

## Links

- Community forum: [https://forum.chirpstack.io/](https://forum.chirpstack.io/)
- Source repository: [https://github.com/chirpstack/chirpstack](https://github.com/chirpstack/chirpstack)
  - Issue and bug reports: [https://github.com/chirpstack/chirpstack/issues](https://github.com/chirpstack/chirpstack/issues)

## Building from source

The easiest way to get started is by using the provided 
[docker-compose](https://docs.docker.com/compose/) environment. This environment
comes with all dependencies pre-installed.

To compile the ChirpStack UI, execute the following command from the root of
this project:

```bash
make build-ui
```

To start a bash shell for development, execute the following command from the
root of this project:

```bash
make devshell
```

**Note:** If you get an `message: #[derive(RustEmbed)] folder '/chirpstack/chirpstack/../ui/build' does not exist. cwd: '/chirpstack'`
error, then this means the UI hasn't been build yet (see above command).

For a list of available make commands that can be executed, please consult
the `Makefile` that can be found in the root of each component.
