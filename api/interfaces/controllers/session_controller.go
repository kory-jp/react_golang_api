package controllers

import (
	"fmt"
	"net/http"

	"github.com/astaxie/session"
	_ "github.com/astaxie/session/providers/memory"
	usecase "github.com/kory-jp/react_golang_api/api/usecase/session"
)

// func Cookie(w http.ResponseWriter, r *http.Request) {
// 	expiration := time.Now()
// 	expiration = expiration.AddDate(0, 0, 1)
// 	cookie := http.Cookie{Name: "username", Value: "golang", Expires: expiration}
// 	http.SetCookie(w, &cookie)
// 	for _, c := range r.Cookies() {
// 		log.Print("Name", c.Name, "Value", c.Value)
// 	}
// }

var GlobalSessions *session.Manager

type SessionController struct {
	Interactor usecase.SessionInteractor
}

func NewManager() *SessionController {
	GlobalSessions, _ = session.NewManager("memory", "gosessionid", 3600)
	go GlobalSessions.GC()
	return &SessionController{
		Interactor: usecase.SessionInteractor{},
	}
}

func (controller *SessionController) Count(w http.ResponseWriter, r *http.Request) {
	sess := GlobalSessions.SessionStart(w, r)
	countup_sess := controller.Interactor.Count(sess)
	fmt.Fprintln(w, countup_sess.Get("countnum"))
}
