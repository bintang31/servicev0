package persistence

import (
	"errors"
	"servicev0/src/domain/entity"
	"servicev0/src/domain/repository"
	"strings"

	"gorm.io/gorm"
)

//UserRepo : Call DB
type UserRepo struct {
	db *gorm.DB
}

//NewUserRepository : User Repository
func NewUserRepository(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}

//UserRepo implements the repository.UserRepository interface
var _ repository.UserRepository = &UserRepo{}

//SaveUser : Save User to DB
func (r *UserRepo) SaveUser(user *entity.User) (*entity.User, map[string]string) {
	dbErr := map[string]string{}
	err := r.db.Debug().Create(&user).Error
	if err != nil {
		//If the email is already taken
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			dbErr["email_taken"] = "email already taken"
			return nil, dbErr
		}
		//any other db error
		dbErr["db_error"] = "database error"
		return nil, dbErr
	}
	return user, nil
}

//GetUser : Get User Detail from DB
func (r *UserRepo) GetUsers() ([]entity.User, error) {
	var user []entity.User
	err := r.db.Debug().Take(&user).Error
	if err != nil {
		return nil, err
	}
	errors.Is(err, gorm.ErrRecordNotFound)
	if err != nil {
		return nil, err
	}
	return user, nil
}
