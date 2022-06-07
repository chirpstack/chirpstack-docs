# Device profile

The device profile defines the device capabilities and boot parameters that are
required to establish the (initial) communication between the device and
ChirpStack.

**Note**: RX parameters are used to establish the initial communication.
ChirpStack will send mac-commands to the device to synchronize these parameters
as configured within the ChirpStack configuration files.

## Device-profile templates

On create or update it is possible to use a device-profile template to fill in
most device-profile values. This also includes payload codecs. Please refer to
[Device-profile templates](device-profile-templates.md) for more information.
