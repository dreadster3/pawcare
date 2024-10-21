package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PetProfile struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	OwnerId     string             `json:"owner_id" bson:"owner_id"`
	Name        string             `json:"name" bson:"name"`
	DateOfBirth primitive.DateTime `json:"date_of_birth" bson:"date_of_birth"`
	Picture     primitive.Binary   `json:"picture" bson:"picture"`
	Species     string             `json:"species" bson:"species"`
	Breed       string             `json:"breed" bson:"breed"`
	Weight      float64            `json:"weight" bson:"weight"`
	Gender      EGender            `json:"gender" bson:"gender"`
}
