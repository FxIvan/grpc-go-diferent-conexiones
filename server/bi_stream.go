package main

import (
	"io"
	"log"

	pb "github.com/fxivan/go-grpc-akhil/proto"
)

func (s *helloServer) SayHelloBidirectionalStreaming(stream pb.GreetService_SayHelloBidirectionalStreamingServer) error {
	log.Print("Init SayHelloBidirectionalStreaming into Server")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		log.Printf("Got request with name : %v", req.Name)

		res := &pb.HelloResponse{
			Message: "Hello into Server: " + req.Name,
		}

		if err := stream.Send(res); err != nil {
			return err
		}
	}
}
