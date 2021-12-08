package usecase

import (
	"github.com/astaxie/session"
)

type SessionInteractor struct {
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
