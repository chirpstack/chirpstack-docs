# Device time

ChirpStack supports the synchronization of the internal device clock with the
network using the `DeviceTimeReq` mac-command. This is useful for devices that
need to have an accurate time-source or devices implementing LoRaWAN<sup>&reg;</sup> Class-B.

If possible, ChirpStack uses the RX timestamp provided by the gateway which
results in the most accurate time. When this timestamp is not available (e.g. the
gateway is not time synchronized), it will use the current server time. Please
note that in this case the returned timestamp is less accurate as ChirpStack
is not aware of the latency between the gateway and the network-server.
