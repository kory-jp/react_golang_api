package infrastructure

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	_ "github.com/astaxie/session/providers/memory"
	"github.com/kory-jp/react_golang_api/api/interfaces/controllers"
)

//正規表現を利用してURLを解析
var validPath = regexp.MustCompile("^/(users|todos)/(show|update|delete)/([0-9]+)$")

//URLからIDを解析して返却
func parseURL(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		fmt.Println(validPath)
		q := validPath.FindStringSubmatch(r.URL.Path)
		fmt.Println(q)
		if q == nil {
			http.NotFound(w, r)
			return
		}
		// 	strconv.Atoi = 文字列 → 数値変換（パース）
		qi, err := strconv.Atoi(q[3])
		if err != nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, qi)
	}
}

func Init() {
	userController := controllers.NewUserController(NewSqlHandler())
	todoController := controllers.NewTodoController(NewSqlHandler())
	sessionController := controllers.NewManager(NewSqlHandler())
	// http.HandleFunc("/user", controllers.GetUser)
	http.HandleFunc("/users/new", userController.Create)
	// http.HandleFunc("/users/index", userController.Index)
	http.HandleFunc("/users/show/", parseURL(userController.Show))
	http.HandleFunc("/users/update/", parseURL(userController.Update))
	http.HandleFunc("/users/delete/", parseURL(userController.Delete))
	http.HandleFunc("/todos/new", todoController.Create)
	http.HandleFunc("/todos/index", todoController.Index)
	http.HandleFunc("/todos/show/", parseURL(todoController.Show))
	http.HandleFunc("/todos/update/", parseURL(todoController.Update))
	http.HandleFunc("/todos/delete/", parseURL(todoController.Delete))
	// http.HandleFunc("/cookie", controllers.Cookie)
	// http.HandleFunc("/session", sessionController.Count)
	http.HandleFunc("/login", sessionController.Login)
	http.HandleFunc("/isloggedin", sessionController.IsLoggedin)
	http.HandleFunc("/logout", sessionController.Logout)
	http.ListenAndServe(":8080", nil)
}
