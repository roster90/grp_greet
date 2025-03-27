package main

import (
	"context"
	"fmt"
	"math"

	"log"

	pd "github.com/roster90/grp_greet/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Sqrt(ctx context.Context, req *pd.SqrtRequest) (*pd.SqrtResponse, error) {
	log.Printf(" Sqrt function was invoked with %v\n", req)
	number := req.GetNumber()

	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"%s", fmt.Sprintf("Received a negative number: %v", number),
		)
	}

	return &pd.SqrtResponse{
		Result: float32(math.Sqrt(float64(number))),
	}, nil
}
