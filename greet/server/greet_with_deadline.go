package main

import (
	"context"
	"log"
	"time"

	pd "github.com/roster90/grp_greet/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GreetWithDeadline(c context.Context, req *pd.GreetRequest) (*pd.GreetResponse, error) {
	log.Printf(" GreetWithDeadline function was invoked with %v\n", req)
	firstName := req.GetFirstName()

	for range 3 {
		if c.Err() == context.DeadlineExceeded {
			log.Println("The client canceled the request!")
			return nil, status.Error(codes.Canceled, "The client canceled the request")
		}
		// Simulate a long running task
		// Sleep for 1 second
		time.Sleep(1 * time.Second)
	}

	return &pd.GreetResponse{
		Result: "Hello, " + firstName,
	}, nil

}
