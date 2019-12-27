package persistence

import (
	"log"

	"github.com/inaohiro/jwt-sample/codes"
	"github.com/inaohiro/jwt-sample/domain/model"
	"github.com/inaohiro/jwt-sample/domain/repository"
)

//var store = map[string]*model.User{}
var store = make(map[string]*model.User)

type mockUserRepository struct {
}

func NewMockUserRepository() repository.User {
	return &mockUserRepository{}
}

func (m *mockUserRepository) FindById(id string) (user *model.User, err error) {
	if found, ok := store[id]; ok {
		return found, nil
	}
	return nil, codes.NotFound
}

func (m *mockUserRepository) Find(user *model.User) (found *model.User, err error) {
	found, ok := store[user.Id()]
	if !ok {
		return nil, codes.NotFound
	}
	if found.Password() != user.Password() {
		return nil, codes.InvalidEntity
	}

	return
}

func (m *mockUserRepository) Save(user *model.User) error {
	store[user.Id()] = user
	return nil
}

func (m *mockUserRepository) Delete(id string) error {
	delete(store, id)
	return nil
}

func Dump() {
	for _, v := range store {
		log.Println(v)
	}
}
