import os
import sys

import grpc
from chirpstack_api import api

# Configuration.

# This must point to the API interface.
server = "localhost:8080"

# The DevEUI for which you want to enqueue the downlink.
dev_eui = "0101010101010101"

# The API token (retrieved using the web-interface).
api_token = "..."

if __name__ == "__main__":
  # Connect without using TLS.
  channel = grpc.insecure_channel(server)

  # Device-queue API client.
  client = api.DeviceServiceStub(channel)

  # Define the API key meta-data.
  auth_token = [("authorization", "Bearer %s" % api_token)]

  # Construct request.
  req = api.EnqueueDeviceQueueItemRequest()
  req.queue_item.confirmed = False
  req.queue_item.data = bytes([0x01, 0x02, 0x03])
  req.queue_item.dev_eui = dev_eui
  req.queue_item.f_port = 10

  resp = client.Enqueue(req, metadata=auth_token)

  # Print the downlink id
  print(resp.id)
