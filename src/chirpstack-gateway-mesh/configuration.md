# Configuration

To list all CLI options, start `chirpstack-gateway-mesh` with the `--help` flag.
This will print:

```text
ChirpStack Gateway Mesh

Usage: chirpstack-gateway-mesh [OPTIONS] [COMMAND]

Commands:
  configfile  Print the configuration template
  help        Print this message or the help of the given subcommand(s)

Options:
  -c, --config <FILE>  
  -h, --help           Print help
  -V, --version        Print version
```

## Configuration example

Example `chirpstack-gateway-mesh.toml` configuration:

```toml
{{#include chirpstack-gateway-mesh.toml}}
```
