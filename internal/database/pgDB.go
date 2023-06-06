package database

import (
	"context"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/pkg/errors"
)

type PgDatabase struct {
	pool *pgxpool.Pool
}

func NewPgDatabase(pool *pgxpool.Pool) *PgDatabase {
	return &PgDatabase{pool: pool}
}

func (db *PgDatabase) SaveLink(ctx context.Context, url, link string) error {
	query := `
	insert into link(
	                 original_url,
	                 short_link
	)values (
	         $1,$2
	)
`
	_, err := db.pool.Exec(ctx, query, url, link)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == "23505" {
				return ErrAlreadyExists
			} else {
				return errors.Wrap(err, "failed to save the link")
			}
		}
	}
	return nil
}
func (db *PgDatabase) GetURL(ctx context.Context, link string) (string, error) {
	query := `
	select original_url
	from link
	where short_link=$1
`
	var url string
	err := db.pool.QueryRow(ctx, query, link).Scan(&url)
	if err != nil {
		return "", errors.Wrap(err, "failed to get the url")
	}

	return url, nil
}
