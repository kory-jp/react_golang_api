package database

import (
	"log"
	"time"

	"github.com/kory-jp/react_golang_api/api/domain"
)

type TodoRepository struct {
	SqlHandler
}

func (repo *TodoRepository) Store(t domain.Todo) (id int, err error) {
	result, err := repo.Execute(`
		insert into
			todos(
				content,
				user_id,
				created_at
			)
		value
			(?, ?, ?)
	`, t.Content, t.UserID, time.Now())
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
	}
	id64, err := result.LastInsertId()
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
		return
	}
	id = int(id64)
	return
}

func (repo *TodoRepository) FindById(identifier int) (todo domain.Todo, err error) {
	row, err := repo.Query(`
		select
			id,
			content,
			user_id,
			created_at
		from
			todos
		where
			id = ?
	`, identifier)
	defer row.Close()
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
	}
	var id int
	var content string
	var user_id int
	var created_at time.Time
	row.Next()
	if err = row.Scan(&id, &content, &user_id, &created_at); err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
	}
	todo.ID = id
	todo.Content = content
	todo.UserID = user_id
	todo.CreatedAt = created_at
	return
}

func (repo *TodoRepository) Update(identifier int, t domain.Todo) (id int, err error) {
	_, error := repo.Execute(`
			update
				todos
			set
				content = ?
			where
				id = ?
	`, t.Content, identifier)
	if error != nil {
		log.SetFlags(log.Llongfile)
		log.Println(error)
		return
	}
	id = int(identifier)
	return
}

func (repo *TodoRepository) Delete(identifier int) (err error) {
	if _, err := repo.Execute(`delete from todos where id = ?`, identifier); err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
	}
	return
}
