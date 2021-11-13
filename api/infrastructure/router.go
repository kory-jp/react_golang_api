package infrastructure

import (
	"net/http"

	"github.com/kory-jp/react_golang_api/api/interfaces/controllers"
)

func Init() {
	http.HandleFunc("/user", controllers.GetUser)
	http.ListenAndServe(":8080", nil)
}
