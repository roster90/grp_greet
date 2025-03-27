package main

import (
	"context"
	"log"

	pb "github.com/roster90/grp_greet/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {

	log.Printf("Starting to do a Greet RPC...")

	req := &pb.GreetRequest{
		FirstName: "Roster",
	}

	res, err := c.Greet(context.Background(), req)

	if err != nil {
		log.Fatalf("Failed to greet: %v", err)
	}
	log.Printf("Greet response: %s", res.Result)
}
