package tools

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var Con *sqlx.DB

func DbConnect() *sqlx.DB {
	//Connect to mysql server
	dbPassword := GetEnv("DB_PASS")
	dbName := GetEnv("DB_NAME")
	dbUser := GetEnv("DB_USER")
	dbPort := GetEnv("DB_PORT")

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable", dbUser, dbPassword, dbName, dbPort)
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
