# Event types

Depending the integration, it is possible to encode events in several ways:

* JSON: JSON based on the Protocol Buffers JSON mapping
* Protobuf: Protocol Buffers binary encoding

For the [Protobuf](https://developers.google.com/protocol-buffers/)
message definitions, please refer to [integration.proto](https://github.com/chirpstack/chirpstack/blob/master/api/proto/integration/integration.proto).

<!-- toc -->

## Event types

### `up` - Uplink event

Defined by the `UplinkEvent` Protobuf message.

JSON example:


```json
{
	"deduplicationId": "3ac7e3c4-4401-4b8d-9386-a5c902f9202d",
	"time": "2022-07-18T09:34:15.775023242+00:00",
	"deviceInfo": {
		"tenantId": "52f14cd4-c6f1-4fbd-8f87-4025e1d49242",
		"tenantName": "ChirpStack",
		"applicationId": "17c82e96-be03-4f38-aef3-f83d48582d97",
		"applicationName": "Test application",
		"deviceProfileId": "14855bf7-d10d-4aee-b618-ebfcb64dc7ad",
		"deviceProfileName": "Test device-profile",
		"deviceName": "Test device",
		"devEui": "0101010101010101",
		"tags": {
			"key": "value"
		}
	},
	"devAddr": "00189440",
	"dr": 1,
	"fPort": 1,
	"data": "qg==",
	"rxInfo": [{
		"gatewayId": "0016c001f153a14c",
		"uplinkId": 4217106255,
		"rssi": -36,
		"snr": 10.5,
		"context": "E3OWOQ==",
		"metadata": {
			"region_name": "eu868",
			"region_common_name": "EU868"
		}
	}],
	"txInfo": {
		"frequency": 867100000,
		"modulation": {
			"lora": {
				"bandwidth": 125000,
				"spreadingFactor": 11,
				"codeRate": "CR_4_5"
			}
		}
	}
}
```

### `status` - Margin and battery status

Defined by the `StatusEvent` Protobuf message.

JSON example:


```json
{
	"deduplicationId": "bf2a9746-1c8b-49be-a92e-f0b89c8534e4",
	"time": "2022-07-18T09:35:06.927172638+00:00",
	"deviceInfo": {
		"tenantId": "52f14cd4-c6f1-4fbd-8f87-4025e1d49242",
		"tenantName": "ChirpStack",
		"applicationId": "17c82e96-be03-4f38-aef3-f83d48582d97",
		"applicationName": "Test application",
		"deviceProfileId": "14855bf7-d10d-4aee-b618-ebfcb64dc7ad",
		"deviceProfileName": "Test device-profile",
		"deviceName": "Test device",
		"devEui": "0101010101010101"
	},
	"margin": 10,
	"batteryLevel": 88.3
}
```

### `join` - Device join event

Defined by the `JoinEvent` Protobuf message.

JSON example:


```json
{
	"deduplicationId": "c9dbe358-2578-4fb7-b295-66b44edc45a6",
	"time": "2022-07-18T09:33:28.823500726+00:00",
	"deviceInfo": {
		"tenantId": "52f14cd4-c6f1-4fbd-8f87-4025e1d49242",
		"tenantName": "ChirpStack",
		"applicationId": "17c82e96-be03-4f38-aef3-f83d48582d97",
		"applicationName": "Test application",
		"deviceProfileId": "14855bf7-d10d-4aee-b618-ebfcb64dc7ad",
		"deviceProfileName": "Test device-profile",
		"deviceName": "Test device",
		"devEui": "0101010101010101",
		"tags": {
			"key": "value"
		}
	},
	"devAddr": "00189440"
}
```

### `ack` - Confirmed downlink (n)ack

Defined by the `AckEvent` Protobuf message.

JSON example:


```json
{
	"deduplicationId": "856a03e7-80f6-4c26-b971-9dcb34a59c2e",
	"time": "2022-07-18T09:38:49.235700766+00:00",
	"deviceInfo": {
		"tenantId": "52f14cd4-c6f1-4fbd-8f87-4025e1d49242",
		"tenantName": "ChirpStack",
		"applicationId": "17c82e96-be03-4f38-aef3-f83d48582d97",
		"applicationName": "Test application",
		"deviceProfileId": "14855bf7-d10d-4aee-b618-ebfcb64dc7ad",
		"deviceProfileName": "Test device-profile",
		"deviceName": "Test device",
		"devEui": "0101010101010101",
		"tags": {
			"key": "value"
		}
	},
	"queueItemId": "1bbcab1d-6ec6-4136-b3f5-d63be1dc08de",
	"acknowledged": true,
	"fCntDown": 3
}
```

### `txack` - Downlink transmission ack

Defined by the `TxAckEvent` Protobuf message.

JSON example:


```json
{
	"downlinkId": 3461958127,
	"time": "2022-07-18T09:37:48.929184085+00:00",
	"deviceInfo": {
		"tenantId": "52f14cd4-c6f1-4fbd-8f87-4025e1d49242",
		"tenantName": "ChirpStack",
		"applicationId": "17c82e96-be03-4f38-aef3-f83d48582d97",
		"applicationName": "Test application",
		"deviceProfileId": "14855bf7-d10d-4aee-b618-ebfcb64dc7ad",
		"deviceProfileName": "Test device-profile",
		"deviceName": "Test device",
		"devEui": "0101010101010101",
		"tags": {
			"key": "value"
		}
	},
	"queueItemId": "42cb6459-a640-46ec-9623-bdff39dc4736",
	"fCntDown": 2,
	"gatewayId": "0016c001f153a14c",
	"txInfo": {
		"frequency": 867900000,
		"power": 14,
		"modulation": {
			"lora": {
				"bandwidth": 125000,
				"spreadingFactor": 7,
				"codeRate": "CR_4_5",
				"polarizationInversion": true
			}
		},
		"timing": {
			"delay": {
				"delay": "1s"
			}
		},
		"context": "ICYVGw=="
	}
}
```

### `log` - Log (or error) event

Defined by the `LogEvent` Protobuf message.

JSON example:


```json
{
	"time": "2022-07-18T09:39:57.311414029+00:00",
	"deviceInfo": {
		"tenantId": "52f14cd4-c6f1-4fbd-8f87-4025e1d49242",
		"tenantName": "ChirpStack",
		"applicationId": "17c82e96-be03-4f38-aef3-f83d48582d97",
		"applicationName": "Test application",
		"deviceProfileId": "14855bf7-d10d-4aee-b618-ebfcb64dc7ad",
		"deviceProfileName": "Test device-profile",
		"deviceName": "Test device",
		"devEui": "0101010101010101",
		"tags": {
			"key": "value"
		}
	},
	"level": "ERROR",
	"code": "UPLINK_MIC",
	"description": "MIC of join-request is invalid, make sure keys are correct",
	"context": {
		"deduplication_id": "4e91db13-5ce3-4c8c-ab0f-dfc626cf0771"
	}
}
```

### `location` - Location event

Defined by the `LocationEvent` Protobuf message.

JSON example:

```json
{
	"deduplicationId": "63d870b9-b769-4d88-8b33-8e43f58851a7",
	"time": "2022-07-18T09:55:11.059481003+00:00",
	"deviceInfo": {
		"tenantId": "52f14cd4-c6f1-4fbd-8f87-4025e1d49242",
		"tenantName": "ChirpStack",
		"applicationId": "17c82e96-be03-4f38-aef3-f83d48582d97",
		"applicationName": "Test application",
		"deviceProfileId": "14855bf7-d10d-4aee-b618-ebfcb64dc7ad",
		"deviceProfileName": "Test device-profile",
		"deviceName": "Test device",
		"devEui": "0101010101010101",
		"tags": {
			"key": "value"
		}
	},
	"location": {
		"latitude": 37.344214,
		"longitude": -8.839375,
		"altitude": 2
	}
}
```

### `integration` - Integration event

Defined by the `IntegrationEvent` Protobuf message.

JSON example:

```json
{
	"deduplicationId": "39c68646-57f7-4973-8ac8-48839b5714d6",
	"time": "2022-07-18T10:02:15.812284501+00:00",
	"deviceInfo": {
		"tenantId": "52f14cd4-c6f1-4fbd-8f87-4025e1d49242",
		"tenantName": "ChirpStack",
		"applicationId": "17c82e96-be03-4f38-aef3-f83d48582d97",
		"applicationName": "Test application",
		"deviceProfileId": "14855bf7-d10d-4aee-b618-ebfcb64dc7ad",
		"deviceProfileName": "Test device-profile",
		"deviceName": "LR1110",
		"devEui": "0016c001f000175a",
		"tags": {
			"key": "value"
		}
	},
	"integrationName": "loracloud",
	"eventType": "modem_UplinkResponse",
	"object": {
		"stream_records": [
			[43.0, "0815a7f492bf248575caf492bf248725c874acb96799de"]
		],
		"dnlink": null,
		"pending_requests": {
			"id": 10.0,
			"requests": [],
			"upcount": 0.0,
			"updelay": 32.0
		},
		"fulfilled_requests": [],
		"fports": {
			"wifiport": 197.0,
			"gnssport": 198.0,
			"dmport": 199.0,
			"streamport": 199.0
		},
		"log_messages": [],
		"info_fields": {
			"voltage": {
				"timestamp": 1648481184.757,
				"value": 3.04
			},
			"charge": {
				"value": 0.0,
				"timestamp": 1648481184.757
			},
			"interval": null,
			"rxtime": null,
			"appstatus": null,
			"joineui": null,
			"streampar": null,
			"uptime": null,
			"firmware": {
				"value": {
					"fwtotal": 65535.0,
					"fwcrc": "c66f3ebe",
					"fwcompleted": 0.0
				},
				"timestamp": 1648481184.757
			},
			"adrmode": null,
			"deveui": null,
			"signal": null,
			"rstcount": {
				"timestamp": 1648481184.757,
				"value": 300.0
			},
			"chipeui": null,
			"crashlog": null,
			"status": null,
			"region": null,
			"session": {
				"value": 304.0,
				"timestamp": 1648481184.757
			},
			"rfu": null,
			"temp": {
				"timestamp": 1648481184.757,
				"value": 25.0
			},
			"alcsync": {
				"value": {
					"token": 0.0,
					"time": 19.0
				},
				"timestamp": 1648481177.033
			}
		},
		"position_solution": null,
		"file": null
	}
}
```
