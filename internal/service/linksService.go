package service

import "context"

type LinksProcessor interface {
	SaveLink(ctx context.Context, url string) (error, string)
	GetUrlByLink(ctx context.Context, link string) (error, string)
}

// LinksService is a struct that implements LinksProcessor interface
// and represents the use-case layer
type LinksService struct {
}

func (l *LinksService) SaveLink(ctx context.Context, url string) (error, string) {

	return nil, ""
}

func (l *LinksService) GetUrlByLink(ctx context.Context, link string) (error, string) {
	return nil, ""
}
