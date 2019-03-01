package driver

import (
	"database/sql"
	"log"
	"os"

	"github.com/lib/pq"
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var db *sql.DB

func Connect() *sql.DB {
	pgURL, err := pq.ParseURL(os.Getenv("DB_URL"))

	logFatal(err)

	db, err := sql.Open("postgres", pgURL)
	err = db.Ping()

	logFatal(err)

	return db
}
