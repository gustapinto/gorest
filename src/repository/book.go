package repository

import (
	"gorest/src/database"
	"log"
)

type Book struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func AddBook(book *Book) {
	db := database.ConnectPostgres()
	defer db.Close()

	query := `insert into "books" ("title", "author") values ($1, $2)`

	if _, err := db.Exec(query, book.Title, book.Author); err != nil {
		log.Fatal(err)
	}
}

func GetAllBooks() []Book {
	var books []Book

	db := database.ConnectPostgres()
	defer db.Close()

	rows, err := db.Query("select * from books")

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var id int
		var title, author string

		if err := rows.Scan(&id, &title, &author); err != nil {
			log.Fatal(err)
		}

		books = append(books, Book{id, title, author})
	}

	return books
}
