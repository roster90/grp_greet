package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/roster90/grp_greet/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) error {
	log.Printf("GreetEveryone function was invoked with a streaming request\n")

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Printf("Error while reading client stream: %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Roster"},
		{FirstName: "Kelley"},
		{FirstName: "XUnicorn"},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Sending request: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)

		}
		if err := stream.CloseSend(); err != nil {
			log.Printf("Error closing send stream: %v", err)
		}
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				// we've reached the end of the stream
				break
			}

			if err != nil {
				log.Fatalf("Error while reading stream: %v", err)
			}
			log.Printf("doGreetEveryone: %s", res.GetResult())
		}
		close(waitc)

	}()
	<-waitc

	return nil
}
