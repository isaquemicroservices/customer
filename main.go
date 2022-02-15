package main

import (
	"log"
	"net"

	app "github.com/isaqueveraslabs/customer-microservice/application/customer"
	config "github.com/isaqueveraslabs/customer-microservice/configuration"
	inter "github.com/isaqueveraslabs/customer-microservice/interfaces/customer"
	gogrpc "google.golang.org/grpc"
)

func main() {
	// Loading config of system
	config.Load()

	// Listen on port
	listen, err := net.Listen("tcp", config.Get().Address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err.Error())
	}

	// Creating new server
	server := gogrpc.NewServer()

	// customer registration server
	app.RegisterCustomerServer(server, &inter.Server{})

	// Message of success
	log.Println(config.Get().Name, "is running in port", config.Get().Address)

	// Initializing server
	if err = server.Serve(listen); err != nil {
		log.Fatalf("Failed to server: %v", err.Error())
	}
}
