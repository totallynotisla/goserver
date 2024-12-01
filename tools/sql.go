package db

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var Con *sqlx.DB
var db_password string = "root"

func DbConnect() *sqlx.DB {
	//Connect to mysql server
	connStr := fmt.Sprintf("user=postgres password=%s dbname=gerawana sslmode=disable", db_password)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		panic(err.Error())
	}

	return db
}

func InitDB(con *sqlx.DB) error {
	schema, err := os.ReadFile("database/schema.sql")
	if err != nil {
		return err
	}

	_, err = con.Exec(string(schema))
	if err != nil {
		return err
	}

	return nil
}
