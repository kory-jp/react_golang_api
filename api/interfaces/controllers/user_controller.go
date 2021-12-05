package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/kory-jp/react_golang_api/api/domain"
	"github.com/kory-jp/react_golang_api/api/interfaces/database"
	usecase "github.com/kory-jp/react_golang_api/api/usecase/user"
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

func (controller *UserController) Create(w http.ResponseWriter, r *http.Request) {
	bytesUser, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
		return
	}
	userType := new(domain.User)
	if err := json.Unmarshal(bytesUser, userType); err != nil {
		fmt.Printf("%#v\n", err)
		log.Fatalln(err)
		return
	}
	user, err := controller.Interfactor.Add(*userType)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(user)
	fmt.Fprintln(w, user)
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

func (controller *UserController) Show(w http.ResponseWriter, r *http.Request, id int) {
	user, err := controller.Interfactor.UserById(id)
	if err != nil {
		log.Println(err)
		return
	}
	jsonUser, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
	}
	fmt.Fprintf(w, string(jsonUser))
}

func (controller *UserController) Update(w http.ResponseWriter, r *http.Request, id int) {
	bytesUser, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
		return
	}
	userType := new(domain.User)
	if err := json.Unmarshal(bytesUser, userType); err != nil {
		fmt.Printf("%#v\n", err)
		log.Fatalln(err)
		return
	}
	user, err := controller.Interfactor.UpdateUser(id, *userType)
	if err != nil {
		log.Println(err)
	}
	jsonUser, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
	}
	fmt.Fprintf(w, string(jsonUser))
}

func (controller *UserController) Delete(w http.ResponseWriter, r *http.Request, id int) {
	if err := controller.Interfactor.DeleteUser(id); err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
	} else {
		fmt.Fprintf(w, "削除成功")
	}
}
