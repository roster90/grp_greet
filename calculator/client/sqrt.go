package main

import (
	"context"

	"log"

	pb "github.com/roster90/grp_greet/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doSqrt(c pb.CalculatorServiceClient, n int32) {

	req := &pb.SqrtRequest{
		Number: n,
	}

	res, err := c.Sqrt(context.Background(), req)

	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			log.Printf("Error message from server: %v", e.Message())
			if e.Code() == codes.InvalidArgument {
				log.Printf("We probably sent a negative number %v!", n)
				return
			}
		} else {
			log.Fatalf("Anone  while calling Sqrt: %v", e)
		}

	}

	log.Printf("Response from Sqrt: %v", res.GetResult())

}
