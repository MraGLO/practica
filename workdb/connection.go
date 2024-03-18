package workdb

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func ConnectDB() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), "host = localhost dbname = NewsDB user = postgres password = 1234 port = 5432")
	if err != nil {
		log.Fatal(err)
	}

	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatal("Connection Error:", err)
	}
	fmt.Println("Connected to the DB Successfully")
	return conn
}
