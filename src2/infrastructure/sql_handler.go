package infrastructure

import (
	"database/sql"
	"fmt"
)

type SqlHandler struct {
	db *sql.DB
}

func NewSqlHandler() *SqlHandler {
	connStr := fmt.Sprintf("postgres://%v:%v@%v:5432/%v?sslmode=disable",
		"postgres",
		"2222",
		"localhost",
		"onlineStore",
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err.Error())
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.db = db
	return sqlHandler
}

func (handler *SqlHandler) FindSome(id int) string {
	return ""
}
