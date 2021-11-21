package usecase

import "github.com/kory-jp/react_golang_api/api/domain"

type UserRepository interface {
	Store(domain.User) (int, error)
	FindById(int) (domain.User, error)
	FindAll() (domain.Users, error)
	Update(int, domain.User) (int, error)
	Delete(int) error
}
