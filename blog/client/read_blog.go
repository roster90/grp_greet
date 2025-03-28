package main

import (
	"context"
	"log"

	pb "github.com/roster90/grp_greet/blog/proto"
)

func doReadBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Printf("Starting to do a doReadBlog ")

	req := &pb.BlogId{
		Id: id,
	}
	res, err := c.ReadBlog(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to greet: %v", err)
	}
	log.Printf("Blog has been read: %v", res)
	return res
}
