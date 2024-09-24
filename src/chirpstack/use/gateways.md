# Gateways

A tenant is able to manage its own set of gateways. Please note that
this feature might be unavailable when the tenant is configured without
gateway support.

That a gateway belongs to a given tenant does not mean that the usage 
of a gateway is limited to the tenant. Every device in the whole network
will be able to communicate using the gateway. The tenant will be
responsible however for managing the gateway details (e.g. name, location)
and will be able to see its statistics.

## Statistics

Gateway statistics are based on the aggregated values sent by the gateway /
packet-forwarder. In case no statistics are visible, it could mean that the
gateway is incorrectly configured.

The statistics grahps display the following values over the past month:
* __Received__ displays the number of Rx packets received from LoRa devices by the gateway.
* __Transmitted__ shows how many Tx packets have been sent from the gateway over LoRa.
* The __/frequency__ graphs display the distribution of packets on each frequency (Hz).
* The __/DR__ graphs show the distribution of packets for each data rate.
* __TX packets / status__ shows the distribution of the Tx packets' status.
