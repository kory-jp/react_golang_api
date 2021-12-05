package usecase

import (
	"log"

	"github.com/kory-jp/react_golang_api/api/domain"
)

type TodoInteractor struct {
	TodoRepository TodoRepository
}

func (interactor *TodoInteractor) Add(t domain.Todo) (todo domain.Todo, err error) {
	identifier, err := interactor.TodoRepository.Store(t)
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
		return
	}
	todo, err = interactor.TodoRepository.FindById(identifier)
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
		return
	}
	return
}

func (interactor *TodoInteractor) Todos() (todos domain.Todos, err error) {
	todos, err = interactor.TodoRepository.FindAll()
	return
}

func (interactor *TodoInteractor) TodoById(indentifier int) (todo domain.Todo, err error) {
	todo, err = interactor.TodoRepository.FindById(indentifier)
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
		return
	}
	return
}

func (interactor *TodoInteractor) UpdateTodo(indentifier int, t domain.Todo) (todo domain.Todo, err error) {
	updateTodoIdentifier, err := interactor.TodoRepository.Update(indentifier, t)
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
		return
	}
	todo, err = interactor.TodoRepository.FindById(updateTodoIdentifier)
	return
}

func (interactor *TodoInteractor) DeleteTodo(indentifier int) (err error) {
	if err := interactor.TodoRepository.Delete(indentifier); err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
	}
	return
}
