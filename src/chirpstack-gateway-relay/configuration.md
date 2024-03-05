# Configuration

To list all CLI options, start `chirpstack-gateway-relay` with the `--help`
flag. This will print:

```text
ChirpStack Gateway Relay

Usage: chirpstack-gateway-relay [OPTIONS] [COMMAND]

Commands:
  configfile  Print the configuration template
  help        Print this message or the help of the given subcommand(s)

Options:
  -c, --config <FILE>  
  -h, --help           Print help
  -V, --version        Print version
```

## Configuration example

Example `chirpstack-gateway-relay.toml` configuration:

```toml
{%raw-%}
{{#include chirpstack-gateway-relay.toml}}
{%endraw-%}
```