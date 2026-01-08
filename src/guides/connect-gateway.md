# Connecting a Gateway

This guide describes how to connect your gateway to ChirpStack, and how to
validate that it is successfully communicating with ChirpStack.

**Note:** This guide does not cover the configuration of MQTT credentials and /
or (client) certificates.

## Requirements

### ChirpStack

It is expected that at this point, ChirpStack is up-and-running.

### ChirpStack Concentratord / Packet Forwarder

Please make sure that your gateway either has the
[ChirpStack Concentratord](../chirpstack-concentratord/index.md),
[Semtech UDP Packet Forwarder](https://github.com/lora-net/packet_forwarder) or
the [LoRa Basics Station](https://doc.sm.tc/station/) installed.

### ChirpStack MQTT Forwarder

The steps needed to install the ChirpStack MQTT Forwarder vary per gateway
vendor and model. Please see
[ChirpStack MQTT Forwarder installation](../chirpstack-mqtt-forwarder/install/index.md)
for more information.

In the case it is not possible to install the ChirpStack MQTT Forwarder on the
gateway, you must install the
[ChirpStack Gateway Bridge](../chirpstack-gateway-bridge/index.md) on a server
and connect your gateway to it.

## Configuration

There are two components that you need to configure; the ChirpStack
Concentratord or packet-forwarder and the ChirpStack MQTT Forwarder or
ChirpStack Gateway Bridge. Please note that this section only covers a summary.
For more information, please refer to the documentation of each component.

### ChirpStack Concentratord / Packet Forwarder

#### ChirpStack Concentratord

ChirpStack provides pre-configured packages for various gateway models. No
further configuration is required. Please see
[ChirpStack Concentratord installation](../chirpstack-concentratord/installation/index.md)
for more information.

#### Packet Forwarder

In case the gateway comes with a Packet Forwarder pre-installed and you wish to
use it instead of the ChirpStack Concentratord, then there are several
scenarios:

##### Semtech UDP Forwarder

In case it is possible to install the
[ChirpStack MQTT Forwarder](../chirpstack-mqtt-forwarder/index.md) on the
gateway, you must configure the Semtech UDP Packet Forwarder to forward its data
to `localhost:1700`. In case it is **not** possible to install the ChirpStack
MQTT Forwarder on the gateway, then you must point it to the hostname / IP of
your [ChirpStack Gateway Bridge](../chirpstack-gateway-bridge/index.md) instance
running in the cloud.

##### LoRa Basics Station

In this case you must configure the LoRa Basics Station to connect to the
hostname / IP of your
[ChirpStack Gateway Bridge](../chirpstack-gateway-bridge/index.md) instance
running in the cloud.

### ChirpStack MQTT Forwarder / ChirpStack Gateway Bridge

#### ChirpStack MQTT Forwarder

There are two important bits to configure:

- You must configure the correct backend which is either the Semtech UDP Packet
  Forwarder or the ChirpStack Concentratord.
- You must configure it to connect to a MQTT broker. Please make sure to use the
  correct MQTT topic prefix, matching the region-configuration of your gateway!

Please refer to
[ChirpStack MQTT Forwarder configuration](../chirpstack-mqtt-forwarder/configuration.md)
for more information.

#### ChirpStack Gateway Bridge

In case it is not possible to install the ChirpStack MQTT Forwarder on the
gateway, then you must install the
[ChirpStack Gateway Bridge](../chirpstack-gateway-bridge/index.md) in the cloud.

There are two important bits to configure:

- You must configure the coorect backend which is either the Semtech UDP Packet
  Forwarder or the LoRa Basics Station.
- You must configure it to connect to a MQTT broker. If all components are
  running on the same machine, then the default configuration will most likely
  to be sufficient.

Please refer to the
[ChirpStack Gateway Bridge configuration](../chirpstack-gateway-bridge/configuration.md)
for more information.

## Adding the gateway to ChirpStack

Log in to the ChirpStack web-interface. The default credentials are:

- Username: admin
- Password: admin

### Add gateway

Navigate to **Gateways** in the web-interface, and click **Add** and complete
the form. Make sure that the **Gateway ID** field is equal to the Gateway ID of
your gateway. If this value is incorrectly configured, data received by your
gateway will be rejected.

## Validate

There are a few ways to validate if your gateway is correctly configured.

### Last seen at

Even when no LoRa(WAN) data is received by the gateway, it will send gateway
statistics periodically. Usually this stats interval is configured to 30
seconds. As ChirpStack will update the gateway **Last seen at** timestamp when
it receives statistics, this is the easiest way to validate that the gateway is
correctly configured. **Note:** it might take a short while before statistics
are sent by your gateway. You must refresh the page in order to see the (new)
**Last seen at** value.

### LoRaWAN frames

After opening the overview page of your gateway, you will see a **LoRaWAN
frames** tab. This will show all LoRaWAN frames that are received and sent by
your gateway. In case of received frames, this means that you will also see
received frames from devices that are not yours and / or that are not yet
configured. Therefore this screen is useful to validate if your gateway is able
to receive LoRaWAN frames and forward these to ChirpStack.

### MQTT

As the ChirpStack Gateway Bridge uses MQTT for communication with ChirpStack it
is also possible to subscribe to the gateway MQTT topics. This will validate
that data from your gateway arrives at least up to your MQTT broker.

An example command to subscribe to the gateway MQTT topic using `mosquitto_sub`:

```bash
mosquitto_sub -v -t "+/gateway/#"
```

If you see data like below, then this does not indicate that there is an issue.
The ChirpStack Gateway Bridge can publish (and receive) events either as JSON or
[Protobuf](https://developers.google.com/protocol-buffers), in which case the
payloads are binary encoded. This option can be configured in the ChirpStack
Gateway Bridge [Configuration](../chirpstack-gateway-bridge/configuration.md).

```text
eu868/gateway/00800000a00016b6/event/up 
��b��鄞▒       }
                ▒4/5▒X
���
  ѣ�����_▒
�$~ �    䨡����^(���������1"@8@zv��d���\�A3Al��
```

## Troubleshooting

In case no data is received by ChirpStack, follow the steps below to
troubleshoot.

### Packet forwarder

The first component to troubleshoot is the packet-forwarder on the gateway. By
sending an uplink message from a LoRaWAN device (it does not have to be
configured in ChirpStack) and checking the packet-forwarder logs, you can verify
if the gateway is able to receive the uplinks. Not receiving any uplink messages
could indicate a mis-configuration of the gateway.

Depending the gateway model, the commands to retrieve the packet-forwarder logs
can vary. Please refer to the ChirpStack Gateway Bridge
[Installation](../chirpstack-gateway-bridge/install/index.md) for model specific
instructions.

#### ChirpStack Concentratord

##### Received uplink

On receiving an uplink message, it is expected to see a log item similar to:

```text
Frame received, uplink_id: 648c85f9-19e8-4186-af96-f5a45eae7ba5, count_us: 105814115, freq: 868500000, bw: 125000, mod: LoRa, dr: SF8
```

#### Semtech UDP Packet Forwarder

##### PULL_ACK

it is expected to periodically see log items similar to the following:

```text
INFO: [down] PULL_ACK received in 1 ms
```

This indicates that the UDP data sent by the packet-forwarder to the ChirpStack
Gateway Bridge is acknowledged.

##### Received uplink

On receiving an uplink message, it is expected to see a log item similar to:

```text
JSON up: {"rxpk":[{"tmst":661201908,"time":"2020-10-06T13:25:02.523074Z","tmms":1286025921523,"chan":1,"rfch":1,"freq":868.300000,"stat":1,"modu":"LORA","datr":"SF11BW125","codr":"4/5","lsnr":8.0,"rssi":-42,"size":23,"data":"AAEAAAAAAAAAAQEBAQEBAQGheWrMbQo="}]}
```

##### tcpdump

An other way to troubleshoot the functioning of the Semtech UDP Packet Forwarder
is to use `tcpdump`.

If the ChirpStack Gateway Bridge is installed on the gateway itself, then the
following command must be executed on the gateway itself:

```bash
sudo tcpdump -AUq -i lo port 1700
```

If the ChirpStack Gateway Bridge is installed on a server, then you must execute
the following command (either on the gateway or on the receiving server):

```bash
sudo tcpdump -AUq port 1700
```

As the above command can be either executed on the gateway or on the server,
this allows you to validate not only if the gateway is forwarding data, but also
if the server is able to receive the data. If the gateway is correctly
forwarding but the server is not receiving, then this could indicate a network
related issue (e.g. firewall).

Output similar to the following is expected:

```text
11:42:00.114726 IP localhost.34268 > localhost.1700: UDP, length 12
E..(..@.@."................'.....UZ.....
11:42:00.130292 IP localhost.1700 > localhost.34268: UDP, length 4
E.. ..@.@.".....................
11:42:10.204723 IP localhost.34268 > localhost.1700: UDP, length 12
E..(.&@.@..................'.x...UZ.....
11:42:10.206503 IP localhost.1700 > localhost.34268: UDP, length 4
E.. .'@.@....................x..
11:42:10.968420 IP localhost.43827 > localhost.1700: UDP, length 113
E....h@.@............3...y.......UZ.....{"stat":{"time":"2017-09-11 11:42:10 GMT","rxnb":0,"rxok":0,"rxfw":0,"ackr":100.0,"dwnb":0,"txnb":0}}
11:42:10.970702 IP localhost.1700 > localhost.43827: UDP, length 4
E.. .i@.@..b...........3........
11:42:20.284752 IP localhost.34268 > localhost.1700: UDP, length 12
E..(..@.@..................'.....UZ.....
11:42:20.289256 IP localhost.1700 > localhost.34268: UDP, length 4
E.. ..@.@.......................
11:42:30.364780 IP localhost.34268 > localhost.1700: UDP, length 12
E..( .@.@..................'.S7..UZ.....
11:42:30.366310 IP localhost.1700 > localhost.34268: UDP, length 4
E..  .@.@....................S7.
```

Explanation:

- `localhost.34268 > localhost.1700`: packet sent from the packet-forwarder to
  the ChirpStack Gateway Bridge
- `localhost.1700 > localhost.34268`: packet sent from the ChirpStack Gateway
  Bridge to the Packet Forwarder

#### Semtech BasicStation

##### Received uplink

On receiving an uplink message, it is expected to see a log item similar to:

```text
RX 867.9MHz DR1 SF11/BW125 snr=10.8 rssi=-55 xtime=0x91000001A21C4B - updf mhdr=40 DevAddr=00D35EC3 FCtrl=00 FCnt=0 FOpts=[] 01A5 mic=643170495 (14 bytes)
```

### ChirpStack Gateway Bridge

When you have verified that the packet-forwarder is receiving LoRa(WAN) frames
sent by your device(s), you can validate that it is communicating succesfully
with the ChirpStack Gateway Bridge. You can do this by inspecting the ChirpStack
Gateway Bridge logs.

#### Gateway detected

When ChirpStack Gateway Bridge detects data from a gateway, it will subscribe to
a MQTT topic containing the gateway ID. You should see a log message similar to:

```text
integration/mqtt: subscribing to topic        qos=0 topic="eu868/gateway/00800000a00016b6/command/#"
```

Note that this log message is only printed once. If you do not see this log
message, you need to confirm that:

1. The packet-forwarder is correctly configured to forward its data to your
   ChirpStack Gateway Bridge instance.
2. In case the packet-forwarder and the ChirpStack Gateway Bridge are not
   installed on the same device, that there are no firewalls blocking the
   communication between the gateway and the server.
3. The ChirpStack Gateway Bridge is configured with the correct backend
   configuration for the used packet-forwarder. Please refer to the ChirpStack
   Gateway Bridge [Configuration](../chirpstack-gateway-bridge/configuration.md)
   documentation.

#### Uplink received

When ChirpStack Gateway Bridge receives an uplink message from the gateway, it
will print a message similar to:

```text
publishing event            event=up qos=0 topic=eu868/gateway/00800000a00016b6/event/up uplink_id=c2f855a9-7f74-4093-98ad-7d2b06a79398
```

### MQTT broker

If you have verified the above components, you should be able to see messages
being published to your MQTT broker. For this, please repeat the MQTT validate
step.

In case the ChirpStack Gateway Bridge does log that it is publishing events, but
you are unable to receive data from the MQTT broker, then verify if you
configured ACLs that prevent the ChirpStack Gateway Bridge from publishing
events and / or prevent your MQTT client from receiving events.

### ChirpStack

If you have verified that (uplink) events are published to your MQTT broker, you
can verify that these are received by ChirpStack. When ChirpStack receives an
uplink frame from one of your gateways through the MQTT broker, it will print a
log item similar to:

```
gateway/mqtt: uplink frame received           gateway_id=00800000a00016b6 uplink_id=be864b60-1392-47a1-98ef-8cd1115b6f05
```

If you do not see similar log items, then validate:

1. That in case you have configured MQTT ACLs, ChirpStack has access to the
   configured MQTT topics.
2. That in case you have modified the MQTT topic templates in the ChirpStack
   Gateway Bridge [Configuration](../chirpstack-gateway-bridge/configuration.md)
   and / or ChirpStack region [Configuration](../chirpstack/configuration.md),
   this configuration is aligned.
