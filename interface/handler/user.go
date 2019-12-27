package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/inaohiro/jwt-sample/usecase"
)

type User interface {
	Create(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Auth(w http.ResponseWriter, r *http.Request)
}

// guarantee the userHandler struct implements all of the method of the User interface
var _ User = (*userHandler)(nil)

type userHandler struct {
	uc usecase.User
}

func NewUserHandler(usecase usecase.User) *userHandler {
	return &userHandler{uc: usecase}
}

type userRequest struct {
	Id       string
	Password string
}

func (user *userHandler) Create(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Printf("close request body: %v", err)
		}
	}()
	log.Println(string(b))

	var m userRequest
	if err := json.Unmarshal(b, &m); err != nil {
		// log tracable error
		log.Printf("unmarshal json: %v", err.Error())
		// get radical errors and return the short message
		http.Error(w, fmt.Sprintf("unmarshaling json: %v", err.Error()), 500)
		return
	}
	log.Println(m)

	_, err = user.uc.Create(m.Id, m.Password)
	if err != nil {
		log.Printf("create user: %v", err.Error())
		http.Error(w, fmt.Sprintf("creating user: %v", err.Error()), 500)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (user *userHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/users")

	if err := user.uc.Delete(id); err != nil {
		// return error
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (user *userHandler) Auth(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	defer func() {
		if err := r.Body.Close(); err != nil {
			// logging ?
		}
	}()

	var m userRequest
	if err := json.Unmarshal(b, &m); err != nil {
		log.Printf("unmarshal json: %v", err.Error())
	}

	token, err := user.uc.Authenticate(m.Id, m.Password)
	if err != nil {
		log.Printf("authenticated user: %v", err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(map[string]string{"token": token.Token()})
	_, _ = w.Write(response)

}
