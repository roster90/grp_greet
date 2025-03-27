package main

import (
	"context"
	"log"
	"time"

	pb "github.com/roster90/grp_greet/calculator/proto"
)

func doAverage(c pb.CalculatorServiceClient) {

	log.Printf("Starting to do a doLongGreet invoke")

	numbers := []float32{12223, 111, 100}

	stream, err := c.ComputeAverage(context.Background())

	if err != nil {
		log.Fatalf("Failed to calling doAverage: %v", err)
	}

	for _, number := range numbers {
		log.Printf("Sending request: %v\n", number)

		stream.Send(&pb.ComputeAverageRequest{
			Number: number,
		})
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Failed to receive response from ComputeAverage: %v", err)
	}

	log.Printf("ComputeAverage Response: %v\n", res.GetResult())

}
