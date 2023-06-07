package service

import (
	"context"
	"database/sql"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	internalErrors "ozonIntern/internal/errors"
	mockdatabase "ozonIntern/internal/mock"
	"testing"
)

// ---------------------------------------------------------------------------------------------------------------------------------------------------------------------//
func TestCreateLinkReturnsLink(t *testing.T) {
	ctx := context.Background()
	expect := "dA5zl5B8Cw"
	testURL := "https://www.example.com"
	ctrl := gomock.NewController(t)
	repos := mockdatabase.NewMockLinksDatabase(ctrl)
	linkService := LinksService{linksDatabase: repos}

	repos.EXPECT().GetURL(ctx, expect).Return("", internalErrors.ErrUrlNotFound)

	resultURL, err := linkService.CreateLink(ctx, testURL)

	assert.Equal(t, expect, resultURL)
	assert.Equal(t, err, nil)
}
func TestCreateLinkReturnsLinkAlreadyExists(t *testing.T) {
	ctx := context.Background()
	expect := "dA5zl5B8Cw"
	testURL := "https://www.example.com"
	ctrl := gomock.NewController(t)
	repos := mockdatabase.NewMockLinksDatabase(ctrl)
	linkService := LinksService{linksDatabase: repos}

	repos.EXPECT().GetURL(ctx, expect).Return(testURL, nil)

	resultURL, err := linkService.CreateLink(ctx, testURL)

	assert.Equal(t, expect, resultURL)
	assert.Equal(t, err, nil)
}
func TestCreateLinkReturnsLinkCollisionError(t *testing.T) {
	ctx := context.Background()
	expect := "dA5zl5B8Cw"
	expect_collision := "zxZmBeSpxo"
	testURL_first := "https://www.example.com"
	testURL_second := "https://www.yandex.ru"
	ctrl := gomock.NewController(t)
	repos := mockdatabase.NewMockLinksDatabase(ctrl)
	linkService := LinksService{linksDatabase: repos}

	repos.EXPECT().GetURL(ctx, expect).Return(testURL_second, nil)

	repos.EXPECT().GetURL(ctx, expect_collision).Return("", internalErrors.ErrUrlNotFound)

	resultURL, err := linkService.CreateLink(ctx, testURL_first)

	assert.Equal(t, expect_collision, resultURL)
	assert.Equal(t, err, nil)
}
func TestCreateLinkReturnsDbError(t *testing.T) {
	ctx := context.Background()
	expect := "dA5zl5B8Cw"
	testURL := "https://www.example.com"
	ctrl := gomock.NewController(t)
	repos := mockdatabase.NewMockLinksDatabase(ctrl)
	linkService := LinksService{linksDatabase: repos}

	repos.EXPECT().GetURL(ctx, expect).Return("", sql.ErrConnDone)

	_, err := linkService.CreateLink(ctx, testURL)

	assert.Equal(t, err.Error(), "failed to check uniqueness of the link: sql: connection is already closed")
}

//---------------------------------------------------------------------------------------------------------------------------------------------------------------------//

func TestProcessLinkReturnsSavedLink(t *testing.T) {
	ctx := context.Background()
	expect := "dA5zl5B8Cw"
	testURL := "https://www.example.com"
	ctrl := gomock.NewController(t)
	repos := mockdatabase.NewMockLinksDatabase(ctrl)
	linkService := LinksService{linksDatabase: repos}
	repos.EXPECT().GetURL(ctx, expect).Return("", internalErrors.ErrUrlNotFound)
	repos.EXPECT().SaveLink(ctx, testURL, expect).Return(nil)
	resultLink, err := linkService.ProcessLink(ctx, testURL)

	assert.Equal(t, err, nil)
	assert.Equal(t, resultLink, expect)
}

func TestProcessLinkReturnsShortLinkCreationError(t *testing.T) {
	ctx := context.Background()
	expect := "dA5zl5B8Cw"
	testURL := "https://www.example.com"
	ctrl := gomock.NewController(t)
	repos := mockdatabase.NewMockLinksDatabase(ctrl)
	linkService := LinksService{linksDatabase: repos}

	repos.EXPECT().GetURL(ctx, expect).Return("", sql.ErrConnDone)
	_, err := linkService.ProcessLink(ctx, testURL)

	assert.Equal(t, err.Error(), "failed to create short link: failed to check uniqueness of the link: sql: connection is already closed")
}

func TestProcessLinkReturnsSaveLinkError(t *testing.T) {
	ctx := context.Background()
	expect := "dA5zl5B8Cw"
	testURL := "https://www.example.com"
	ctrl := gomock.NewController(t)
	repos := mockdatabase.NewMockLinksDatabase(ctrl)
	linkService := LinksService{linksDatabase: repos}

	repos.EXPECT().GetURL(ctx, expect).Return(testURL, nil)
	repos.EXPECT().SaveLink(ctx, testURL, expect).Return(sql.ErrConnDone)
	_, err := linkService.ProcessLink(ctx, testURL)

	assert.Equal(t, err.Error(), "failed to save the link: sql: connection is already closed")
}

//---------------------------------------------------------------------------------------------------------------------------------------------------------------------//

func TestGetUrlByLinkReturnsUrlNoErr(t *testing.T) {
	ctx := context.Background()
	expect := "dA5zl5B8Cw"
	testURL := "https://www.example.com"
	ctrl := gomock.NewController(t)
	repos := mockdatabase.NewMockLinksDatabase(ctrl)
	linkService := LinksService{linksDatabase: repos}

	repos.EXPECT().GetURL(ctx, expect).Return(testURL, nil)

	resultURL, err := linkService.GetUrlByLink(ctx, expect)

	assert.Equal(t, err, nil)
	assert.Equal(t, resultURL, testURL)
}
func TestGetUrlByLinkReturnsDbErr(t *testing.T) {
	ctx := context.Background()
	expect := "dA5zl5B8Cw"
	ctrl := gomock.NewController(t)
	repos := mockdatabase.NewMockLinksDatabase(ctrl)
	linkService := LinksService{linksDatabase: repos}

	repos.EXPECT().GetURL(ctx, expect).Return("", sql.ErrConnDone)

	_, err := linkService.GetUrlByLink(ctx, expect)

	assert.Equal(t, err.Error(), "failed to get url by link: sql: connection is already closed")
}
