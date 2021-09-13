module github.com/mkamadeus/grpc-demo/server

go 1.17

require (
	github.com/spf13/cobra v1.1.3
	google.golang.org/grpc v1.40.0
	google.golang.org/protobuf v1.27.1
)

replace github.com/grpc-demo/protobuf => ../protobuf
