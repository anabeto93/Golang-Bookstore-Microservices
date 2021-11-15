package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

const (
	DB_USERNAME = "MYSQL_USERS_DB_USERNAME"
	DB_PASSWORD = "MYSQL_USERS_DB_PASSWORD"
	DB_HOST     = "MYSQL_USERS_DB_HOST"
	DB_SCHEMA   = "MYSQL_USERS_DB_SCHEMA"
)

var (
	Client *sql.DB
)

func init() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading Env")
	}

	username := os.Getenv(DB_USERNAME)
	password := os.Getenv(DB_PASSWORD)
	host := os.Getenv(DB_HOST)
	schema := os.Getenv(DB_SCHEMA)

	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, schema)
	Client, err = sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("users database successfully initialized.")
}
