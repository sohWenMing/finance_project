package database_connection

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "finance_postgres"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "finance_db"
)

func Connect_db() (db *sql.DB, err error) {
	dbInfo := fmt.Sprintf(
		"host=%s "+
			"port=%d "+
			"user=%s "+
			"password=%s "+
			"dbname=%s "+
			"sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open("postgres", dbInfo)

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
