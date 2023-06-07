package database

import (
	"context"
	"fmt"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pressly/goose/v3"
	"log"
	"ozonIntern/internal/config"
	"ozonIntern/migrations"
)

type LinksDatabase interface {
	SaveLink(ctx context.Context, url, link string) error
	GetURL(ctx context.Context, link string) (string, error)
}

func CreateDatabase(ctx context.Context, flag bool) LinksDatabase {
	if !flag {
		return NewInMemDatabase()
	} else {
		connectionPattern := "host=%s port=%s dbname=%s user=%s password=%s sslmode=disable"
		dbConfig, err := config.New()
		if err != nil {
			return nil
		}
		connURL := fmt.Sprintf(connectionPattern,
			dbConfig.Host,
			dbConfig.Port,
			dbConfig.DBName,
			dbConfig.User,
			dbConfig.Password,
		)
		connection, _ := pgxpool.New(ctx, connURL)
		db, err := goose.OpenDBWithDriver("postgres", connURL)
		if err != nil {
			log.Fatal(err)
		}
		err = migrations.RunMigrations(db)
		if err != nil {
			log.Fatal(err)
		}
		return NewPgDatabase(connection)
	}
}
