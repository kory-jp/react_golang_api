package usecase

import "github.com/kory-jp/react_golang_api/api/domain"

type SessionRepository interface {
	FindByEmail(domain.User) (domain.User, error)
}
