# Device metrics

ChirpStack can visualize measurements collected by devices in simple graphs.
This makes testing and setting up proof of concepts really easy, as no external
software or services are required to visualize collected data.

In order to start visualizing collected measurements, you must configure the
following:

- Payload codec: As ChirpStack must decode the data in order to collect the
  measurements.
- Measurements: The measurements that must be aggregated by ChirpStack.

Both options can be found in the [Device Profile](../use/device-profiles.md)
configuration.
