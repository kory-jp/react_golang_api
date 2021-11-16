package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/kory-jp/react_golang_api/api/interfaces/database"
	"github.com/kory-jp/react_golang_api/api/usecase"
)

type UserController struct {
	Interfactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Interfactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) Index(w http.ResponseWriter, r *http.Request) {
	users, err := controller.Interfactor.Users()
	if err != nil {
		log.Panicln(err)
		return
	}
	us, err := json.Marshal(users)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Fprintf(w, string(us))
}
