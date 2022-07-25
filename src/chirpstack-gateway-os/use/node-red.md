# Node-RED

[Node-RED](https://nodered.org/) provides a browser-based flow editor that makes
it easy to wire together flows using the wide range of nodes in the palette.

If you have installed the **chirpstack-gateway-os-full** image, then Node-RED
comes pre-installed with ChirpStack Gateway OS since v3.5.0, including the
[node-red-contrib-chirpstack](https://github.com/brocaar/node-red-contrib-chirpstack/)
package which provides several nodes to interact with ChirpStack.

Links:

* [Node-RED integration guide](../../guides/node-red-integration.md)

## Enabling Node-RED

By default, Node-RED is disabled and must be enabled through the `gateway-config`
utility. In the root menu, select **Enable / disable applications** and then
**Enable Node-RED**. Note that once enabled, the option will change to
**Disable Node-RED**. Starting Node-RED will take some time (~ one minute).

## Access Node-RED

Once started, Node-RED is accessible in your browser on port
`http://[IP ADDRESS]:1880`. This means that if the IP address of your gateway
is `192.168.0.1`, that you can access Node-RED at `http://192.168.0.1:1880`.

