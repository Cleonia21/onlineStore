package dataBase

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

type DataBase struct {
	db *sql.DB
}

func (db *DataBase) ConnectDB() {
	var err error
	db.db, err = sql.Open("postgres", connectString())

	if err != nil {
		fmt.Errorf("failed to connect to database err=%v\n", err.Error())
		os.Exit(2)
	}
	//fmt.Println("data base connected")
}

func connectString() string {
	connStr := fmt.Sprintf("postgres://%v:%v@%v:5432/%v?sslmode=disable",
		"postgres",
		"2222",
		"localhost",
		"onlineStore",
	)
	return connStr
}

func (db *DataBase) Close() error {
	err := db.db.Close()
	return err
}
