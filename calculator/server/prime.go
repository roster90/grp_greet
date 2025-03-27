package main

import (
	"log"

	pd "github.com/roster90/grp_greet/calculator/proto"
)

func (s *Server) Primes(req *pd.PrimesRequest, stream pd.CalculatorService_PrimesServer) error {
	log.Printf(" Greet function was invoked with %v\n", req)

	N := req.Number
	k := int64(2)

	for N > 1 { // giống như while N > 1

		if N%k == 0 {
			stream.Send(&pd.PrimesResponse{
				Result: k,
			})
			N = N / k
		} else {
			k++
		}
	}

	return nil
}
