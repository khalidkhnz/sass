package services

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/khalidkhnz/sass/go-ecom/config"
)

var DATABASE *pgx.Conn

func ConnectToDb() (*pgx.Conn, error) {
	dbURL := config.DbUri()

	conn, err := pgx.Connect(context.TODO(), dbURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	DATABASE = conn
	return conn, nil
}