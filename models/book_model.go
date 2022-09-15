package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID       primitive.ObjectID `json:"id,omitempty"`
	Title    string             `json:"title" validate:"required"`
	Author   string             `json:"author" validate:"required"`
	Progress string             `json:"progress" validate:"required"`
	Volume   string             `json:"volume"`
}

//omitempty
