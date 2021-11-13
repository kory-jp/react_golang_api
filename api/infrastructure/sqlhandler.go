package infrastructure

import (
	"database/sql"
	"fmt"

	"github.com/kory-jp/react_golang_api/api/config"
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
