package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/kory-jp/react_golang_api/api/config"
	"github.com/kory-jp/react_golang_api/api/interfaces/database"
)

type SqlHandler struct {
	Conn *sql.DB
}

const (
	tableNameUser = "users"
	tableNameTodo = "todos"
	// tableNameSession = "sessions"
)

func NewSqlHandler() *SqlHandler {
	DSN := fmt.Sprintf("%s:%s@%s/%s?parseTime=true", config.Config.UserName, config.Config.Password, config.Config.DBPort, config.Config.DBname)
	conn, err := sql.Open(config.Config.SQLDriver, DSN)
	if err != nil {
		panic(err.Error())
	}

	cmdU := fmt.Sprintf(`
	create table if not exists %s(
		id integer primary key auto_increment,
		uuid varchar(50) NOT NULL UNIQUE,
		name varchar(50),
		email varchar(50),
		password varchar(50),
		created_at datetime default current_timestamp
	)`, tableNameUser)

	_, errU := conn.Exec(cmdU)
	if errU != nil {
		log.Fatalln(errU)
	}

	cmdT := fmt.Sprintf(`
	create table if not exists %s (
		id integer primary key auto_increment,
			content text,
			user_id integer,
			created_at datetime default current_timestamp
	)`, tableNameTodo)

	_, errT := conn.Exec(cmdT)
	if errT != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
	}

	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}

func (handler *SqlHandler) Execute(statemant string, args ...interface{}) (database.Result, error) {
	res := SqlResult{}
	result, err := handler.Conn.Exec(statemant, args...)
	if err != nil {
		log.SetFlags(log.Llongfile)
		log.Println(err)
		return res, err
	}
	res.Result = result
	return res, nil
}

func (handler *SqlHandler) Query(statement string, args ...interface{}) (database.Row, error) {
	rows, err := handler.Conn.Query(statement, args...)
	if err != nil {
		return new(SqlRow), err
	}
	row := new(SqlRow)
	row.Rows = rows
	return row, nil
}

type SqlResult struct {
	Result sql.Result
}

func (r SqlResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r SqlResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

type SqlRow struct {
	Rows *sql.Rows
}

func (r SqlRow) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest...)
}

func (r SqlRow) Next() bool {
	return r.Rows.Next()
}

func (r SqlRow) Close() error {
	return r.Rows.Close()
}
