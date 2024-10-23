package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Pet struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	OwnerId     string             `json:"-" bson:"owner_id" swaggerignore:"true"`
	Name        string             `json:"name" bson:"name"`
	DateOfBirth primitive.DateTime `json:"date_of_birth" bson:"date_of_birth"`
	Picture     primitive.Binary   `json:"picture" bson:"picture"`
	Species     string             `json:"species" bson:"species"`
	Breed       string             `json:"breed" bson:"breed"`
	Weight      float64            `json:"weight" bson:"weight"`
	Gender      EGender            `json:"gender" bson:"gender"`
}

type Owner struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	UserId      string             `json:"-" bson:"user_id" swaggerignore:"true"`
	Name        string             `json:"name" bson:"name"`
	DateOfBirth primitive.DateTime `json:"date_of_birth" bson:"date_of_birth"`
	Picture     primitive.Binary   `json:"picture" bson:"picture"`
}
