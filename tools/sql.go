package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DbConnect() *sql.DB {
	//Connect to mysql server
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}

	return db;
}


