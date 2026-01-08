# FUOTA usage

<!-- toc -->

## Performing a FUOTA deployment

### Prerequisite

Before continuing with the FUOTA deployment, please make sure that:

- In the device-profile of the devices, you have configured the implemented
  application-layer parameters related to FUOTA.
- In the device-profile of the devices, you can configured the correct _Expected
  uplink interval_.
- In case of LoRaWAN 1.0.x, you have configured the _Gen App Key_ for the FUOTA
  devices.

### Create FUOTA deployment

- In the left menu, click **Applications**, then click the application
  containing your FUOTA devices. Click the **FUOTA** tab and then click the
  **New FUOTA deployment** button.

- The most important options to highlight in the form:

  - **Deployment**
    - _Unicast retry count_: This defines the number of times ChirpStack will
      try to send an unicast command. When ChirpStack sends an unicast command
      to the device, it will wait for the _Expected uplink interval_
      (device-profile) until it will retry (if a retry count has configured). A
      higher number does mean the FUOTA process will take longer to complete.
    - _Multicast group-type_: Only select Class-B if this is supported by the
      devices and gateways.
    - _Fragmentation redundancy_: It is recommended to configure this field,
      such that even in case of packet-loss when sending the multicast
      fragments, the device will be able to reconstruct the payload using the
      redundancy fragments.
    - _Calculate multicast-timeout_: If enabled, ChirpStack will try to
      calculate the best multicast-timeout value based on ChirpStack settings
      and number of multicast fragments.
    - _Calculate fragment size_: If enabled, ChirpStack will set the fragment
      size to the max. value allowed based on the _Regional parameters revision_
      selected in the device-profile.

  - **Set device tags (on complete)**
    - If configured, then these tags will be set once a device completes the
      FUOTA process. E.g. you could configure a `version` tag, which will help
      when filtering devices in the device table.

### Add devices

After the FUOTA deployment has been created, you need to add the devices to it.
Please note that you can only add devices with the same device-profile as
selected in the previous step!

- To select the devices, use the checkbox in the first column of each device
  record.
- To add the selected devices, click **Selected devices > Add to FUOTA
  deployment**, select the FUOTA deployment and then click **OK**.

### Add gateways

This step is optional. By default, ChirpStack will try to select the minimum set
of gateways needed to multicast to the devices within the multicast-group. If
you would like to override this behavior:

- In the left menu, click **Gateways**.
- To select the gateways, use the checkbox in the first column of each gateway
  record.
- To add the selected gateways, click **Selected gateways > Add to FUOTA
  deployment**, select the FUOTA deployment and then click **OK**.

### Start FUOTA deployment

Once the devices (and optionally gateways) have been added to the FUOTA
deployment, it is time to start the deployment.

- In the left menu, click **Applications**, click the application, click the
  **FUOTA** tab and then the FUOTA deployment.
- Click the **Start deployment** button.

### State of FUOTA deployment

#### Dashboard

The _Dashboard_ tab will show the executed and running jobs. This view will
automatically refresh. In case there are warnings or errors, you can find the
full message by hovering over the status label.

#### Per device

The _Devices_ tab will show per device when each FUOTA step was completed. This
view will automatically refresh. In case there are warnings or errors, you can
find the full message by hovering over the status label.
