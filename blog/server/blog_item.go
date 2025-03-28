package main

import (
	pd("github.com/roster90/grp_greet/blog/proto")
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogItem struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string `bson:"author_id,omitempty"`
	Content string `bson:"content,omitempty"`
	Title string `bson:"title,omitempty"`
}