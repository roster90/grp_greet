package main

import (
	"io"
	"log"

	pd "github.com/roster90/grp_greet/calculator/proto"
)

func (s *Server) ComputeAverage(streamRq pd.CalculatorService_ComputeAverageServer) error {
	log.Printf("ComputeAverage function was invoked with a streaming request\n")

	sum := float32(0.0)
	count := 0
	for {
		req, err := streamRq.Recv()

		if err == io.EOF {
			return streamRq.SendAndClose(&pd.ComputeAverageResponse{
				Result: sum / float32(count),
			})
		}

		count++
		sum += req.GetNumber()

		log.Printf("Received number: %v\n", req.GetNumber())
		if err != nil {
			log.Printf("Error while reading client stream: %v\n", err)
			break
		}

	}
	return nil

}
