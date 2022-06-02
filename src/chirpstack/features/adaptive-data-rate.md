# Adaptive data-rate (ADR)

Adaptive data-rate lets ChirpStack control the data-rate and
tx-power of the device, so that it uses less airtime and less power to
transmit the same amount of data. This is not only beneficial for the
energy consumption of the device, but also optimizes the spectrum.

**Note:** ADR should only be used for static devices (devices that
do not move).

## Activating ADR

The activation of ADR is controlled by the device. Only when the device
sends an uplink frame with the ADR flag set to `true` will ChirpStack 
adjust the data-rate and tx-power of the device if needed.

## Configuration

To make sure there is enough link margin left after setting the ideal
data-rate and tx-power, it is important to configure the installation margin
correctly. A higher installation margin makes the ADR algorithm more
conservative.

## Custom ADR algorithm

ChirpStack provides several ADR algorithms out-of-the-box that should cater
most use-cases. However, it is possible to configure additional ADR algorithms.
ADR algorithms must be implemented in JavaScript, and the path(s) to these
file(s) must be configured in the ChirpStack [Configuration](../configuration.md)
file.

### Example skeleton

```js
// This must return the name of the ADR algorithm.
export function name() {
  return "Example plugin";
}

// This must return the id of the ADR algorithm.
export function id() {
  return "example_id";
}

// This handles the ADR request.
//
// Input object example:
// {
//  regionName: "eu868",
//  regionCommonName: "EU868",
//  devEui: "0102030405060708",
//  macVersion: "1.0.3",
//  regParamsRevision: "A",
//  adr: true,
//  dr: 1,
//  txPowerIndex: 0,
//  nbTrans: 1,
//  maxTxPowerIndex: 15,
//  requiredSnrForDr: -17.5,
//  installationMargin: 10,
//  minDr: 0,
//  maxDr: 5,
//  uplinkHistory: [
//    {
//      "fCnt": 10,
//      "maxSnr": 7.5,
//      "maxRssi": -110,
//      "txPowerIndex": 0,
//      "gatewayCount": 3
//    }
//  ]
// }
//
// This function must return an object, example:
// {
//  dr: 2,
//  txPowerIndex: 1,
//  nbTrans: 1
// }
export function handle(req) {
  return {
    dr: req.dr,
    txPowerIndex: req.txPowerIndex,
    nbTrans: req.nbTrans
  };
}
```
