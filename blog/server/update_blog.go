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
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) UpdateBlog(ctx context.Context, in *pb.Blog) (*emptypb.Empty, error) {
	log.Printf("UpdateBlog function was invoked with %v\n", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"%s", fmt.Sprintf("Cannot parseUD: %v\n", err),
		)
	}
	data := &BlogItem{
		AuthorID: in.AuthorId,
		Content:  in.Content,
		Title:    in.Title,
	}

	res, err := collection.UpdateOne(context.Background(), bson.M{"_id": oid}, bson.M{"$set": data})

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"%s", fmt.Sprintf("Could not update  %v\n", err),
		)
	}

	if res.ModifiedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"%s", fmt.Sprintf("Cannot find blog with specified ID: %v\n", err),
		)
	}

	return &emptypb.Empty{}, nil

}
