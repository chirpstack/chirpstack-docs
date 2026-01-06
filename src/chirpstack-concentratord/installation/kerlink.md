# Kerlink

<!-- toc -->

## Kerlink iFemtoCell

* [Product detail page](https://www.kerlink.com/product/wirnet-ifemtocell/)
* [Product documentation page](https://wikikerlink.fr/)

**Note:** This was tested with the following firmware version: `5.5.4`.

### SSH into the gateway

The first step is to log in into the gateway using ssh:

```bash
ssh root@GATEWAY-IP-ADDRESS
```

Please refer to the [Kerlink wiki](http://wikikerlink.fr/wirnet-productline)
for login instructions.

### Disable Kerlink CPF

Before you can install the ChirpStack Concentratord, you must disable the 
Kerlink CPF services (`lorad` and `lorafwd`).

Disable `lorad`:

Edit `/etc/default/lorad` and change `DISABLE_LORAD="no"` to `DISABLE_LORAD="yes"`.

Disable `lorafwd`:

Edit `/etc/default/lorafwd` and change `DISABLE_LORAFWD="no"` to `DISABLE_LORAFWD="yes"`.

### Download IPK

Use the following commands to download the latest version of the
`chirpstack-concentratord` package:

```bash
cd /user/.updates
wget https://artifacts.chirpstack.io/downloads/chirpstack-concentratord/vendor/kerlink/ifemtocell/chirpstack-concentratord_%CONCENTRATORD_VERSION-r1_klkgw.ipk
```

### Install IPK

Run the following command to install the IPK package:

```bash
sync
kerosd -u
reboot
```

### Configuration

Configuration files can be found at:

* `/etc/chirpstack-concentratord/concentratord.toml` - Concentratord configuration
* `/etc/chirpstack-concentratord/channels.toml` - Channel configuration

#### (Re)start and stop commands

Use the following commands to (re)start and stop the ChirpStack Concentratord
service:

```bash
# Status
sudo monit summary

# start
monit start chirpstack-concentratord

# stop
monit stop chirpstack-concentratord

# restart
monit restart chirpstack-concentratord
```

#### Log output

To view the ChirpStack Concentratord log output, use the following command:

```bash
tail -f -n 100 /var/log/messages |grep chirpstack-concentratord
```
