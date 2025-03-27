package main

import (
	"io"
	"log"

	pd "github.com/roster90/grp_greet/calculator/proto"
)

func (s *Server) ComputeAverage(streamRq pd.CalculatorService_ComputeAverageServer) error {
	log.Printf("ComputeAverage function was invoked with a streaming request\n")

	total := float32(0.0)
	counter := 0
	for {
		req, err := streamRq.Recv()

		if err == io.EOF {
			return streamRq.SendAndClose(&pd.ComputeAverageResponse{
				Result: total / float32(counter),
			})
		}

		counter++
		total += req.GetNumber()

		log.Printf("Received number: %v\n", req.GetNumber())
		if err != nil {
			log.Printf("Error while reading client stream: %v\n", err)
			break
		}

	}
	return nil

}
