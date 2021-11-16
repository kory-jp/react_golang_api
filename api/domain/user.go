package domain

type User struct {
	ID   int    `josn:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Users []User
