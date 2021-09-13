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
	weatherServer := &servers.WeatherServer{WeatherCtl: weatherCtl}

	studentCtl, err := controllers.NewStudentController()
	if err != nil {
		log.Fatalf("failed to create student controller: %v", err)
	}
	studentServer := &servers.StudentServer{StudentCtl: studentCtl}

	grpcServer := grpc.NewServer()
	schemas.RegisterWeatherServer(grpcServer, weatherServer)
	schemas.RegisterStudentServer(grpcServer, studentServer)

	log.Printf("Serving on port: %d", PORT)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
