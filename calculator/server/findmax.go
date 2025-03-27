package main

import (
	"io"
	"log"

	pd "github.com/roster90/grp_greet/calculator/proto"
)

func (s *Server) FindMaximum(stream pd.CalculatorService_FindMaximumServer) error {
	log.Printf("FindMaximum was invoked with a streaming request\n")

	max := int32(0)
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

		if number := req.GetNumber(); number > max {
			max = number
			err := stream.Send(&pd.FindMaximumResponse{
				Result: max,
			})

			if err != nil {
				log.Printf("Error while sending data to client: %v\n", err)
			}

		}
	}

	return nil

}
