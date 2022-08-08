# Subscribe to MQTT

If you have installed the **chirpstack-gateway-os-full** image, then a
MQTT broker ([Mosquitto](https://mosquitto.org/)) is pre-installed and
pre-configured.

## Subscribing

The MQTT broker is pre-configured with two listeners. Connections from the
same host (Raspberry Pi) no not require a client-certificate or any credentials.
If you are connecting from a different host, then you must use a
client-certificate in order to connect to the MQTT broker. This client-certificate
can be generated using the ChirpStack web-interface. Please find instructions
below.

### Same host

If connecting from the same machine (Raspberry Pi), you can use the following
command to receive data:

```bash
mosquitto_sub -h localhost -t "application/#" -v
```

### Different host

If connecting from a different host, you must obtain a client-certificate first
using the ChirpStack web-interface.

* Navigate to the application for which you want to receive data
* Click the **Integrations** tab
* Click **Get certificate** under the MQTT integration card
* Click **Generate certificate** to obtain the certificate files
* Store the three files on the machine from which you want to connect to the
  MQTT broker:
	* CA certificate: `ca.pem`
	* TLS certificate: `cert.pem`
	* TLS key: `key.pem`

After completing the above steps, you can use the following command to receive
data:

```bash
mosquitto_sub -h IPADDRESS -p 8883 --insecure --cafile ca.pem --cert cert.pem --key key.pem  -t "application/#" -v
```

**Note:** Replace `IPADDRESS` with the IP address of the Raspberry Pi. Also
note the `--insecure` flag, which is required to ignore the certificate
hostname validation.
