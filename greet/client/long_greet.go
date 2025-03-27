package main

import (
	"context"
	"log"
	"time"

	pb "github.com/roster90/grp_greet/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {

	log.Printf("Starting to do a doLongGreet invoke")

	req := []*pb.GreetRequest{
		{FirstName: "Roster"},
		{FirstName: "Kelley"},
		{FirstName: "XUnicorn"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Failed to calling doLongGreet: %v", err)
	}

	for _, req := range req {
		log.Printf("Sending request: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Failed to receive response from LongGreet: %v", err)
	}

	log.Printf("LongGreet Response: %v\n", res.GetResult())

}
