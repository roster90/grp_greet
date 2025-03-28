package main

import (
	"context"
	"io"
	"log"

	pb "github.com/roster90/grp_greet/blog/proto"
)

func doGetBlogsFiltered(client pb.BlogServiceClient, cond string) {

	log.Printf("doGetBlogsFiltered to do a doGetBlogsFiltered RPC...")

	filter := &pb.BlogFilter{
		Condition: cond,
	}

	stream, err := client.ListBlogsFiltered(context.Background(), filter)

	if err != nil {
		log.Fatalf("Failed to doGetBlogsFiltered: %v", err)
	}

	for {
		blog, err := stream.Recv()

		if err == io.EOF {
			// we've reached the end of the stream
			break
		}
		if err != nil {
			log.Fatalf("Error while calling ListBlog stream: %v", err)
		}
		log.Printf("Blog: %v", blog)

	}

}
