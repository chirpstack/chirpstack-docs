# Device status

ChirpStack supports periodically requesting the Device Status, using the
`DevStatusReq` mac-command specified by the LoRaWAN<sup>&reg;</sup> protocol.
A device will respond to this request by sending its battery status
(if available) and its demodulation signal-to-noise ratio in dB
for the last successfully received request.

When the device status is available, it will be exposed through the configured
integrations.

## Exposed information

* **Margin**: the demodulation signal-to-noise ratio in dB rounded to the nearest
  integer value for the last successfully received DevStatusReq command.
* **External power source**: if the device was powered by an external power-source.
* **Battery level unavailable**: if this is set to true, no battery level
  is available.
* **Battery level**: battery level as percentage.

