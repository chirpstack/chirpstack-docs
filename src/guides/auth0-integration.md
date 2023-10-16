# Auth0.com integration

This guide describes the steps to setup ChirpStack with [Auth0.com](https://auth0.com/)
authentication backend using the [OpenID Connect (OIDC)](https://openid.net/developers/how-connect-works/)
authentication option.

Before starting with this guide:

* Make sure you have a fully functioning ChirpStack setup
* You have an Auth0.com account

During this guide we assume that the URL to access the ChirpStack web-interface
is `https://example.com/`. You need to replace this with the actual URL (or IP)
of your ChirpStack installation!

## Auth0.com configuration

### Create application

Within the Auth0.com console, click the **Applications > Applications** in the
menu. Then click the **Create Application** button and select the
_Regular Web Application_ option. Click **Create** to create the application.

Click on the **Settings** tab. Here you must configure the following options:

* Application URIs
    * **Allowed Callback URLs** This must be configured to the
      `/auth/oidc/callback` endpoint of your ChirpStack installation, e.g.
      `https://example.com/auth/oidc/callback`.

Click **Save Changes**.

## ChirpStack configuration

In order to make use of Auth0.com, you must make some modifications to you
ChirpStack [Configuration](../chirpstack/configuration.md).

Example OpenID Connection configuration section:

```toml
[user_authentication]

  [user_authentication.openid_connect]
    enabled=true
    registration_enabled=false
    registration_callback_url=""
    provider_url="https://tenant.eu.auth0.com/"
    client_id="..."
    client_secret="..."
    redirect_url="https://example.com/auth/oidc/callback"
    login_label="Log in using Auth0.com"
```

Please see the [Configuration](../chirpstack/configuration.md) page for
documentation of these options. Below we will go over these fields within
the context of the Auth0.com setup:

* **`enabled`**: You need to set this to `true`.
* **`registration_enabled`**: Set this to `true` to in case you would like to
  create a new ChirpStack user, in case the Auth0.com user does not (yet) exist
  in ChirpStack. In case you leave this to `false` then a ChirpStack user must
  exist with the same e-mail address as provided by Auth0.com.
* **`registration_callback_url`**: In case the `registration_enabled` is set to
  `true`, most likely you would like to set this URL to an internal endpoint
  which can take care of onboarding new users. E.g. this service could create
  a new Tenant and assign the user to this tenant.
* **`provider_url`**: You will find this information in the Auth0.com
  **Settings** tab under **Domain**. For example if your **Domain** is
  `example.eu.auth0.com`, then you must use `provider_url="https://example.eu.auth0.com/"`.
* **`client_id`**: You will find this information in the Auth0.com
  **Settings** tab under **Client ID**.
* **`client_secret`**: You will find this information in the Auth0.com
  **Settings** tab under **Client Secret**.
* **`redirect_url`**: This must be configured to the
      `/auth/oidc/callback` endpoint of your ChirpStack installation, e.g.
      `https://example.com/auth/oidc/callback`.
* **`login_label`**: The label to use under the login form.

Make sure to restart ChirpStack after making configuration changes.

## Validation

When navigating to the ChirpStack login page, you should no longer see the
username and password fields. Instead there should be a button which once
clicked, will redirect you to the Auth0.com login screen of your Auth0.com
tenant.
