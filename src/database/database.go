package database

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/jackc/pgx/v5"
)

type DB struct {
	*pgx.Conn
}

func NewDB() (*DB, error) {
	var db DB

	// Reading database pass from compose secret file. Check compose file for more info
	pass, err := os.ReadFile("/run/secrets/db_password")
	if err != nil {
		return &db, err
	}

	// need to reference db here since this is the database service name and this is how we can reach it in the docker network
	connection_str := fmt.Sprintf("postgres://%s:%s@db:5432/%s",
		os.Getenv("DB_USER"),
		pass,
		os.Getenv("DB_DB"),
	)

	conn, err := pgx.Connect(context.Background(), connection_str)
	if err != nil {
		return &db, err
	}

	db.Conn = conn

	return &db, nil
}

func (db DB) Migrate() error {
	dir, err := os.ReadDir("./migrations")
	if err != nil {
		return err
	}

	tx, err := db.Conn.Begin(context.Background())
	if err != nil {
		return err
	}

	for _, d := range dir {
		if d.IsDir() {
			fmt.Printf("%: is a dir, skipping...\n", d.Name())
			continue
		}

		b, err := os.ReadFile(path.Join("migrations", d.Name()))
		if err != nil {
			tx.Rollback(context.Background())
			return err
		}

		_, err = tx.Exec(context.Background(), string(b))
		if err != nil {
			tx.Rollback(context.Background())
			return err
		}
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}

	return nil
}
