package main

import (
	"context"
	"log"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"

	pd "github.com/roster90/grp_greet/blog/proto"
)

var addr string = "0.0.0.0:50051"

var collection *mongo.Collection

type Server struct {
	pd.BlogServiceServer
}

func main() {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://dev:pQkZzHzuIfV0u1sP@cluster0.xyhwi3x.mongodb.net"))

	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	err = client.Connect(context.Background())

	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	collection = client.Database("db-dev").Collection("blog")

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Listening on %s", addr)

	s := grpc.NewServer()
	pd.RegisterBlogServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
