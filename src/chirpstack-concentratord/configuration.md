# Configuration

ChirpStack Concentratord provides different binaries in order to target different
hardware platforms.

ChirpStack Concentratord makes it possible to use multiple configuration files,
by repeating the `-c` argument. For example:

* `concentratord.toml` for generic configuration
* `region.toml` for region specific settings (e.g. Class-B beacon configuration)
* `channels.toml` to configure the channels

When using ChirpStack Concentratord within a region supporting multiple
sub-bands, this means that only the `channels.toml` file needs to be updated
when changing the sub-band, avoiding duplication.

<!-- toc -->

## `chirpstack-concentratord-sx1301`

The `chirpstack-concentratord-sx1301` binary has the following command-line
flags:

```text
concentratord {{ concentratord_version }}
Orne Brocaar <info@brocaar.com>
LoRa concentrator HAL daemon (v1)

USAGE:
    chirpstack-concentratord-sx1301 [OPTIONS] [SUBCOMMAND]

FLAGS:
    -h, --help       Prints help information
    -V, --version    Prints version information

OPTIONS:
    -c, --config <FILE>...    Path to configuration file

SUBCOMMANDS:
    configfile    Print the configuration template
    help          Prints this message or the help of the given subcommand(s)
```

### Configuration example

```toml
{%raw%}
{{#include chirpstack-concentratord-sx1301.toml}}
{%endraw%}
```

## `chirpstack-concentratord-sx1302`

The `chirpstack-concentratord-sx1302` binary has the following command-line
flags:

```text
chirpstack-concentratord-sx1302 {{ concentratord_version }}
Orne Brocaar <info@brocaar.com>
LoRa concentrator HAL daemon for SX1302

USAGE:
    chirpstack-concentratord-sx1302 [OPTIONS] [SUBCOMMAND]

FLAGS:
    -h, --help       Prints help information
    -V, --version    Prints version information

OPTIONS:
    -c, --config <FILE>...    Path to configuration file

SUBCOMMANDS:
    configfile    Print the configuration template
    help          Prints this message or the help of the given subcommand(s)
```

### Configuration example

Unlike the `-sx1301` binary, there is no option for configuring the
`gateway_id`. The unique gateway ID is embedded in the SX1302 and read
by the Concentratord daemon.

```toml
{%raw%}
{{#include chirpstack-concentratord-sx1302.toml}}
{%endraw%}
```

## `chirpstack-concentratord-2g4`

The `chirpstack-concentratord-2g4` binary has the following command-line
flags:

```text
concentratord {{ concentratord_version }}
Orne Brocaar <info@brocaar.com>
LoRa concentrator HAL daemon (2.4GHz)

USAGE:
    chirpstack-concentratord-2g4 [OPTIONS] [SUBCOMMAND]

FLAGS:
    -h, --help       Prints help information
    -V, --version    Prints version information

OPTIONS:
    -c, --config <FILE>...    Path to configuration file

SUBCOMMANDS:
    configfile    Print the configuration template
    help          Prints this message or the help of the given subcommand(s)
```

### Configuration example

Unlike the `-sx1301` binary, there is no option for configuring the
`gateway_id`. The unique gateway ID is embedded in the concentrator and read
by the Concentratord daemon.

```toml
{%raw%}
{{#include chirpstack-concentratord-2g4.toml}}
{%endraw%}
```
