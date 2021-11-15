package infrastructure

import (
	"net/http"

	"github.com/kory-jp/react_golang_api/api/interfaces/controllers"
)

func Init() {
	// http.HandleFunc("/user", controllers.GetUser)
	userController := controllers.NewUserController(NewSqlHandler())
	http.HandleFunc("/users", userController.Index)
	http.ListenAndServe(":8080", nil)
}
