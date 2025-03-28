package main

import (
	"context"
	"log"

	pb "github.com/roster90/grp_greet/blog/proto"
)

func doUpdateBlog(c pb.BlogServiceClient, id string) {
	log.Printf("Starting to do a doUpdateBlog ")

	req := &pb.Blog{
		Id:       id,
		AuthorId: "Roster",
		Title:    "My First Blog (edited2)",
		Content:  "Content of the first blog, with some awesome additions!2",
	}
	_, err := c.UpdateBlog(context.Background(), req)

	if err != nil {
		log.Fatalf("Failed to update Blog: %v", err)
	}
	log.Printf("Blog has been updated successfully")
}
