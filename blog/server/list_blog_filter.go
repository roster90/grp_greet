package main

import (
	"context"
	"encoding/json"

	pb "github.com/roster90/grp_greet/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ListBlogsFiltered(req *pb.BlogFilter, stream pb.BlogService_ListBlogsFilterServer) error {
	var filter bson.M

	if err := json.Unmarshal([]byte(req.Condition), &filter); err != nil {
		return status.Errorf(codes.InvalidArgument, "Invalid filter condition: %v", err)
	}

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return status.Errorf(codes.Internal, "Failed to find blogs: %v", err)
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var blog BlogItem
		if err := cursor.Decode(&blog); err != nil {
			return status.Errorf(codes.Internal, "Decode error: %v", err)
		}
		stream.Send(documentToBlog(&blog))
	}
	if err := cursor.Err(); err != nil {
		return status.Errorf(codes.Internal, "Unknown internal error %v", err)
	}

	return nil
}
