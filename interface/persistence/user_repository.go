package persistence

import (
	"github.com/inaohiro/jwt-sample/domain/model"
	"github.com/inaohiro/jwt-sample/domain/repository"
)

type userRepository struct {
	db *DB
}

func NewUserRepository(db *DB) repository.User {
	return &userRepository{db: db}
}

func (repo *userRepository) FindById(id string) (user *model.User, err error) {
	panic("implement me")
}

func (repo *userRepository) Find(user *model.User) (found *model.User, err error) {
	panic("implement me")
}

func (repo *userRepository) Save(user *model.User) error {
	panic("implement me")
}

func (repo *userRepository) Delete(id string) error {
	panic("implement me")
}
