# gRPC

ChirpStack provides a gRPC API interface which can be used to integrate ChirpStack
with external application and / or to integrate ChirpStack into other platforms.

## gRPC

[gRPC](https://grpc.io/) is a high-performance, open-source universal RPC
framework. [Protocol Buffers](https://developers.google.com/protocol-buffers)
definitions are used to define this API. All definitions are hosted in the
[chirpstack](https://github.com/chirpstack/chirpstack/tree/master/api/proto) repository.

Using the gRPC toolsets, it is possible to generate API clients for various
programming languages. The following languages are officially supported by
gRPC but there are additional community-based implementations:

* C++
* Go
* Node.js
* Java
* Ruby
* Android Java
* PHP
* Python
* C#
* Objective-C

### Authentication

In order to use the gRPC API methods, you must provide per-RPC credentials.
The key for this metadata is `authorization`, the value `Bearer <API TOKEN>`
(replace **&lt;API TOKEN&gt;** with the API TOKEN obtained using the web-interface).

### ChirpStack SDKs

The ChirpStack project provides SDKs for the following programming languages:

* [Go](https://pkg.go.dev/github.com/chirpstack/chirpstack/api/go/v4)
* [Python](https://pypi.org/project/chirpstack-api/)
* [JavaScript](https://www.npmjs.com/package/@chirpstack/chirpstack-api)
* [Rust](https://crates.io/crates/chirpstack_api)

### API console / UI

ChirpStack itself does not provide an API console / UI interface, but it is
compatible third-party gRPC GUI applications. See for example:

* [BloomRPC](https://github.com/bloomrpc/bloomrpc)
* [gRPC UI](https://github.com/fullstorydev/grpcui)
