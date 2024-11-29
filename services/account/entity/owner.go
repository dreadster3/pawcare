package entity

import (
	"fmt"
	"time"
)

type OwnerId string

type Owner struct {
	Id          OwnerId
	Name        string
	DateOfBirth time.Time
}

func (o Owner) String() string {
	return fmt.Sprintf("Owner(id=%s,name=%s,dob=%s)", o.Id, o.Name, o.DateOfBirth.String())
}
