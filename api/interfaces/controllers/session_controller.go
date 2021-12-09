package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/astaxie/session"
	_ "github.com/astaxie/session/providers/memory"
	"github.com/kory-jp/react_golang_api/api/domain"
	"github.com/kory-jp/react_golang_api/api/interfaces/database"
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

func NewManager(sqlHandler database.SqlHandler) *SessionController {
	GlobalSessions, _ = session.NewManager("memory", "gosessionid", 3600)
	go GlobalSessions.GC()
	return &SessionController{
		Interactor: usecase.SessionInteractor{
			SessionRepository: &database.SessionRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *SessionController) Count(w http.ResponseWriter, r *http.Request) {
	sess := GlobalSessions.SessionStart(w, r)
	countup_sess := controller.Interactor.Count(sess)
	fmt.Fprintln(w, countup_sess.Get("countnum"))
}

func (controller *SessionController) Login(w http.ResponseWriter, r *http.Request) {
	sess := GlobalSessions.SessionStart(w, r)
	bytesUser, err := io.ReadAll(r.Body)
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
	}
	userType := new(domain.User)
	if err := json.Unmarshal(bytesUser, userType); err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
		return
	}
	valid, cookie, err := controller.Interactor.Login(*userType, sess)
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
	}
	http.SetCookie(w, &cookie)
	fmt.Println(valid)
	fmt.Fprintln(w, valid)
}

func (controller *SessionController) IsLoggedin(w http.ResponseWriter, r *http.Request) {
	sess := GlobalSessions.SessionStart(w, r)
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
	}
	valid := controller.Interactor.IsLoggedin(sess, cookie)
	fmt.Println(valid)
	fmt.Fprintln(w, valid)
}

func (controller *SessionController) Logout(w http.ResponseWriter, r *http.Request) {
	s := GlobalSessions.SessionStart(w, r)
	c, err := r.Cookie("_cookie")
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
	}
	valid, cookie, err := controller.Interactor.Logout(s, c)
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
	}
	http.SetCookie(w, &cookie)
	fmt.Println(valid)
	fmt.Fprintln(w, valid)
}
