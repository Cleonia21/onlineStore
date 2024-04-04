package main

import (
	"database/sql"
	"fmt"
	"os"
)


type dataBase struct {
	db *sql.DB
}

func (db *dataBase) ConnectDB() {
	var err error
	db.db, err = sql.Open("postgres", connectString())

	if err != nil {
		fmt.Errorf("Failed to connect to database. \n", err)
		os.Exit(2)
	}
	fmt.Println("data base connected")
}

func connectString() string {
	connStr := fmt.Sprintf("postgres://%v:%v@%v:5432/%v?sslmode=disable",
		"cleonia",
		"2222",
		"localhost",
		"onlineStore",
	)
	return connStr
}
