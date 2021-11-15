package database

import (
	"log"

	"github.com/kory-jp/react_golang_api/api/domain"
)

type UserRepository struct {
	SqlHandler
}

func (repo *UserRepository) Store(u domain.User) (id int, err error) {
	result, err := repo.Execute(`
		insert into
			users
			(name, age)
		values
			(?, ?)
	`, u.Name, u.Age)
	if err != nil {
		return
	}
	id64, err := result.LastInsertId()
	if err != nil {
		return
	}
	id = int(id64)
	return
}

func (repo *UserRepository) FindById(identifier int) (user domain.User, err error) {
	// SQL文で改行を含んで記載したい場合は``(バッククォート)
	row, err := repo.Query(`
		select
			id,
			name,
			age
		from
			users
		where
			id = ?
	`, identifier)
	defer row.Close()
	if err != nil {
		log.Fatalln(err)
		return
	}
	var id int
	var name string
	var age int
	row.Next()
	if err = row.Scan(&id, &name, &age); err != nil {
		log.Fatalln(err)
		return
	}
	user.ID = id
	user.Name = name
	user.Age = age
	return
}

func (repo *UserRepository) FindAll() (users domain.Users, err error) {
	rows, err := repo.Query(`
		select
			id, 
			name,
			age
		from
			users
	`)
	defer rows.Close()
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		var id int
		var name string
		var age int
		if err := rows.Scan(&id, &name, &age); err != nil {
			continue
		}
		user := domain.User{
			ID:   id,
			Name: name,
			Age:  age,
		}
		users = append(users, user)
	}
	return
}
