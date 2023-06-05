package database

import "context"

type LinksDatabase interface {
	SaveLink(ctx context.Context, url string) error
	GetURL(ctx context.Context, link string) (string, error)
}
