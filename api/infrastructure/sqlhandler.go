package infrastructure

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/kory-jp/react_golang_api/api/config"
	"github.com/kory-jp/react_golang_api/api/interfaces/database"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler() *SqlHandler {
	DSN := fmt.Sprintf("%s:%s@%s/%s?parseTime=true", config.Config.UserName, config.Config.Password, config.Config.DBPort, config.Config.DBname)
	conn, err := sql.Open(config.Config.SQLDriver, DSN)
	if err != nil {
		panic(err.Error())
	}

	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}

func (handler *SqlHandler) Execute(statemant string, args ...interface{}) (database.Result, error) {
	res := SqlResult{}
	result, err := handler.Conn.Exec(statemant, args...)
	if err != nil {
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
