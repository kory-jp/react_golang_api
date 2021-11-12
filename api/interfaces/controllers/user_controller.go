package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// type Test struct {
// 	Title  string `json:"title"`
// 	Number int    `json:"number"`
// }

func GetUser(w http.ResponseWriter, r *http.Request) {
	// t := Test{
	// 	Title:  "test",
	// 	Number: 2,
	// }

	t := struct {
		Title  string `json:"title"`
		Number int    `json:"number"`
	}{
		Title:  "test",
		Number: 2,
	}
	s, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(s))
}
