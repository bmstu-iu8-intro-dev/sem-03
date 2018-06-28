package main

import (
	_ "github.com/lib/pq"
	"database/sql"
	"log"
)

type Author struct {
	firstName string
	lastName string
}

type Book struct {
	title string
	author Author
	pages int
}

func selectAll(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int8
		var title string
		var author string
		err := rows.Scan(&id, &title, &author)
		if err != nil {
			log.Fatal(err)
		} else {
			log.Printf("%d: %v - %v", id, title, author)
		}
	}
}

func selectColumns(db *sql.DB) {
	rows, err := db.Query("SELECT author, title FROM books")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var title string
		var author string
		err := rows.Scan(&author, &title)
		if err != nil {
			log.Fatal(err)
		} else {
			log.Printf("%v - %v", title, author)
		}
	}
}

func selectBooksByAuthor(db* sql.DB, author string) {
	rows, err := db.Query("SELECT id, title FROM books WHERE author = $1", author)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int8
		var title string
		err := rows.Scan(&id, &title)
		if err != nil {
			log.Fatal(err)
		} else {
			log.Printf("%v - %v", id, title)
		}
	}
}

func selectBooks(db* sql.DB) {
	rows, err := db.Query("SELECT title FROM books WHERE LOWER(title) LIKE '%effective%'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var title string
		err := rows.Scan(&title)
		if err != nil {
			log.Fatal(err)
		} else {
			log.Printf(title)
		}
	}
}

func selectById(db* sql.DB) {
	rows, err := db.Query("SELECT id, title FROM books WHERE id > 3")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int8
		var title string
		err := rows.Scan(&id, &title)
		if err != nil {
			log.Fatal(err)
		} else {
			log.Printf("%v - %v", id, title)
		}
	}
}

func main() {
	db, err := sql.Open("postgres", "postgres://test_user:0@localhost/first_db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	selectAll(db)
	selectColumns(db)
	selectBooksByAuthor(db, "Scott Meyers")
	selectBooks(db)
	selectById(db)
}
