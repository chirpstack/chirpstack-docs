# Mosquitto TLS configuration

This guide describes how to configure [Mosquitto](https://www.mosquitto.org/)
with TLS configuration. It describes how you can:

* Generate your own CA (certificate authority)
* Generate a server-certificate
* How to setup ChirpStack such that it can use the generated CA to sign
  client-certificates for gateways and the application MQTT integration

Please note:

While this guide describes the steps for Mosquitto, with some changes these
can be applied to other MQTT brokers too.
These steps might vary slightly based on your setup.
The provided configuration are examples, based on your requirements you
might need to modify these.

 
<!-- toc -->

## Requirements

Before proceeding, please make sure that you have installed the [cfssl](https://github.com/cloudflare/cfssl)
utility. You should also already have a working ChirpStack environment.

If using Debian or Ubuntu, this package can be installed using:

```bash
sudo apt-get install golang-cfssl
```

## Generating CA

The CA certificate (and key) has two purposes:

* It is used to sign certificates
* It is used to validate certificates (checking it has been signed by the CA)

Save the following files to disk:

`ca-csr.json`:

```json
{
  "CN": "ChirpStack CA",
  "key": {
    "algo": "rsa",
    "size": 4096
  }
}
```

`ca-config.json`:

```json
{
  "signing": {
    "default": {
      "expiry": "8760h"
    },
    "profiles": {
      "server": {
        "expiry": "8760h",
        "usages": [
          "signing",
          "key encipherment",
          "server auth"
        ]
      }
    }
  }
}
```

Then execute the following command to generate the CA certificate and key:

```bash
cfssl gencert -initca ca-csr | cfssljson -bare ca
```

## Generate MQTT server-certificate

The MQTT server-sertificate is used to establish a secure TLS connection between
the MQTT client (gateway or integration) and the MQTT broker.

Save the following file to disk and change `example.com` into the hostname
that will be used by clients that will connect to the MQTT broker:

`mqtt-server.json`:

```json
{
	"CN": "example.com",
	"hosts": [
		"example.com"
	],
	"key": {
		"algo": "rsa",
		"size": 4096
	}
}
```

Then execute the following command to generate the MQTT server-certificate:

```bash
cfssl gencert -ca ca.pem -ca-key ca-key.pem -config ca-config.json -profile server mqtt-server.json | cfssljson -bare mqtt-server
```

## Configure ChirpStack

You must configure ChirpStack with the CA certificate and CA key which you have
generated in the previous step. ChirpStack will use the CA certificate (+ key)
to sign the gateway or application MQTT integration client-certificates when
generated using the web-interface.

Create a directory for the certificate files and copy these files
using the following command:

```bash
mkdir -p /etc/chirpstack/certs
cp ca.pem /etc/chirpstack/certs
cp ca-key.pem /etc/chirpstack/certs
```

Depending your setup, you might need to modify the ownership and / or permissions
of the created directory and files.

Update the following configuration sections in the `chirpstack.toml` configuration
file:

`[gateway]` section:

```toml
client_cert_lifetime="12months"
ca_cert="/etc/chirpstack/certs/ca.pem"
ca_key="/etc/chirpstack/certs/ca-key.pem"
```

`[integration.mqtt.client]` section:

```toml
client_cert_lifetime="12months"
ca_cert="/etc/chirpstack/certs/ca.pem"
ca_key="/etc/chirpstack/certs/ca-key.pem"
```

Make sure to restart ChirpStack. Also verify the logs for possible errors.

## Configure Mosquitto

Create a directory for the certificate files using the following command:

```bash
mkdir -p /etc/mosquitto/certs
cp ca.pem /etc/mosquitto/certs
cp mqtt-server.pem /etc/mosquitto/certs
cp mqtt-server-key.pem /etc/mosquitto/certs
```

Depending your setup, you might need to modify the ownership and / or permissions
of the created directory and files.

To restrict MQTT clients (gateway and integrations) to their own topics,
create the following ACL file:

`/etc/mosquitto/acl`:

```text
pattern readwrite +/gateway/%u/#
pattern readwrite application/%u/#
```

Note that the `%u` will be automatically replaced by the `CN` field of the
client-certificate. The `+` prefix in `+/gateway/%u/#` is used for the region
prefix.

The following configuration file will configure two listeners:

* One listener on port `1883` (no TLS), which is accessible by any MQTT on
  the same machine (`localhost` only).
* One listener on port `8883` (TLS), which is accessible on any network
  interface. Client must use a client-certificate in order to connect to this
  listener.

`/etc/mosquitto/conf.d/listeners.conf`:

```text
per_listener_settings true

listener 1883 127.0.0.1
allow_anonymous true

listener 8883 0.0.0.0
cafile /etc/mosquitto/certs/ca.pem
certfile /etc/mosquitto/certs/mqtt-server.pem
keyfile /etc/mosquitto/certs/mqtt-server-key.pem
allow_anonymous false
require_certificate true
use_identity_as_username true
acl_file /etc/mosquitto/acl
```

Make sure to restart ChirpStack. Also verify the logs for possible errors.

## Verify

### localhost (no TLS)

From the same machine on which Mosquitto is running, you should be able to
connect without providing any credentials, using the following command:

```bash
mosquitto_sub  -h localhost -t "#" -v -d
```

### Any host (TLS)

From any other machine, you should be able to connect to the MQTT broker using
the following command:

```bash
mosquitto_sub -h example.com -p 8883 --cafile ca.pem --cert cert.pem --key key.pem  -t "#" -v -d
```

Please note:

* The `ca.pem`, `cert.pem` and `key.pem` must be obtained from the ChirpStack
  web-interface (gateway certificate or application MQTT integration certificate).
* Verify that your firewall rules allow incoming connections to the MQTT broker.
* In case you see TLS related errors, please verify that the hostname
  (of the `-h` flag) matches the MQTT server-certificate. Validation of the 
  server-certificate can be disabled using the `--insecure` flag.
