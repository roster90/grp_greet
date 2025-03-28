package main

import (
	"context"
	"io"
	"log"

	pb "github.com/roster90/grp_greet/blog/proto"
)

func doGetBlogsFiltered(client pb.BlogServiceClient) {

	log.Printf("doGetBlogsFiltered to do a doGetBlogsFiltered RPC...")

	filter := &pb.BlogFilter{
		Condition: `{"author_id": "Roster"}`,
	}

	stream, err := client.ListBlogsFilter(context.Background(), filter)

	if err != nil {
		log.Fatalf("Failed to greet: %v", err)
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
