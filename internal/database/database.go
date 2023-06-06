package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"ozonIntern/internal/config"
)

type LinksDatabase interface {
	SaveLink(ctx context.Context, url, link string) error
	GetURL(ctx context.Context, link string) (string, error)
}

func CreateDatabase(ctx context.Context, flag bool) LinksDatabase {
	if !flag {
		return NewInMemDatabase()
	} else {
		connectionPattern := "postgresql://%s:%s@%s:%s/%s"
		dbConfig, err := config.New()
		if err != nil {
			return nil
		}
		connURL := fmt.Sprintf(connectionPattern,
			dbConfig.User,
			dbConfig.Password,
			dbConfig.Host,
			dbConfig.Port,
			dbConfig.DBName,
		)
		connection, _ := pgxpool.New(ctx, connURL)
		return NewPgDatabase(connection)
	}
}
