package persistence

import (
	"fmt"
	"time"

	pgv2 "gorm.io/driver/postgres"
	v2 "gorm.io/gorm"

	//Gorm POSTGRES

	"servicev0/src/domain/repository"
)

//Repositories : Assign Repository
type Repositories struct {
	User repository.UserRepository
	dbv2 *v2.DB
}

//NewRepositories : Register Repository
func NewRepositories(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) (*Repositories, error) {
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s application_name=%s search_path=%s", DbHost, DbPort, DbUser, DbName, DbPassword, "mobileloket-engine", "public")
	dbv2, err := v2.Open(pgv2.Open(DBURL), &v2.Config{})
	if err != nil {
		return nil, err
	}
	sqlDB, err := dbv2.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(200)

	sqlDB.SetMaxIdleConns(100)

	sqlDB.SetConnMaxLifetime(2 * time.Minute)

	return &Repositories{
		User: NewUserRepository(dbv2),
		dbv2: dbv2,
	}, nil
}

//Close : closes the  database connection
func (s *Repositories) CloseEngine() error {
	sqlDB, err := s.dbv2.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
