package infrastructure

import (
	"net/http"
	"react_golang_api/api/interfaces/controllers"
)

func Init() {
	http.HandleFunc("/user", controllers.GetUser)
	http.ListenAndServe(":8080", nil)
}
