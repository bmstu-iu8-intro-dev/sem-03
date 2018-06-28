package main

import (
	_ "github.com/lib/pq"
	"database/sql"
	"log"
)

func createTable(db * sql.DB) {
	db.Query("CREATE TABLE books (" +
		"id SERIAL PRIMARY KEY" +
		"title VARCHAR(64)" +
		"author VARCHAR(64)" +
		")")
}

func main() {
	db, err := sql.Open("postgres", "postgres://test_user:0@localhost/first_db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	createTable(db)
}