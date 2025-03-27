package main

import (
	"context"
	"io"
	"log"

	pb "github.com/roster90/grp_greet/calculator/proto"
)

func doPrime(c pb.CalculatorServiceClient) {

	log.Printf("Starting to do a Greet RPC...")

	req := &pb.PrimesRequest{
		Number: 120,
	}

	stream, err := c.Primes(context.Background(), req)

	if err != nil {
		log.Fatalf("Failed to greet: %v", err)
	}

	var result string = ""
	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			// we've reached the end of the stream
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}
		result += "*" + string(msg.GetResult())
	}
	log.Printf("Prime number: %s", result)

}
