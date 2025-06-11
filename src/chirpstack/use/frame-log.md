# Frame log

For debugging purposes, ChirpStack keeps a log of the last LoRaWAN frames that
are exposed through the available integrations. This feature is available both
for gateways and devices.

### Gateway frame logs

The **LoRaWAN frames** tab on the gateway detail page will display all frames
sent and received by the selected gateway. This includes frames received from
unknown devices.

### Device frame logs

The **LoRaWAN frames** tab on the device detail page will display all frames
which have been authenticated with the (session)keys of the device. This means
that in case your device is experiencing MIC calculation issues, you will see
the uplink frame under the gateway frame logs, but not under the device frame
logs.
