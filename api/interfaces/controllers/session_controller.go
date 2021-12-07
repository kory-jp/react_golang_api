package controllers

import (
	"fmt"
	"net/http"

	"github.com/astaxie/session"
	_ "github.com/astaxie/session/providers/memory"
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

func NewManager() {
	GlobalSessions, _ = session.NewManager("memory", "gosessionid", 3600)
	go GlobalSessions.GC()
}

func Count(w http.ResponseWriter, r *http.Request) {
	sess := GlobalSessions.SessionStart(w, r)
	ct := sess.Get("countnum")
	fmt.Println(sess)
	if ct == nil {
		sess.Set("countnum", 1)
	} else {
		sess.Set("countnum", (ct.(int) + 1))
	}
	w.Header().Set("Content-type", "application/json")
	fmt.Fprintln(w, sess.Get("countnum"))
}
