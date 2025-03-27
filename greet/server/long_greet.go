package main

import (
	"io"
	"log"

	pd "github.com/roster90/grp_greet/greet/proto"
)

func (s *Server) LongGreet(streamRq pd.GreetService_LongGreetServer) error {
	log.Printf("LongGreet function was invoked with a streaming request\n")

	result := ""
	for {
		req, err := streamRq.Recv()

		if err == io.EOF {
			return streamRq.SendAndClose(&pd.GreetResponse{
				Result: result,
			})
		}

		if err != nil {
			log.Printf("Error while reading client stream: %v\n", err)
			break
		}

		result += "Hello, " + req.GetFirstName() + "!\n"
		log.Printf("Sending response: %v\n", result)

	}
	return nil

}
