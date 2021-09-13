package main

import (
	"fmt"
	"log"
	"net"

	"github.com/mkamadeus/grpc-demo/server/schemas/"
	"google.golang.org/grpc"
)

const PORT = 1337

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", PORT))
	if err != nil {
		log.Fatalf("failed to listed: %v", err)
	}
	defer listener.Close()

	server := &schemas.Server{}

	grpcServer := grpc.NewServer()
	schemas.RegisterWeatherServer(grpcServer, server)

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
