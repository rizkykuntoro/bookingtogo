package models

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB
var _ = godotenv.Load(".env")

// const (
// 	username = "root"
// 	password = "root"
// 	host     = "localhost"
// 	database = "bookingtogo"
// )

func Init() *sql.DB {
	var err error

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	database := os.Getenv("DB_DATABASE")

	connStr := "postgres://" + username + ":" + password + "@" + host + "/" + database + "?sslmode=disable"
	db, err = sql.Open("postgres", connStr)

	if err != nil {
		panic(err.Error())
	}

	return db
}
