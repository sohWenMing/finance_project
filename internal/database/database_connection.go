package database_connection

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

var connectionString = "postgres://postgres:postgres@finance_postgres:5432/finance_db?sslmode=disable"

func Connect_db() (db *sql.DB, err error) {
	db, err = sql.Open("postgres", connectionString)

	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	fmt.Println("connected to postgres database.")
	return db, nil
}
