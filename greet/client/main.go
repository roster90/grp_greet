package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	pb "github.com/roster90/grp_greet/greet/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr string = "localhost:50051"

func main() {
	godotenv.Load() // load .env => os.Getenv

	tls := os.Getenv("ENABLE_TLS") == "true" //change to false to disable TLS
	certFile := os.Getenv("TLS_CA_FILE")
	log.Printf("Loading TLS keys from %s", certFile)

	opts := []grpc.DialOption{}

	if tls {

		creds, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			log.Fatalf("Failed to load TLS keys: %v\n", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))

	}
	conn, err := grpc.NewClient(addr, opts...)

	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	log.Printf("Connected to %s", addr)

	defer conn.Close()

	cs := pb.NewGreetServiceClient(conn)

	doGreet(cs)
	// doGreetManyTime(cs)
	// doLongGreet(cs)
	// doGreetEveryone(cs)
	// doGreetWithDeadline(cs, 5*time.Second)
}
