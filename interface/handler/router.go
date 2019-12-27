package handler

import (
	"net/http"

	"github.com/inaohiro/go-multiplexer"
	"github.com/inaohiro/jwt-sample/interface/persistence"
)

func NewRouter(user User) http.Handler {
	mux := multiplexer.NewMultiplexer()
	mux.POST("/users", user.Create)
	mux.DELETE("/users", user.Delete)

	mux.POST("/token", user.Auth)

	//mux.GET("/key")
	mux.GET("/", func(writer http.ResponseWriter, r *http.Request) {
		persistence.Dump()
	})

	return mux
}
