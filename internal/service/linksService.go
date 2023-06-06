package service

import (
	"context"
	"github.com/pkg/errors"
	"ozonIntern/internal/database"
)

type LinksProcessor interface {
	ProcessLink(ctx context.Context, url string) (string, error)
	IsAlreadyExist(ctx context.Context, url string) (bool, error)
	GetUrlByLink(ctx context.Context, link string) (string, error)
}

// LinksService is a struct that implements LinksProcessor interface
// and represents the use-case layer
type LinksService struct {
	linksDatabase database.LinksDatabase
}

func IsAlreadyExist(ctx context.Context, url string) (bool, error) {

	return true, nil
}
func NewLinksService(linksDatabase database.LinksDatabase) *LinksService {
	return &LinksService{linksDatabase: linksDatabase}
}

func (l *LinksService) ProcessLink(ctx context.Context, url string) (string, error) {
	link := url + "salt"
	// create short link here
	err := l.linksDatabase.SaveLink(ctx, url, link)
	if err != nil {
		return "", errors.Wrap(err, "failed to save the link")
	}
	return link, nil
}

func (l *LinksService) GetUrlByLink(ctx context.Context, link string) (string, error) {
	return l.linksDatabase.GetURL(ctx, link)
}
