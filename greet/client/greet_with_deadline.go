package main

import (
	"context"
	"log"
	"time"

	pb "github.com/roster90/grp_greet/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Printf("Starting to do a GreetWithDeadline RPC...")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	defer cancel()

	req := &pb.GreetRequest{
		FirstName: "Roster",
	}
	res, err := c.GreetWithDeadline(ctx, req)

	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Printf("Timeout was hit! Deadline was exceeded")
				return
			} else {
				log.Fatalf("Unexpected error: %v", e)
			}
		} else {
			log.Fatalf("A non gRPC error %v", e)
		}
	}
	log.Printf("GreetWithDeadline: %s", res.Result)
}
