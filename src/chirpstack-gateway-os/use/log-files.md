# Log files

## Web-interface

To read log output using the web-interface, click **Status** and then
**System Log** in the left menu.

## SSH

To read log output using SSH, you need to use the `logread` utility.
Examples:

```bash
# Print log output
logread

# Get last 20 lines and follow log messages
logread -n 20 -f 
```
