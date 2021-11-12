package infrastructure

import (
	"fmt"
	"net/http"
)

func Init() {
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("接続失敗", err)
	}
}
