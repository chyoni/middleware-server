package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"

	"github.com/cenkalti/backoff/v4"
)

func InitStore() (*sql.DB, error) {

	pgConString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		os.Getenv("PGHOST"), os.Getenv("PGPORT"), os.Getenv("PGDATABASE"), os.Getenv("PGUSER"), os.Getenv("PGPASSWORD"))

	var (
		db  *sql.DB
		err error
	)

	openDB := func() error {
		db, err = sql.Open("postgres", pgConString)
		return err
	}

	err = backoff.Retry(openDB, backoff.NewExponentialBackOff())
	if err != nil {
		return nil, err
	}
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS message (value STRING PRIMARY KEY)"); err != nil {
		return nil, err
	}
	return db, nil
}
