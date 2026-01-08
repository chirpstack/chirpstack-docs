# Python examples

* [Python gRPC quickstart](https://grpc.io/docs/languages/python/quickstart/)
* [`chirpstack-api` Python package](https://pypi.org/project/chirpstack-api/)

ChirpStack provides a Python package `chirpstack-api` that can be installed
using `pip`:

```bash
pip install chirpstack-api
```

## Enqueue downlink

The example below demonstrates:

* Connecting to a gRPC server
* Defining a service client / stub
* Performing an API call (in this case `Enqeue`)

### `enqueue_downlink.py`

```python
{{#include ../../../examples/chirpstack/api/python/enqueue_downlink.py}}
```
