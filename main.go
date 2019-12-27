package main

import (
	"log"

	"github.com/inaohiro/jwt-sample/config"
	"github.com/inaohiro/jwt-sample/domain/service"
	"github.com/inaohiro/jwt-sample/infra"
	"github.com/inaohiro/jwt-sample/interface/handler"
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
	userUsecase := usecase.NewUserUsecase(userService)

	userHandler := handler.NewUserHandler(userUsecase)
	router := handler.NewRouter(userHandler)
	serverEnv := config.ServerEnv{
		Addr: "localhost",
		Port: 8080,
	}
	httpServer := infra.NewHttpServer(&serverEnv, router)

	log.Fatal(httpServer.ListenAndServe())
}
