package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/roster90/grp_greet/calculator/proto"
)

func doFindMax(c pb.CalculatorServiceClient) {
	log.Printf("Starting to do a doFindMax invoke")

	numbers := []int32{1, 5, 3, 6, 2, 20}

	stream, err := c.FindMaximum(context.Background())

	if err != nil {
		log.Fatalf("Failed to calling doFindMax: %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		for _, number := range numbers {
			log.Printf("Sending request: %v\n", number)

			stream.Send(&pb.FindMaximumRequest{
				Number: number,
			})
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
			log.Printf("doFindMax: %v", res.GetResult())
		}
		close(waitc)

	}()
	<-waitc
}
