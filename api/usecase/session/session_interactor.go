package usecase

import (
	"log"

	"github.com/astaxie/session"
	"github.com/kory-jp/react_golang_api/api/domain"
)

type SessionInteractor struct {
	SessionRepository SessionRepository
}

func (interactor *SessionInteractor) Count(sess session.Session) (coutup_sess session.Session) {
	ct := sess.Get("countnum")
	if ct == nil {
		sess.Set("countnum", 1)
	} else {
		sess.Set("countnum", (ct.(int) + 1))
	}
	coutup_sess = sess
	return
}

func (interactor *SessionInteractor) Login(u domain.User, s session.Session) (valid bool, err error) {
	user, err := interactor.SessionRepository.FindByEmail(u)
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
		return
	}
	if user.Password == u.Encrypt(u.Password) {
		s.Set("uuid", user.UUID)
		// fmt.Println(s.Get("uuid"))
		valid = true
	} else {
		valid = false
	}
	return
}
