package persistence

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type SQLConnection struct {
	conn *pgxpool.Pool
}

func NewSQLConnection() *SQLConnection {
	connString := "postgres://postgres:yourpassword@db:5432/yourdbname?sslmode=disable"
	pool, err := pgxpool.New(context.Background(), connString)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return &SQLConnection{
		conn: pool,
	}
}

func (s *SQLConnection) Close() {
	if s.conn != nil {
		s.Close()
	}
}
