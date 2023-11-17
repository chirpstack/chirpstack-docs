# Browan

<!-- toc -->

## MiniHub Pro

The MiniHub Pro gateway supports:

* Semtech UDP Packet Forwarder
* Semtech Basics Station

### Semtech UDP Packet Forwarder

**Note:** You must install the [ChirpStack Gateway Bridge](../chirpstack-gateway-bridge/index.md)
on your server first and configure the `semtech_udp` backend.

Steps:

* In the MiniHub Pro web-interface, click **Configure LoRa Setting**.
* Under _Mode_, click **LoRa Packet Forwarder**.
* Under the _LoRa Packet Forwarder_ section, enter the following details:
  * **Server Address**: The hostname / IP address of your server on which the
    ChirpStack Gateway Bridge is installed, for example `example.com`.
  * **Server Uplink Port**: The ChirpStack Gateway Bridge port (default `1700`).
  * **Server Downlink Port**: The ChirpStack Gateway Bridge port (default `1700`).
* Under _Radio 0 / 1 Settings_ and _Channel Assignment_ validate that the
  channels match the channels that you have configured in ChirpStack.
* Click **Save**.

### Semtech Basics Station

**Note:** You must install the [ChirpStack Gateway Bridge](../chirpstack-gateway-bridge/index.md)
on your server first and configure the `basic_station` backend.

Steps:

* In the MiniHub Pro web-interface, click **Configure LoRa Settings**.
* Under _Mode_, click **LoRa Basics Station**.
* Under the _LoRa Basics Station_ section, enter the following details:
  * **Enable CUPS**: ChirpStack does not provide a CUPS endpoint, unless you are
    a different CUPS server, **disable** this option.
  * **LNS URI**: URL of the ChirpStack Gateway Bridge Basics Station backend.
    E.g. `ws://example.com:3000` (non-TLS) or `wss://example.com:3000` (TLS).
  * If the ChirpStack Gateway Bridge is configured with TLS or mutual-TLS:
    * **Install LNS Trust**: Upload the CA certificate used to sign the server-certificate.
    * **Install LNS CRT**: If using mutual-TLS, upload the gateway client-certificate.
    * **Install LNS Key**: If using mutual-TLS, upload the gateway private-key.
* Click **Save**.

## Pico Next Indoor Gateway

The Pico Next Indoor Gateway supports:

* Semtech Packet Forwarder
* Semtech Basics Station

### Semtech UDP Packet Forwarder

**Note:** You must install the [ChirpStack Gateway Bridge](../chirpstack-gateway-bridge/index.md)
on your server first and configure the `semtech_udp` backend.

Steps (in Pico Next Indoor Gateway web-interface):

* In the menu, click **LoRa > Mode Selection**.
* Under _Mode_, select **Packet Forwarder**.
* Select the _Region Plan_ select your region.
* Click **Apply**.
* In the menu, click **LoRa > Packet Forwarder**.
* Enter the following details in the _Gateway Info_ form:
  * **Server Address**: The hostname / IP of your server on which the
    ChirpStack Gateway Bridge is installed, for example `example.com`.
  * **Server Uplink Port**: The ChirpStack Gateway Bridge port (default `1700`).
  * **Server Downlink Port**: The ChirpStack Gateway Bridge port (default `1700`).
* Click **Apply**.

### Semtech Basics Station

**Note:** You must install the [ChirpStack Gateway Bridge](../chirpstack-gateway-bridge/index.md)
on your server first and configure the `basic_station` backend.

Steps:

* In the menu, click **LoRa > Mode Selection**.
* Under _Mode_, select **Basic Station**.
* Click **Apply**.
* In the menu, click **LoRa > Basic Station**.
* At the top, click the **Connection Configuration** tab.
* In the form, enter the following details:
  * **Basic Station Mode**: Select _LNS Mode_.
  * **Protocol**: Select _WebSocket_ if you have configured the ChirpStack
    Gateway Bridge **without** TLS certificates, select _WebSocket Secure_
    if you have configured the ChirpStack Gateway Bridge with a TLS
    server-certificate (and optionally mutual-TLS).
  * **Server Address:** The hostname / IP of your server on which the
    ChirpStack Gateway Bridge is installed, for example `example.com`.
  * **Server Port**: The ChirpStack Gateway Bridge port (default `3001`).
  * If the ChirpStack Gateway Bridge is configured with TLS or mutual-TLS:
    * **Trust**: Upload the CA certificate used to sign the server-certificate.
    * **CRT**: If using mutual-TLS, upload the gateway client-certificate.
    * **Key**: If using mutual-TLS, upload the gateway private-key.
* Click **Apply**.
