package workdb

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDB() *pgxpool.Pool {
	dbpool, err := pgxpool.New(context.Background(), "host = localhost dbname = NewsDB user = postgres password = 1234 port = 5432")
	if err != nil {
		log.Println(err)
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	return dbpool
}
