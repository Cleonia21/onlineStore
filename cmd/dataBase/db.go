package dataBase

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type DataBase struct {
	db *sql.DB
}

// ConnectDB осуществляет подключение к базе данных
func (db *DataBase) ConnectDB() error {
	var err error
	db.db, err = sql.Open("postgres", connectString())

	if err != nil {
		return fmt.Errorf("failed to connect to database: %v\n", err.Error())
	}
	//fmt.Println("data base connected")
	return nil
}

// connectString возвращает строку подключения к базе данных
func connectString() (connStr string) {
	connStr = fmt.Sprintf("postgres://%v:%v@%v:5432/%v?sslmode=disable",
		"postgres",
		"2222",
		"localhost",
		"onlineStore",
	)
	return connStr
}

// Close закрывает подключение к базе данных
func (db *DataBase) Close() error {
	err := db.db.Close()
	return err
}
