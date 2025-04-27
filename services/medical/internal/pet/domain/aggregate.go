package domain

type (
	PetId  string
	UserId string
)

type Pet struct {
	Id     PetId
	UserId UserId
}
