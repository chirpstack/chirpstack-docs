const grpc = require("@grpc/grpc-js");
const device_grpc = require("@chirpstack/chirpstack-api/api/device_grpc_pb");
const device_pb = require("@chirpstack/chirpstack-api/api/device_pb");

// This must point to the ChirpStack API interface.
const server = "localhost:8080";

// The DevEUI for which we want to enqueue the downlink.
const devEui = "0101010101010101";

// The API token (can be obtained through the ChirpStack web-interface).
const apiToken = "...";


// Create the client for the DeviceService.
const deviceService = new device_grpc.DeviceServiceClient(
  server,
  grpc.credentials.createInsecure(),
);

// Create the Metadata object.
const metadata = new grpc.Metadata();
metadata.set("authorization", "Bearer " + apiToken);

// Enqueue downlink.
const item = new device_pb.DeviceQueueItem();
item.setDevEui(devEui);
item.setFPort(10);
item.setConfirmed(false);
item.setData(new Uint8Array([1, 2, 3]));

const enqueueReq = new device_pb.EnqueueDeviceQueueItemRequest();
enqueueReq.setQueueItem(item);

deviceService.enqueue(enqueueReq, metadata, (err, resp) => {
  if (err !== null) {
    console.log(err);
    return;
  }

  console.log("Downlink has been enqueued with id: " + resp.getId());
});
