package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/roster90/grp_greet/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error) {
	log.Printf("CreateBlog function was invoked with %v\n", in)
	data := BlogItem{
		AuthorID: in.AuthorId,
		Content:  in.Content,
		Title:    in.Title,
	}

	res, err := collection.InsertOne(context.Background(), data)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"%s", fmt.Sprintf("Internal Error: %v\n", err),
		)
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			"Cannot convert to OID",
		)
	}

	return &pb.BlogId{Id: oid.Hex()}, nil

}
