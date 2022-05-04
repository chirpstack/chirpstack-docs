# Changelog

## v4.0.0 (in development)

### Main changes

#### ChirpStack v4

This release replaces ChirpStack Network Server (v3) and ChirpStack Application
Server (v3) with ChirpStack v4.0.0.

#### MQTT certificates

If using the _full_ image, MQTT TLS certificates will be automatically generated on
first boot and configured in ChirpStack. For consuming device events throug the
MQTT integration, you can generate certificates in the web-interface under
the application integrations.

**Note:** localhost connections (e.g. if using Node-RED) to the MQTT do not
need TLS certificates for authentication and authorization.

#### gateway-config improvements

If using the _full_ ChirpStack Gateway OS image, which includes the ChirpStack
Network Server, the `gateway-config` script will automatically create the
gateway in the ChirpStack database for you.

#### Region configuration

The _full_ image contains update region configuration for the common regions.
The desired region can be selected using the `gateway-config` configuration script.

#### New shields

Support for the following shields has been added:

- RAK2247
- RAK5147

#### Yocto Kirkstone

ChirpStack Gateway OS has been updated to Yocto Kirkstone. This should fix
boot problems on some Pi 4 revisions, as this release contains an updated
bootloader.

### Other updates

* ChirpStack Concentratord to v3.2.2
* ChirpStack Gateway Bridge to v3.13.1
* Node-RED to v2.1.6
