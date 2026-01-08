# Modifying files

ChirpStack Gateway OS uses an
[OverlayFS](https://en.wikipedia.org/wiki/OverlayFS) on top of the read-only
root filesystem. This means you can make modifications that are persisted even
after a software update as these changes are written to a different data
partition.

## Important

Before performing a backup or upgrade, please make sure that the directories /
files are listed in the files that must be included during a sysupgrade. This
list can be configured in the web-interface by clicking the **System**, then
**Backup / Flash Firmware** in the left menu and then clicking the
**Configuration** tab.

## Reset

To perform a factory reset using the web-interface, click **System**, then
**Backup / Flash Firmware** in the left menu. Then click **Perform reset**.
