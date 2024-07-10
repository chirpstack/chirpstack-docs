[![Milesight Logo](http://resource.milesight-iot.com/Support/lockon/ChirpStack-v4-doc/Milesight-logo.png)](https://www.milesight.com/)

# How to Connect a Milesight Gateway to Chirpstack v4 Platform

## Table of Contents

1. [Prerequisites](#prerequisites)
   - [Device Details](#device-details)
   - [Network Connectivity](#network-connectivity)
2. [Configuring the Gateway](#configuring-the-gateway)
   - [Accessing the Gateway Management Interface](#accessing-the-gateway-management-interface)
   - [Setting Up Multi-Destination](#setting-up-multi-destination)
   - [Saving and Applying the Configuration](#saving-and-applying-the-configuration)
3. [Adding the Gateway to Chirpstack](#adding-the-gateway-to-chirpstack)
   - [Logging into Chirpstack](#logging-into-chirpstack)
   - [Adding the Gateway](#adding-the-gateway)
   - [Verifying the Gateway Connection](#verifying-the-gateway-connection)
4. [Finally](#finally)

## Prerequisites

### Device Details
- [UG65 Product page](https://www.milesight.com/iot/product/lorawan-gateway/ug65)
- [UG67 Product page](https://www.milesight.com/iot/product/lorawan-gateway/ug67)

**Note:** UG65/UG67 has built-in ChirpStack MQTT Forwarder, please ensure the firmware version is 60.0.0.43 and later.



### Network Connectivity

- Ensure that the gateway and the Chirpstack v4 system are on the same network and can communicate with each other.

  

## Configuring the Gateway

### Accessing the Gateway Management Interface

1. Log in to the gateway's management interface using the default credentials (username: `admin`, password: `password`).
2. Navigate to `Packet Forward -> General`.
3. Find `Multi-Destination` and click the `Operation` button for `Embedded NS`.
4. In the pop-up window, uncheck the `Enable` option and click `Save`.

### Setting Up Multi-Destination

1. Under the `Multi-Destination` section, click the `+` button.
2. In the pop-up window, check the `Enable` option.
3. Select `ChirpStack-v4` from the `Type` dropdown menu.
4. Enter the Chirpstack address in the `Server Address` field (e.g., `192.168.45.221`).
5. The default `MQTT Port` is `1883`. If you have changed it, enter the modified port number.
6. Select the appropriate `Region ID` based on your device's frequency band (e.g., `US915` for this test).
7. `User Credentials` and `TLS Authentication` can be left empty unless your setup requires them.
8. Click `Save` to close the window.

### Saving and Applying the Configuration

1. At the bottom of the `Multi-Destination` page, click `Save & Apply`.
2. Wait 3-5 minutes and refresh the page.
3. Check the `Connect Status` under the `Multi-Destination` section for the `Type` `ChirpStack-v4`. If it shows `Connected`, the gateway configuration is successful.
   <br><img src="http://resource.milesight-iot.com/Support/lockon/ChirpStack-v4-doc/Config-Gateway-02.png" alt="Gateway configuration complete" style="width:50%;">

## Adding the Gateway to Chirpstack

### Logging into Chirpstack

1. Log in to the Chirpstack management interface using the default credentials (username: `admin`, password: `admin`).

### Adding the Gateway

1. Navigate to `Tenants -> Gateways`.
2. Click `Add gateway`.
3. Fill in the necessary information in the pop-up window:
   - `Name`: Enter `UG65-US915`.
   - `Gateway ID (EUI64)`: Enter `24E124FFFEF5408E`.
   - Fill in other parameters as required.
4. Click `Submit`.
   <br><img src="http://resource.milesight-iot.com/Support/lockon/ChirpStack-v4-doc/Config-Chirpstack-01.png" alt="Config-Chirpstack" style="width:50%;">

### Verifying the Gateway Connection

1. Navigate back to `Tenants -> Gateways`.
2. Check the status of the newly added gateway. If it shows `Online`, the gateway has been successfully added and is operational.
3. Click on the blue ID code under the `Gateway ID` column.
4. In the new interface, go to the `LoRaWAN frames` tab.
5. Ensure there is data scrolling, indicating a successful data link between the gateway and Chirpstack.
   <img src="http://resource.milesight-iot.com/Support/lockon/ChirpStack-v4-doc/Config-Chirpstack-02.png" alt="Config-Chirpstack" style="width:50%;">

## Finally

The gateway has been successfully connected to the Chirpstack v4 platform.
