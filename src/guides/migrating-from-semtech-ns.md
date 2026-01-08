# Migrating from Semtech Network Server

Semtech has announced that it will discontinue its
[Semtech Network Server](https://ns.loracloud.com/), which is based on
ChirpStack (v4). This guide describes how you can set up your own ChirpStack
instance and migrate your devices. As ChirpStack is an open-source
LoRaWAN<sup>&reg;</sup> Network Server, there won’t be additional charges apart
from hosting costs.

<!-- toc -->

## What is the Semtech Network Server

The Semtech Network Server is a by Semtech hosted instance of ChirpStack (v4) to
help developers get up to speed with testing their LoRaWAN projects. Users can
register up to 3 gateways and 10 devices for testing and evaluation.

## When is the Semtech Network Server sunset

The sunset of the Semtech Network Server will be October 31. Before this time,
you need to migrate your devices.

## What do I need to do

As the Semtech Network Server is based on ChirpStack, all you need to do is set
up your own ChirpStack instance and migrate your gateways and devices.

## Setup ChirpStack

Below you will find different ways how you can set up ChirpStack on a Raspberry
Pi, your local machine or a cloud-hosted virtual machine.

### ChirpStack Gateway OS

If you are using a Raspberry Pi as LoRa® gateway, then the easiest way to set up
ChirpStack is by using the ChirpStack Gateway OS Full image. This is ready to
use SD Card image which includes the ChirpStack Network Server. On first boot,
it will create a Wi-Fi access-point to which you can connect for the
configuration of LoRa concentrator shield and LoRaWAN region through the
ChirpStack Gateway OS web-interface.

If you have more than one Raspberry Pi based LoRa gateway, then it is
recommended to install the Base image and connect this to a ChirpStack instance
hosted on a separate machine. In this case, you should combine the ChirpStack
Gateway OS (Base) installation with one of the ChirpStack deployments listed
below.

Please refer to the
[Raspberry Pi installation guide](../chirpstack-gateway-os/install/raspberry-pi.md)
for installation instructions.

### Docker Compose

The Docker Compose example can be used to set up either a local ChirpStack
installation on your computer or to deploy ChirpStack on a cloud-hosted virtual
machine. The Docker Compose example will setup and configure the databases, the
MQTT broker and the ChirpStack Network Server. You can connect your gateway
either using the Semtech UDP Packet Forwarder, Semtech Basics Station or using
the ChirpStack MQTT protocol.

Please refer to the
[Docker Compose installation guide](../getting-started/docker.md) for
installation instructions.

### Debian / Ubuntu VM using Ansible

If you would like to install ChirpStack onto a Debian or Ubuntu machine, then
the easiest way is by using the
[ChirpStack Ansible Playbook](https://github.com/chirpstack/chirpstack-ansible-playbook).
[Ansible](https://github.com/ansible/ansible) is an IT automation tool and the
ChirpStack Ansible Playbook automates the setup of ChirpStack including all its
dependencies and the initialization of the database. Optionally, it can also
automate the request and setup of TLS certificates.

Please refer to the
[Ansible installation guide](../getting-started/ansible-vagrant.md) for
installation instructions.

## Migrating gateway(s)

You need to manually copy the Gateway ID(s) from the Semtech Network Server and
re-create these in your local ChirpStack instance. Then you need to re-configure
your gateway to point to your local ChirpStack installation. Per ChirpStack
deployment option, we describe which actions you need to perform.

### ChirpStack Gateway OS

No setup is required, as this is automated by the ChirpStack Gateway OS
web-interface configuration.

### Docker Compose

Login into the ChirpStack web-interface to add your gateway. Based on the
software running on the gateway, you must update your gateway configuration.

#### Semtech UDP Packet Forwarder

You need to point your gateway to the IP address or hostname of the machine
running the Docker Compose instance. By default port `1700` is used for uplink
and downlink.

#### Semtech Basics Station

You need to point your gateway to `ws://IP-OR-HOST:3001`, where `IP-OR-HOST` is
the IP address or hostname of the machine running the Docker Compose example. By
default no TLS certificate or credentials are required to connect to the Basics
Station endpoint.

#### ChirpStack MQTT Forwarder / ChirpStack Gateway Bridge

You need to point your gateway to `tcp://IP-OR-HOST:1883`, where `IP-OR-HOST` is
the IP address or hostname of the machine running the Docker Compose example. By
default no TLS certificate or credentials are required to connect to the MQTT
broker.

### Debian / Ubuntu VM using Ansible

Login into the ChirpStack web-interface to add your gateway. Based on the
software running on the gateway, you must update your gateway configuration.

#### Semtech UDP Packet Forwarder

You need to point your gateway to the IP address or hostname of the virtual
machine. By default port `1700` is used for uplink and downlink.

#### Semtech Basics Station

You need to point your gateway to `wss://IP-OR-HOST:3001`, where `IP-OR-HOST` is
the IP address or hostname of the virtual machine. You must obtain the gateway
CA and TLS certificates through the ChirpStack web-interface.

#### ChirpStack MQTT Forwarder / ChirpStack Gateway Bridge

You need to point your gateway to `ssl://IP-OR-HOST:8883`, where `IP-OR-HOST` is
the IP address or hostname of the virtual machine. You must obtain the gateway
CA and TLS certificates through the ChirpStack web-interface.

## Migrating device(s)

You need to re-create the devices provisioned in the Semtech Network Server in
your local ChirpStack installation. Make sure that you will use the same Device
Profile settings for your device(s), and that you do not make any mistakes when
copying the DevEUI and AppKey of your devices.

Please refer to the
[Connecting a device guide](https://www.chirpstack.io/docs/guides/connect-device.html)
for more information on connecting your devices and troubleshooting connectivity
issues.

Important: After migrating your device(s), please make sure to disable these
devices in the Semtech Network Server. You will find the disable device option
in the device configuration screen. After this, devices need to be reactivated
(e.g. rejoin in case of OTAA).
