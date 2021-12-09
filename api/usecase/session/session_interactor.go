package usecase

import (
	"fmt"
	"log"
	"net/http"

	"github.com/astaxie/session"
	"github.com/kory-jp/react_golang_api/api/domain"
)

type SessionInteractor struct {
	SessionRepository SessionRepository
}

func (interactor *SessionInteractor) Login(u domain.User, s session.Session) (valid bool, cookie http.Cookie, err error) {
	user, err := interactor.SessionRepository.FindByEmail(u)
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
		return
	}
	if user.Password == u.Encrypt(u.Password) {
		s.Set("uuid", user.UUID)
		cookie = http.Cookie{
			Name:     "_cookie",
			Value:    user.UUID,
			HttpOnly: true,
		}
		valid = true
	} else {
		valid = false
	}
	return
}

func (interactor *SessionInteractor) IsLoggedin(s session.Session, c *http.Cookie) (valid bool) {
	if s.Get("uuid") == c.Value {
		valid = true
	} else {
		valid = false
	}
	return
}

func (interactor *SessionInteractor) Logout(s session.Session, c *http.Cookie) (valid bool, cookie http.Cookie, err error) {
	if s.Get("uuid") == c.Value {
		cookie = http.Cookie{
			Name:     "_cookie",
			Value:    "",
			HttpOnly: true,
		}
		s.Delete("uuid")
		valid = true
	} else {
		valid = false
	}
	fmt.Println(cookie)
	return
}
