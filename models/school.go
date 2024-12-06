package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type School struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name" validate:"required"`
	Address     string             `json:"address" bson:"address" validate:"required"`
	Contact     string             `json:"contact" bson:"contact" validate:"required"`
	Email       string             `json:"email" bson:"email" validate:"required"`
	Description string             `json:"description" bson:"description" validate:"max=255"`
}