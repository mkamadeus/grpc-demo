module github.com/mkamadeus/grpc-demo/server

go 1.17

require (
	github.com/spf13/cobra v1.1.3
	google.golang.org/grpc v1.40.0
	google.golang.org/protobuf v1.27.1
)

require (
	github.com/golang/protobuf v1.5.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/net v0.0.0-20200822124328-c89045814202 // indirect
	golang.org/x/sys v0.0.0-20200323222414-85ca7c5b95cd // indirect
	golang.org/x/text v0.3.2 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
)

replace github.com/grpc-demo/protobuf => ../protobuf
