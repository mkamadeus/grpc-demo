package main

import (
	"fmt"
	"log"
	"net"

	"github.com/mkamadeus/grpc-demo/server/controllers"
	"github.com/mkamadeus/grpc-demo/server/schemas"
	"github.com/mkamadeus/grpc-demo/server/servers"
	"google.golang.org/grpc"
)

const PORT = 1337

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", PORT))
	if err != nil {
		log.Fatalf("failed to listed: %v", err)
	}
	defer listener.Close()

	weatherCtl := controllers.NewWeatherController()
	server := &servers.WeatherServer{WeatherCtl: weatherCtl}

	grpcServer := grpc.NewServer()
	schemas.RegisterWeatherServer(grpcServer, server)

	log.Printf("Serving on port: %d", PORT)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
