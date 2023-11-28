# Guessing game with gRPC

This example provides a simple version of the guessing game example using gRPC.
At present, it only supports sending guesses--no other features from later
lectures are supported.

## Relevant files

 - `proto/guessing-game.proto`:  gRPC/protobuf specification for the messages
 - `cmd/server/main.go`:  Guessing game server (Golang)
 - `client/client.py`:  Guessing game client (Python)
 - `pkg/proto/*`:  protobuf-generated RPC code (Server)
 - `client/*_pb2.py`:  protobuf-generated RPC code (Client)

 ## Setup

 You can compile the server using `make`.  To run the client, you will need a
 Python environment with gRPC--to set this up, you can create a virtual
 environment using the provided `requirements.txt`.

 If you want to change the protobuf spec, you need to recompile the generated
 code.  To do this, you need protobuf and gRPC installed for each language for
 which you intend to generate code (in this case, golang and Python).  For
 details on how to do this yourself, see gRPC's setup docs:
- Go: https://grpc.io/docs/languages/go/quickstart/
- Python: https://grpc.io/docs/languages/python/quickstart/