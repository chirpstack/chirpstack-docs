# C# examples

- [Call gRPC services with the .NET client](https://learn.microsoft.com/en-us/aspnet/core/grpc/client?view=aspnetcore-8.0)

## Preparation

Need to install several NuGet packages:

- Chirpstack.Api
- Grpc.Net.Client.Web

## Enqueue

The example below demonstrates:

- Configuration of gRPC _dial options_ including API token
- Connect to a gRPC API
- Define service client (in this case for the `DeviceService`)
- Perform an API call for a service (in this case `Enqueue`)

### `Program.cs`

```csharp
{{#include ../../../examples/chirpstack/api/csharp/Program.cs}}
```
