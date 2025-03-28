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
		Title:    "My second Blog",
		Content:  "Content of the second blog",
	}

	res, err := (c).CreateBlog(context.Background(), req)

	if err != nil {
		log.Fatalf("Failed to greet: %v", err)
	}

	log.Printf("Blog has been created: %v", res)
}
