package main

import (
	"context"
	"log"

	pb "github.com/roster90/grp_greet/calculator/proto"
)

func doSumNumber(c pb.CalculatorServiceClient) {

	log.Printf("Starting to do a Greet RPC...")

	req := &pb.CalculatorRequest{
		Number1: 1,
		Number2: 2,
	}

	res, err := c.Sum(context.Background(), req)

	if err != nil {
		log.Fatalf("Failed to greet: %v", err)
	}
	log.Printf("Sum request is : %d", res.Result)
}
