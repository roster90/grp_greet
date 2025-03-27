package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pd "github.com/roster90/grp_greet/calculator/proto"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pd.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Listening on %s", addr)

	s := grpc.NewServer()
	pd.RegisterCalculatorServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
