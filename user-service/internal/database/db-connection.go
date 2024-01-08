package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type DbModel struct {
	DB                *sql.DB
	InfoLog, ErrorLog *log.Logger
}

func DbConnection() *DbModel {
	var dbModel DbModel

	log.Println("connection is ", os.Getenv("DSN"))

	db, err := sql.Open("pgx", os.Getenv("DSN"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	var greeting string
	err = db.QueryRow("select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	log.Println("sulodo ------- : ", greeting)

	dbModel.DB = db
	return &dbModel
}

func DSNConnection(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil

}
