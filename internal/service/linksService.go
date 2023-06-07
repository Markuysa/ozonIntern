package service

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/pkg/errors"
	"ozonIntern/internal/database"
	itnernalErrors "ozonIntern/internal/errors"
)

type LinksProcessor interface {
	ProcessLink(ctx context.Context, url string) (string, error)
	CreateLink(ctx context.Context, url string) (string, error)
	GetUrlByLink(ctx context.Context, link string) (string, error)
}

// LinksService is a struct that implements LinksProcessor interface
// and represents the use-case layer
type LinksService struct {
	linksDatabase database.LinksDatabase
}

func NewLinksService(linksDatabase database.LinksDatabase) *LinksService {
	return &LinksService{linksDatabase: linksDatabase}
}
func (l *LinksService) CreateLink(ctx context.Context, url string) (string, error) {
	hash := sha1.Sum([]byte(url))

	// Кодируем хэш в строку
	encoded := base64.URLEncoding.EncodeToString(hash[:])

	// Отбрасываем лишние символы и оставляем только первые 10 символов
	shortLink := encoded[:10]
	checkURL, err := l.linksDatabase.GetURL(ctx, shortLink)
	if err != nil {
		if errors.Is(err, itnernalErrors.ErrUrlNotFound) {
			return shortLink, nil
		}
		return "", errors.Wrap(err, "failed to check uniqueness of the link")
	}
	if checkURL == url {
		return shortLink, nil
	} else {
		return l.CreateLink(ctx, fmt.Sprint(url, "a"))
	}
}
func (l *LinksService) ProcessLink(ctx context.Context, url string) (string, error) {
	link, err := l.CreateLink(ctx, url)
	if err != nil {
		return "", errors.Wrap(err, "failed to create short link")
	}
	err = l.linksDatabase.SaveLink(ctx, url, link)
	if err != nil {
		return "", errors.Wrap(err, "failed to save the link")
	}
	return link, nil
}

func (l *LinksService) GetUrlByLink(ctx context.Context, link string) (string, error) {
	url, err := l.linksDatabase.GetURL(ctx, link)
	if err != nil {
		return "", errors.Wrap(err, "failed to get url by link")
	}
	return url, err
}
