module github.com/mkamadeus/grpc-demo/server

go 1.17

require (
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013
	google.golang.org/grpc v1.40.0
	google.golang.org/protobuf v1.27.1
)

replace github.com/grpc-demo/protobuf => ../protobuf
