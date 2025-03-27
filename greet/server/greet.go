package main

import (
	"context"
	"log"

	pd "github.com/roster90/grp_greet/greet/proto"
)

func (s *Server) Greet(ctx context.Context, req *pd.GreetRequest) (*pd.GreetResponse, error) {
	log.Printf(" Greet function was invoked with %v\n", req)
	return &pd.GreetResponse{
		Result: "Hello, " + req.FirstName}, nil
}
