.PHONY: proto

build:
	go build ./cmd/server

proto:
	protoc --go_out=. \
	--go-grpc_out=. \
	proto/guessing-game.proto

	python3 -m grpc_tools.protoc -I. --python_out=./client --grpc_python_out=./client proto/guessing-game.proto

clean:
	rm -fv client server