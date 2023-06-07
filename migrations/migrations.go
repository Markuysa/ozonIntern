package migrations

import (
	"database/sql"
	"embed"
	"github.com/pressly/goose/v3"
)

func RunMigrations(connection *sql.DB) error {
	var embedMigrations embed.FS
	goose.SetBaseFS(embedMigrations)
	if err := goose.Up(connection, "."); err != nil {
		return err
	}
	return nil
}
