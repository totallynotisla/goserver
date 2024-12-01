package db

import (
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var Con *sqlx.DB

func DbConnect() *sqlx.DB {
	//Connect to mysql server
	connStr := "user=postgres password=kmzwa8awaa dbname=gerawana sslmode=disable"
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		panic(err.Error())
	}

	return db;
}

func InitDB(con *sqlx.DB) (error) {
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