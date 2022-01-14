package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
	"os"
)

type Book struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func AddBook(book *Book) {
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"postgres", 5432, os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"))

	log.Println(connString)

	db, err := sql.Open("pgx", connString)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	query := `insert into "books" ("title", "author") values ($1, $2)`

	if _, err := db.Exec(query, book.Title, book.Author); err != nil {
		log.Fatal(err)
	}
}
