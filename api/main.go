package main

import (
	"fmt"
	"react_golang_api/api/infrastructure"
)

func main() {
	fmt.Println("hello")
	infrastructure.Init()

	// infrastructure.Init() (no value) used as value = Initメソッドで戻り値を設定していないのにresに戻り値を代入しようとしていた
	// res, _ := infrastructure.Init()
	// fmt.Println(res)
}
