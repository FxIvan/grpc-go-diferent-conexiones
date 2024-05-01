package main

import (
	"log"
	"net"

	pb "github.com/fxivan/go-grpc-akhil/proto"
	"google.golang.org/grpc"
)

const (
	port = ":50410"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	list, err := net.Listen("tcp", "0.0.0.0"+port)
	if err != nil {
		log.Fatalf("Failed to start server %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Printf("Server started at %v", list.Addr())
	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
}
