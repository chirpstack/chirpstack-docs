# Node-RED integration

This guide explains how to setup [Node-RED](https://nodered.org/) with the
[ChirpStack Node-RED nodes](https://github.com/brocaar/node-red-contrib-chirpstack/)
and setup a simple echo flow. This flow will consume uplink event using MQTT
and will enqueue downlink messages using the downlink node, which uses the
[gRPC API](../chirpstack/api/grpc.md).

Before you start, please make sure that you have a working ChirpStack and
Node-RED and that you have already a device connected. As we will setup an
echo flow, the best device to use would be a device of which you can read the
received downlink messages using a serial interface.

## ChirpStack Node-RED package

ChirpStack provides a package called [node-red-contrib-chirpstack](https://github.com/brocaar/node-red-contrib-chirpstack/)
which provides several Node-RED nodes that will help you to integrate with
ChirpStack.

Please refer to the [Adding nodes to the palette](https://nodered.org/docs/user-guide/runtime/adding-nodes)
documentation for installation instructions. Depending your Node-RED
installation, the command you need to execute could be:

```bash
npm install @chirpstack/node-red-contrib-chirpstack
```

**Note:** If you are using the Node-RED instance of the [ChirpStack Gateway OS](../chirpstack-gateway-os/index.md),
then you can skip this step as this package is already installed for you.

## Echo flow

![echo flow](node-red-echo-flow.png)

Below steps describe how to build an echo flow, which will enqueue the received
uplink from the device as downlink. While this probably doesn't cover a real
use-case, it does demonstrate how to receive uplink messages, process these
and schedule downlinks.

#### Receive from MQTT

After opening Node-RED in your browser, first setup a **mqtt in** node. This
node will be used to receive device events that are published by the MQTT
integration of ChirpStack.

The following node properties must be set:

* Server:
	* Server: e.g. `localhost`
	* Port: `1883` or `8883` in case TLS is used
	* Select *Use TLS* in case TLS is used, in which case also add a new TLS config
* Topic: `application/+/device/+/event/+`
* QoS: `0`
* Output: *auto-detect*

Note:  	

To test the connection to the MQTT broker, **Deploy** the flow. Under the newly
added node you should see a label *Connected*.

#### Parse events

Now that you are able to receive events from MQTT, setup a **device event**
node (you will find this node under the ChirpStack section). This node parses
the payloads generated by the **mqtt in** node to objects and it will perform
filtering based on event types.

The following node properties must be set:

* Event Type: *Uplink*

Connect the output side of the **mqtt in** node with the input side of the
**device event** node.

#### Debug events

For debugging, add a **debug** node and connect it to the **device event**
output and re-**Deploy** the flow. Open the debug messages tab in the Node-RED
interface. When one of your devices sends an uplink, you should see the uplink
event as debug message.

#### Echo payload function

To model the uplink event as downlink enqueue payload, add a **function** node
and connect the input to the output of the **device event** node.

The following node properties must be set:

**On Message**

```js
return {
    devEui: msg.payload.deviceInfo.devEui,
    fPort: msg.payload.fPort,
    confirmed: false,
    payload: msg.payload.data
};
```

This will make the **function** node output a message that can be used with the
**device downlink** node, which will be added in the next step.

#### Enqueue downlink

**Note:** For this step you need to obtain an API token from the ChirpStack
web-interface first!

To send the downlink message generated by the echo function to ChirpStack, add
a **device downlink** node and connect the input to the output of the
**function** node.

The following node properties must be set:

* Server: `localhost:8080`
* API token: *the API token obtained from the ChirpStack web-interface*
* Payload Encoding: *Base64*

#### Enqueue output (debug)

To view the message generated by the **device downlink** node, add an other
**debug** node and connect the input to the output of the **device downlink**
node.

The following node properties must be set:

* Output: *complete msg object*

#### Testing the flow

Make sure to re-**Deploy** the flow, so that all changes are deployed. When
sending an uplink frame, you should see two messages in the debug panel:

* The uplink event
* The response from the downlink enqueue

An example of the downlink enqueue output:

```
{"id": "c0768c9a-3ff3-4855-846c-7fdfa8894666", "_msgid": "a8264ce66a2c4830"}
```

**Note:** it is possible that the device receives the downlink after the next
uplink message.
