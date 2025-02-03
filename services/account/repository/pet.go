package repository

import (
	"github.com/dreadster3/pawcare/services/account/aggregate"
)

type IPetRepository interface {
	FindByOwnerId(ownerId aggregate.OwnerId) ([]*aggregate.Pet, error)
}
