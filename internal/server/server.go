package server

import (
	"context"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	internalErrors "ozonIntern/internal/errors"
	"ozonIntern/internal/service"
	gen2 "ozonIntern/pkg/proto/gen"
)

type Server struct {
	linkService service.LinksProcessor
	logger      *zap.Logger
	gen2.UnimplementedLinksCreatorServer
}

func NewServer(linkService service.LinksProcessor, logger *zap.Logger) *Server {
	return &Server{linkService: linkService, logger: logger}
}

func (s *Server) GetLink(ctx context.Context, req *gen2.GetLinkRequest) (*gen2.GetLinkResponse, error) {
	url, err := s.linkService.GetUrlByLink(ctx, req.Short)
	if err != nil {
		if errors.Is(err, internalErrors.ErrUrlNotFound) {
			return &gen2.GetLinkResponse{Url: url}, status.Error(codes.NotFound, internalErrors.ErrUrlNotFound.Error())
		}
		return &gen2.GetLinkResponse{Url: url}, status.Error(codes.Internal, "failed to get")
	}
	return &gen2.GetLinkResponse{Url: url}, nil
}
func (s *Server) SaveLink(ctx context.Context, req *gen2.SaveLinkRequest) (*gen2.SaveLinkResponse, error) {
	link, err := s.linkService.ProcessLink(ctx, req.Url)
	if err != nil {
		if errors.Is(err, internalErrors.ErrAlreadyExists) {
			return &gen2.SaveLinkResponse{}, status.Error(codes.AlreadyExists, internalErrors.ErrAlreadyExists.Error())
		}
		return &gen2.SaveLinkResponse{}, status.Error(codes.Internal, "failed to save")
	}
	return &gen2.SaveLinkResponse{Short: link}, nil
}
