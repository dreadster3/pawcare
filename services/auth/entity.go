package auth

import "fmt"

type UserId string

type User struct {
	Id UserId
}

func (u User) String() string {
	return fmt.Sprintf("User(id=%s)", u.Id)
}
