# Device-profile templates

Device-profile templates can be created by ChirpStack admin users and can be used
by ChirpStack users as templates to create new device profiles. Device-profile
templates can be created using the web-interface, API or from the command-line
to import from external repositories (read below).

**Note**: As creating device profiles based on these templates copies the 
template values, changing or deleting the device-profile templates will not
affect the created device profiles.

## Import device repository

ChirpStack supports importing the [device repository](https://github.com/TheThingsNetwork/lorawan-devices)
as device-profile templates by executing the following commands from the
Linux command-line. Depending the filesystem permissions, you might need
`sudo` to perform these steps.

### Cloning the device repository

First you need to clone the device repository. In this case, we will clone the
repository to `/opt/lorawan-devices`:

```bash
git clone https://github.com/brocaar/lorawan-devices /opt/lorawan-devices
```

**Note:** an older snapshot of the `lorawan-devices` repository is cloned as the
latest revision no longer contains a `LICENSE` file.

### Import into ChirpStack

Once the repository is cloned, you need to import it into ChirpStack. Please
note that this command can be executed multiple times. For example, you could
periodically pull the latest changes from the `lorawan-devices` repository and
re-run this command to update existing templates / create new templates for
devices that have been added to the repository.

```
chirpstack -c /etc/chirpstack import-legacy-lorawan-devices-repository -d /opt/lorawan-devices
```
