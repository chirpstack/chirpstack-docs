# Event types

Depending the integration, it is possible to encode events in several ways:

* JSON: JSON based on the Protocol Buffers JSON mapping
* Protobuf: Protocol Buffers binary encoding

For the [Protobuf](https://developers.google.com/protocol-buffers/)
message definitions, please refer to [integration.proto](https://github.com/chirpstack/chirpstack/blob/master/api/proto/integration/integration.proto).

## Types

* **up**: Uplink event
* **status**: Margin and battery status
* **join**: Device join
* **ack**: Acknowledgement report of confirmed downlink
* **txack**: Downlink transmission acknowledgement
* **log**: Log (or error) event

### up

Defined by the `UplinkEvent` Protobuf message.

JSON example:


```json
```

### status

Defined by the `StatusEvent` Protobuf message.

JSON example:


```json
```

### join

Defined by the `JoinEvent` Protobuf message.

JSON example:


```json
```

### ack

Defined by the `AckEvent` Protobuf message.

JSON example:


```json
```

### txack

Defined by the `TxAckEvent` Protobuf message.

JSON example:


```json
```

### log

Defined by the `LogEvent` Protobuf message.

JSON example:


```json
```
