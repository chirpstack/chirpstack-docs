# Configuration

To list all CLI options, start `chirpstack-mqtt-forwarder` with the `--help`
flag. This will print:

```text
ChirpStack MQTT Forwarder

Usage: chirpstack-mqtt-forwarder [OPTIONS] [COMMAND]

Commands:
  configfile  Print the configuration template
  help        Print this message or the help of the given subcommand(s)

Options:
  -c, --config <FILE>  
  -h, --help           Print help information
  -V, --version        Print version information
```

## Configuration example

Example `chirpstack-mqtt-forwarder.toml` configuration:

```toml
{{#include chirpstack-mqtt-forwarder.toml}}
```
