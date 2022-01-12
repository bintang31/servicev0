package repository

import (
	"servicev0/src/domain/entity"
)

//UserRepository : User collection of methods that the infrastructure
type UserRepository interface {
	SaveUser(*entity.User) (*entity.User, map[string]string)
	GetUsers() ([]entity.User, error)
}
