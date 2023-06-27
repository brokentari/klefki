package db

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/jackc/pgx/v5"
)

type Database struct {
	*Queries
}

var (
	db   *Database
	once sync.Once
)

func GetDB() *Database {
	once.Do(func() {
		dbUser := os.Getenv("POSTGRES_USER")
		dbPass := os.Getenv("POSTGRES_PASSWORD")
		dbHost := os.Getenv("POSTGRES_HOST")
		dbPort := os.Getenv("POSTGRES_PORT")
		dbName := os.Getenv("POSTGRES_DB")

		ctx := context.Background()
		connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName)
		conn, err := pgx.Connect(ctx, connString)

		if err != nil {
			panic(err)
		}

		queries := New(conn)

		db = &Database{queries}
	})

	return db

}
