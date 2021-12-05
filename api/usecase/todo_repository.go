package usecase

import "github.com/kory-jp/react_golang_api/api/domain"

type TodoRepository interface {
	Store(domain.Todo) (int, error)
	FindById(int) (domain.Todo, error)
	FindAll() (domain.Todos, error)
	Update(int, domain.Todo) (int, error)
	Delete(int) error
}
