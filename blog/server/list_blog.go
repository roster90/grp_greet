package main

import (
	"context"
	"log"

	pb "github.com/roster90/grp_greet/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) ListBlogs(req *emptypb.Empty, stream pb.BlogService_ListBlogsFilterServer) error {
	log.Printf("ListBlogs function was invoked with %v\n", req)

	if collection == nil {
		log.Fatal("🔥 MongoDB collection is nil!")
	}

	cursor, err := collection.Find(context.Background(), primitive.D{})

	if err != nil {
		return status.Errorf(codes.Internal, "Unknown internal error %v", err)
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var blog BlogItem
		if err := cursor.Decode(&blog); err != nil {
			return status.Errorf(codes.Internal, "Decode error: %v", err)
		}

		res := documentToBlog(&blog)

		if err := stream.Send(res); err != nil {
			return status.Errorf(codes.Internal, "Stream send error: %v", err)
		}

		stream.Send(documentToBlog(&blog))
	}

	if err := cursor.Err(); err != nil {
		return status.Errorf(codes.Internal, "Unknown internal error %v", err)
	}
	return nil
}
