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

func (repo *TodoRepository) FindAll() (todos domain.Todos, err error) {
	rows, err := repo.Query(`
		select
			id,
			user_id,
			content,
			created_at
		from
			todos
	`)
	defer rows.Close()
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
	}
	for rows.Next() {
		var id int
		var user_id int
		var content string
		var created_at time.Time
		if err := rows.Scan(&id, &user_id, &content, &created_at); err != nil {
			continue
		}
		todo := domain.Todo{
			ID:        id,
			UserID:    user_id,
			Content:   content,
			CreatedAt: created_at,
		}
		todos = append(todos, todo)
	}
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
