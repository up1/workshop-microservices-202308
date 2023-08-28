package demo

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func CreateDbConnection() *pgxpool.Pool {
	databaseUrl := "postgres://demo:mypassword@db:5432/sample_db?sslmode=disable"
	dbPool, err := pgxpool.Connect(context.Background(), databaseUrl)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return dbPool
}
