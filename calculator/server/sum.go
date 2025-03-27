package main

import (
	"context"
	"log"

	pd "github.com/roster90/grp_greet/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, req *pd.CalculatorRequest) (*pd.CalculatorResponse, error) {
	log.Printf(" Greet function was invoked with %v\n", req)

	return &pd.CalculatorResponse{
		Result: req.Number1 + req.Number2,
	}, nil
}
