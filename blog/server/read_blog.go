package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/roster90/grp_greet/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ReadBlog(ctx context.Context, in *pb.BlogId) (*pb.Blog, error) {

	log.Printf("ReadBlog function was invoked with blogId %v\n", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"%s", fmt.Sprintf("Internal Error: %v\n", err),
		)
	}
	data := &BlogItem{}
	filter := bson.M{"_id": oid}
	res := collection.FindOne(context.Background(), filter)

	if condition := res.Decode(data); condition != nil {
		return nil, status.Errorf(
			codes.NotFound,
			"%s", fmt.Sprintf("Cannot find blog with specified ID: %v\n", err),
		)

	}

	return documentToBlog(data), nil
}
