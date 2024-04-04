package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
)

func DataSource() (*Application, error) {
	conn, err := pgx.Connect(context.Background(), "postgresql://test:test@localhost:5432/test")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return &Application{
		Database: conn,
		Memory:   make([]*UserData, 0),
	}, nil
}
