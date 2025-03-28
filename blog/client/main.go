package main

import (
	"log"

	pb "github.com/roster90/grp_greet/blog/proto"

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

	cs := pb.NewBlogServiceClient(conn)

	// createBlog(cs)

	// doUpdateBlog(cs, "67e65095f1adcb156b922e8a")
	// doReadBlog(cs, "67e65095f1adcb156b922e8a")
	// doGetBlogs(cs)
	doGetBlogsFiltered(cs)
}
