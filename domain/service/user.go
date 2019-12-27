package service

import (
	"errors"
	"fmt"

	"github.com/inaohiro/jwt-sample/codes"
	"github.com/inaohiro/jwt-sample/domain/model"
	"github.com/inaohiro/jwt-sample/domain/repository"
)

type User interface {
	Add(id, password string) (*model.User, error)
	Delete(id string) error
	Authenticate(id, password string) (*model.Token, error)
}

type userService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *userService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) IsDuplicated(id string) error {
	user, err := s.repo.FindById(id)
	if err != nil && !errors.Is(err, codes.NotFound) {
		return err
	}
	if user != nil {
		return errors.New("duplicated user")
	}

	return nil
}

func (s *userService) Add(id string, password string) (*model.User, error) {
	if err := s.IsDuplicated(id); err != nil {
		return nil, err
	}
	user, err := model.NewUser(id, password)
	if err != nil {
		return nil, fmt.Errorf("initialize user model: %v", err)
	}
	if err := s.repo.Save(user); err != nil {
		return nil, fmt.Errorf("save user: %v", err)
	}

	return user, nil
}

func (s *userService) Delete(id string) error {
	// check user exists
	user, err := s.repo.FindById(id)
	if err != nil {
		return err
	}
	if user == nil {
		return fmt.Errorf("could not find the user: %s", id)
	}

	if err := s.repo.Delete(id); err != nil {
		return err
	}

	return nil
}

// pass user instance ?
func (s *userService) Authenticate(id string, password string) (*model.Token, error) {
	user, err := model.NewUser(id, password)
	if err != nil {
		return nil, err
	}
	found, err := s.repo.Find(user)
	if err != nil {
		return nil, err
	}
	if found == nil {
		return nil, errors.New("invalid user id or password")
	}

	// generate token
	token, err := model.NewToken()
	if err != nil {
		return nil, err
	}
	return token, nil
}
