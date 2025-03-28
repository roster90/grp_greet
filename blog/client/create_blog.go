package main

import (
	"context"
	"log"

	pb "github.com/roster90/grp_greet/blog/proto"
)

func createBlog(c pb.BlogServiceClient) {
	log.Printf("Starting to do a createBlog RPC...")

	req := &pb.Blog{
		AuthorId: "Roster",
		Title:    "My First Blog",
		Content:  "Content of the first blog",
	}

	res, err := (c).CreateBlog(context.Background(), req)

	if err != nil {
		log.Fatalf("Failed to greet: %v", err)
	}

	log.Printf("Blog has been created: %v", res)
}
