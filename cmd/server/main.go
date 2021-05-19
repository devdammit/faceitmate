package main

import (
	"github.com/devdammit/faceitmate/pkg/api"
	"github.com/devdammit/faceitmate/pkg/faceitmate"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	server := grpc.NewServer()
	srv := faceitmate.NewGRPCServer()

	api.RegisterFaceitmateServer(server, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := server.Serve(l); err != nil {
		log.Fatal(err)
	}
}
