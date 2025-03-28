package main

import (
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pd "github.com/roster90/grp_greet/greet/proto"
)

var addr string = "localhost:50051"

type Server struct {
	pd.GreetServiceServer
}

func main() {
	godotenv.Load() // load .env => os.Getenv

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Printf("Listening on %s", addr)

	opts := []grpc.ServerOption{}
	tls := os.Getenv("ENABLE_TLS") == "true" //change to false to disable TLS

	if tls {
		// Create the TLS credentials
		certFile := os.Getenv("TLS_CERT_FILE")
		keyFile := os.Getenv("TLS_KEY_FILE")
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)

		if err != nil {
			log.Fatalf("Failed to load TLS keys: %v", err)
		}
		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)

	pd.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
