package repository

import "github.com/inaohiro/jwt-sample/domain/model"

type User interface {
	FindById(id string) (user *model.User, err error)
	Find(user *model.User) (found *model.User, err error)
	Save(user *model.User) error
	Delete(id string) error
}
