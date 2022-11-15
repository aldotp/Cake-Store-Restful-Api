package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func ConnectToDatabase() *sql.DB {
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASS")
	HOST := os.Getenv("DB_HOST")
	PORT := os.Getenv("DB_PORT")
	DBNAME := os.Getenv("DB_NAME")

	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	// test := "root:@tcp(localhost:3306)/cakestore?charset=utf8&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", URL)
	if err != nil {
		panic(err)
	}

	return db
}

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
		os.Exit(1)
	}
}

// "root:root@tcp(localhost:3306)/cakestore?charset=utf8&parseTime=True&loc=Local"
