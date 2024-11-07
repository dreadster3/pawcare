package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Record struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	UserId      string             `json:"-" bson:"user_id"`
	PetId       string             `json:"pet_id" bson:"pet_id"`
	Type        RecordType         `json:"type" bson:"type"`
	Description string             `json:"description" bson:"description"`
	Date        primitive.DateTime `json:"date" bson:"date"`
}

type RecordType string

const (
	RecordTypeVaccination RecordType = "vaccination"
	RecordTypeTreatment   RecordType = "treatment"
	RecordTypeDeworming   RecordType = "deworming"
	RecordTypeSurgery     RecordType = "surgery"
	RecordTypeCheckUp     RecordType = "checkup"
	RecordTypeOther       RecordType = "other"
)
