package usecase

import (
	"github.com/inaohiro/jwt-sample/domain/model"
	"github.com/inaohiro/jwt-sample/domain/service"
)

type User interface {
	Create(id, password string) (*model.User, error)
	Delete(id string) error
	Authenticate(id, password string) (*model.Token, error)
}

type userUsecase struct {
	svc service.User
}

func NewUserUsecase(svc service.User) *userUsecase {
	return &userUsecase{svc: svc}
}

func (u *userUsecase) Create(id, password string) (*model.User, error) {
	user, err := u.svc.Add(id, password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUsecase) Delete(id string) error {
	if err := u.svc.Delete(id); err != nil {
		return err
	}
	return nil
}

func (u *userUsecase) Authenticate(id, password string) (*model.Token, error) {
	token, err := u.svc.Authenticate(id, password)
	if err != nil {
		return nil, err
	}

	return token, nil
}
