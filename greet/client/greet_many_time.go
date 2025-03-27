package main

import (
	"context"
	"io"
	"log"

	pb "github.com/roster90/grp_greet/greet/proto"
)

func doGreetManyTime(c pb.GreetServiceClient) {

	log.Printf("Starting to do a Greet RPC...")

	req := &pb.GreetRequest{
		FirstName: "Roster",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Failed to greet: %v", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			// we've reached the end of the stream
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}
		log.Printf("Greet many time: %s", msg.GetResult())
	}

}
