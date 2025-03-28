package main

import (
	"context"
	"io"
	"log"

	pb "github.com/roster90/grp_greet/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func doGetBlogs(c pb.BlogServiceClient) {
	log.Printf("doGetListBlog to do a ListBlog RPC...")

	stream, err := c.ListBlogs(context.Background(), &emptypb.Empty{})

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
