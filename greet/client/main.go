package main

import (
	"log"
	"time"

	pb "github.com/roster90/grp_greet/greet/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "0.0.0.0:50051"

func main() {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	log.Printf("Connected to %s", addr)

	defer conn.Close()

	cs := pb.NewGreetServiceClient(conn)

	// doGreet(cs)
	// doGreetManyTime(cs)
	// doLongGreet(cs)
	// doGreetEveryone(cs)
	doGreetWithDeadline(cs, 5*time.Second)
}
