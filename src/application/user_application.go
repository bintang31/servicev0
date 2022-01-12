package application

import (
	"servicev0/src/domain/entity"
	"servicev0/src/domain/repository"
)

type userApp struct {
	us repository.UserRepository
}

//UserApp implements the UserAppInterface
var _ UserAppInterface = &userApp{}

//UserAppInterface : Interfacing User App to Repository
type UserAppInterface interface {
	SaveUser(*entity.User) (*entity.User, map[string]string)
	GetUsers() ([]entity.User, error)
}

func (u *userApp) SaveUser(user *entity.User) (*entity.User, map[string]string) {
	return u.us.SaveUser(user)
}

func (u *userApp) GetUsers() ([]entity.User, error) {
	return u.us.GetUsers()
}
