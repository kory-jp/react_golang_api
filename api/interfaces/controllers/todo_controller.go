package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/kory-jp/react_golang_api/api/domain"
	"github.com/kory-jp/react_golang_api/api/interfaces/database"
	"github.com/kory-jp/react_golang_api/api/usecase"
)

type TodoController struct {
	Interactor usecase.TodoInteractor
}

func NewTodoController(sqlHandler database.SqlHandler) *TodoController {
	return &TodoController{
		Interactor: usecase.TodoInteractor{
			TodoRepository: &database.TodoRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *TodoController) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test1")
	bytesTodo, err := io.ReadAll(r.Body)
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
	}
	todoType := new(domain.Todo)
	if err := json.Unmarshal(bytesTodo, todoType); err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
		return
	}
	todo, err := controller.Interactor.Add(*todoType)
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
	}
	fmt.Println(todo)
	fmt.Fprintln(w, todo)
}

func (controller *TodoController) Show(w http.ResponseWriter, r *http.Request, id int) {
	todo, err := controller.Interactor.TodoById(id)
	if err != nil {
		log.Println(err)
		return
	}
	jsonTodo, err := json.Marshal(todo)
	if err != nil {
		log.Println(err)
	}
	fmt.Fprintln(w, string(jsonTodo))
}

func (controller *TodoController) Update(w http.ResponseWriter, r *http.Request, id int) {
	bytesTodo, err := io.ReadAll(r.Body)
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
		return
	}
	todoType := new(domain.Todo)
	if err := json.Unmarshal(bytesTodo, todoType); err != nil {
		fmt.Printf("%#v\n", err)
		log.Fatalln(err)
		return
	}
	todo, err := controller.Interactor.UpdateTodo(id, *todoType)
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
	}
	jsonTodo, err := json.Marshal(todo)
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
	}
	fmt.Fprintf(w, string(jsonTodo))
}
