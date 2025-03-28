package main

import (
	"context"
	"log"

	pb "github.com/roster90/grp_greet/blog/proto"
)

func doDeleteBlog(c pb.BlogServiceClient, id string) {
	log.Printf("Starting to do a doDeleteBlog ")

	req := &pb.BlogId{
		Id: id,
	}
	_, err := c.DeleteBlog(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to delete Blog: %v", err)
	}
	log.Printf("Blog has been deleted successfully")
}
