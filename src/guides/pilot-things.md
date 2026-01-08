# Pilot Things integration

[Pilot Things](https://www.pilot-things.com/) is a platform to manage,
visualize, automate and integrate your IoT device fleet. After following this
guide, ChirpStack will be setup to forward device data to
[Pilot Things](https://www.pilot-things.com/).

## Integrate with Pilot Things

### Get Auth Token

In order to let ChirpStack push data to your Pilot Things server, you need to
obtain a Pilot Things _Authentication Token_. To obtain an authentication token,
send an email to [contact@pilot-things.com](mailto:contact@pilot-things.com).

### Setup Pilot Things integration

In the ChirpStack web-interface, navigate to the Application you want to add the
integration. Find the Pilot Things integration under **Integration**, then click
**Add**. Fill in the two required fields:

- Pilot Things server: This is the URL you normally would use to access the
  Pilot Things user interface. For example, https://kerlink.pilot-things.com/.
- Authentication token: This is the token you got in the previous step

## Validate integration

If you completed all the steps, then Pilot Things is ready to receive uplink
data and ChirpStack is configured to forward data for your Device, using th
_Authorization Token_ for authentication.

The last step is to let your device send some data and validate that this data
is received by Pilot Things. You will find this data under the _Activity_ tab
when navigating to the Device within Pilot Things.
