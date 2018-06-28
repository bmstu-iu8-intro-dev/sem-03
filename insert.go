package main

import (
	_ "github.com/lib/pq"
	"database/sql"
	"log"
)

func insertToDB(db * sql.DB) {
	db.Query("INSERT INTO books (author, title) VALUES ($1, $2)", "Bjarne Stroustrup", "A Tour of C++")
	db.Query("INSERT INTO books (author, title) VALUES ($1, $2)", "Scott Meyers", "Effective C++")
	db.Query("INSERT INTO books (author, title) VALUES ($1, $2)", "Scott Meyers", "More Effective C++")
	db.Query("INSERT INTO books (author, title) VALUES ($1, $2)", "Exceptional C++", "Herb Sutter")
	db.Query("INSERT INTO books (author, title) VALUES ($1, $2)", "Pushkin", "Eugene Onegin")
	db.Query("INSERT INTO books (author, title) VALUES ($1, $2)", "Pushkin", "The Captain's Daughter")
}

func updateDB(db * sql.DB) {
	db.Query("UPDATE books SET author = 'A.S.Pushkin' WHERE author = 'Pushkin'")
}

func removeValues(db * sql.DB) {
	db.Query("DELETE FROM books WHERE author <> 'A.S.Pushkin';")
}

func main() {
	db, err := sql.Open("postgres", "postgres://test_user:0@localhost/first_db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	updateDB(db)
}
