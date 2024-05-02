package main

import (
	"io"

	pb "github.com/fxivan/go-grpc-akhil/proto"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var messages []string

	for {

		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessageList{Messages: messages})
		}

		if err != nil {
			return err
		}

		messages = append(messages, "Hola "+req.Name)
	}

}
