package controllers

import (
	"fmt"
	"net/http"

	"github.com/kory-jp/react_golang_api/api/interfaces/database"
	"github.com/kory-jp/react_golang_api/api/usecase"
)

// func GetUser(w http.ResponseWriter, r *http.Request) {

// 	t := struct {
// 		Title  string `json:"title"`
// 		Number int    `json:"number"`
// 	}{
// 		Title:  "test",
// 		Number: 2,
// 	}
// 	s, err := json.Marshal(t)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(string(s))
// }

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

// func (controller *UserController) Index(c Context) {
func (controller *UserController) Index(W http.ResponseWriter, r *http.Request) {
	users, err := controller.Interfactor.Users()
	if err != nil {
		// c.JSON(500, NewError(err))
		return
	}
	// c.JSON(200, users)
	fmt.Println(users)
}
