package main

import (
	"io"
	"log"

	pd "github.com/roster90/grp_greet/greet/proto"
)

func (s *Server) GreetEveryone(stream pd.GreetService_GreetEveryoneServer) error {
	log.Printf("GreetEveryOne was invoked with a streaming request\n")

	for {
		req, err := stream.Recv()

		//khi ket thuc stream thi tra ve nil
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Printf("Error while reading client stream: %v\n", err)
			break
		}

		firstName := req.GetFirstName()
		result := "Hello, " + firstName + "!"

		log.Printf("Sending response: %v\n", result)

		stream.Send(&pd.GreetResponse{
			Result: result,
		})
	}

	return nil

}
