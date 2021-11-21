package usecase

import (
	"fmt"
	"log"

	"github.com/kory-jp/react_golang_api/api/domain"
)

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Add(u domain.User) (user domain.User, err error) {
	identifier, err := interactor.UserRepository.Store(u)
	if err != nil {
		return
	}
	user, err = interactor.UserRepository.FindById(identifier)
	return
}

func (interactor *UserInteractor) Users() (user domain.Users, err error) {
	user, err = interactor.UserRepository.FindAll()
	return
}

func (interactor *UserInteractor) UserById(identifier int) (user domain.User, err error) {
	user, err = interactor.UserRepository.FindById(identifier)
	return
}

func (interactor *UserInteractor) UpdateUser(identifier int, u domain.User) (user domain.User, err error) {
	updateUseridentifier, err := interactor.UserRepository.Update(identifier, u)
	if err != nil {
		fmt.Println(err)
		return
	}
	user, err = interactor.UserRepository.FindById(updateUseridentifier)
	return
}

func (interactor *UserInteractor) DeleteUser(identifier int) (err error) {
	if err := interactor.UserRepository.Delete(identifier); err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
	}
	return
}
