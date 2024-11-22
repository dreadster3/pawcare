package entity

import "time"

type OwnerId string

type Owner struct {
	Id          OwnerId
	Name        string
	DateOfBirth time.Time
}
