package main

import (
	"fmt"
	"log"

	pd "github.com/roster90/grp_greet/greet/proto"
)

func (s *Server) GreetManyTimes(req *pd.GreetRequest, stream pd.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes function was invoked with a streaming request\n")

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello, %s, number %d", req.FirstName, i)
		stream.Send(&pd.GreetResponse{
			Result: res,
		})
	}
	return nil

}
