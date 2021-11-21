package database

import (
	"crypto/sha1"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/kory-jp/react_golang_api/api/domain"
)

type UserRepository struct {
	SqlHandler
}

func (repo *UserRepository) Store(u domain.User) (id int, err error) {
	// uuid作成
	uuidobj, _ := uuid.NewUUID()
	//パスワードを暗号化
	cryptext := fmt.Sprintf("%x", sha1.Sum([]byte(u.Password)))
	result, err := repo.Execute(`
		insert into
			users(
				uuid, 
				name, 
				email,
				password,
				created_at)
		values
			(?, ?, ?, ?, ?)
	`, uuidobj, u.Name, u.Email, cryptext, time.Now())
	if err != nil {
		log.Fatalln(err)
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
			uuid,
			name,
			email,
			password,
			created_at
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
	var uuid string
	var name string
	var email string
	var password string
	var created_at time.Time
	row.Next()
	if err = row.Scan(&id, &uuid, &name, &email, &password, &created_at); err != nil {
		// log.Fatalln(err)
		fmt.Println(err)
		return
	}
	user.ID = id
	user.UUID = uuid
	user.Name = name
	user.Email = email
	user.Password = password
	user.CreatedAt = created_at
	return
}

func (repo *UserRepository) FindAll() (users domain.Users, err error) {
	rows, err := repo.Query(`
		select
			id,
			uuid,
			name,
			email,
			password,
			created_at,
		from
			users
	`)
	defer rows.Close()
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		var id int
		var uuid string
		var name string
		var email string
		var password string
		var created_at time.Time
		if err := rows.Scan(&id, &uuid, &name, &email, &password, &created_at); err != nil {
			continue
		}
		user := domain.User{
			ID:        id,
			UUID:      uuid,
			Name:      name,
			Email:     email,
			Password:  password,
			CreatedAt: created_at,
		}
		users = append(users, user)
	}
	return
}

func (repo *UserRepository) Update(identifier int, u domain.User) (id int, err error) {
	// result, err := repo.Execute(`
	_, error := repo.Execute(`
		update
			users
		set
			name = ?,
			email = ?
		where
			id = ?
`, u.Name, u.Email, identifier)
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(error)
		return
	}
	// ----- IDが取得できない------
	// id64, err := result.LastInsertId()
	// fmt.Println("UR142", id64)
	// if err != nil {
	// 	return
	// }
	id = int(identifier)
	return
}

func (repo *UserRepository) Delete(identifier int) (err error) {
	if _, err := repo.Execute(`delete from users where id = ?`, identifier); err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
	}
	return
}
