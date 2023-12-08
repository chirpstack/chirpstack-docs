using Chirpstack.Api;
using Google.Protobuf;
using Grpc.Core;
using Grpc.Net.Client;
using Grpc.Net.Client.Web;

// Configuration.

// This must point to the API interface.
var server = "localhost:8080";

// The DevEUI for which you want to enqueue the downlink.
var devEui = "0101010101010101";

// The API token (retrieved using the web-interface).
var apiToken = "...";

// Connect without using TLS.
using var channel = GrpcChannel.ForAddress(
    address: server,
    channelOptions: new GrpcChannelOptions()
    {
        HttpHandler = new GrpcWebHandler(new HttpClientHandler()),
        Credentials = ChannelCredentials.Insecure,
    }
);

// Device-queue API client.
var client = new DeviceService.DeviceServiceClient(channel);

// Define the API key meta-data.
var authToken = new Metadata
{
    { "Authorization", "Bearer " + apiToken },
};

// Construct request.
var req = new EnqueueDeviceQueueItemRequest()
{
    QueueItem = new DeviceQueueItem()
    {
        Confirmed = false,
        Data = ByteString.CopyFrom(0x01, 0x02, 0x03),
        DevEui = devEui,
        FPort = 10,
    },
};

var resp = await client.EnqueueAsync(req, headers: authToken);

// Print the downlink id.
Console.WriteLine(resp.Id);
Console.ReadKey();
