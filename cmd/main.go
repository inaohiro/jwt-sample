package main

import (
	"fmt"
	"log"

	"github.com/inaohiro/jwt-sample/domain/model"
	"github.com/inaohiro/jwt-sample/domain/service"
	"github.com/inaohiro/jwt-sample/interface/persistence"
	"github.com/inaohiro/jwt-sample/usecase"
)

func main() {
	//dbenv, err := config.NewDBConfig()
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	//db, cleanup, err := persistence.NewConnection(dbenv)
	//defer cleanup()
	//userRepository := persistence.NewUserRepository(db)

	userRepository := persistence.NewMockUserRepository()
	userService := service.NewUserService(userRepository)
	u := usecase.NewUserUsecase(userService)

	// create user
	user, err := u.Create("root", "password")
	if err != nil {
		log.Fatal(err.Error())
	}
	persistence.Dump()

	// get token
	token, err := u.Authenticate(user.Id(), user.Password())
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(token.Token())

	// validate token
	if err := model.Validate(token); err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("ok")
}
