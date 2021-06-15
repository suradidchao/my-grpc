# My gRPC

## Compile protobuf
```bash
protoc --go-grpc_out=. --go_out=. <path-to-proto-file>/<file-name>.proto
```

## Run testing UI
```bash
grpcui -plaintext 127.0.0.1:9000
```

## Run gRPC Server
```bash
cd server
go run main.go
```

## Run gRPC Client
```bash
cd client
go run client.go
```
