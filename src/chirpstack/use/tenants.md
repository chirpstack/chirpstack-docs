# Tenants

A tenant can be used to let organizations or teams manage their
own applications and optionally their own gateways.

An tenant can have:

* Device profiles
* Gateways (if allowed)
* Applications and devices
* Users

## Device Profiles

[Device profiles](device-profiles.md) can be created by (tenant) admin users
and can be assigned when creating a [Device](devices.md).

## Gateways

An tenant can manage its own set of gateways. Note that when an tenant
is created by a global administrator, it can decide that an tenant can not
have any gateways. In this case, the gateway option is not available to an
tenant.

That an tenant is able to manage its own set of gateways does not mean
that the coverage is limited to this set of gateways. Gateways connectivity
will be shared across the whole network.

## Applications

[Applications](applications.md) can be created by (tenant)
admin users and define a group of devices with the same purpose.

## Users

Users can be assigned to a tenant to grant them access to the
tenant. Within the context of that assignment, an user can be an
tenant administrator or a regular user. Please note that only existing users
can be assigned to a tenant.

### Tenant administrator

An tenant administrator is authorized to manage the users assigned
with the tenant and manage the gateways, applications and devices.

### Regular user

Regular users are able to see all data, but are not able to make any
modifications.
