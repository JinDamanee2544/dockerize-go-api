package post

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type DB interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Exec(query string, args ...interface{}) (sql.Result, error)
}

var db DB

type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func ConnectDB() {
	url := os.Getenv("DATABASE_URL")
	if url == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	var err error
	db, err = sql.Open("postgres", url)

	// log.Print(db)

	if err != nil {
		log.Fatal(err.Error())
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS posts (
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL
	);
	`

	_, err = db.Exec(createTable)

	if err != nil {
		log.Fatal(err.Error())
	}
}
