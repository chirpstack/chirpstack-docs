# ChirpStack Concentratord

The Concentratord backend implements the ChirpStack Concentratord
[ZeroMQ](https://zeromq.org/) based interface.

## Deployment

The ChirpStack Gateway Bridge and the ChirpStack Concentratord must be deployed
on the gateway.

## Prometheus metrics

The ChirpStack Concentratord backend exposes several
[Prometheus](https://prometheus.io/) metrics for monitoring.

#### backend_concentratord_event_count

The number of received events from Concentratord (per event type).

#### backend_concentratord_command_count

The number of commands sent to Concentratord (per command type).
