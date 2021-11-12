package main

import (
	"react_golang_api/api/infrastructure"
)

func main() {
	infrastructure.Init()
	// curl localhost:8080 でアクセスして404ページが返却されたら接続成功
}
