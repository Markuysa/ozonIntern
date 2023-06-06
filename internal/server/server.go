package server

import (
	"context"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	internalErrors "ozonIntern/internal/errors"
	"ozonIntern/internal/proto/gen"
	"ozonIntern/internal/service"
)

type Server struct {
	linkService service.LinksProcessor
	logger      *zap.Logger
	gen.UnimplementedLinksCreatorServer
}

func NewServer(linkService service.LinksProcessor, logger *zap.Logger) *Server {
	return &Server{linkService: linkService, logger: logger}
}

func (s *Server) GetLink(ctx context.Context, req *gen.GetLinkRequest) (*gen.GetLinkResponse, error) {
	url, err := s.linkService.GetUrlByLink(ctx, req.Short)
	if err != nil {
		if errors.Is(err, internalErrors.ErrAlreadyExists) {
			return &gen.GetLinkResponse{}, status.Error(codes.InvalidArgument, internalErrors.ErrAlreadyExists.Error())
		}
		return &gen.GetLinkResponse{Url: url}, status.Error(codes.Internal, "failed to get")
	}
	return &gen.GetLinkResponse{Url: url}, nil
}
func (s *Server) SaveLink(ctx context.Context, req *gen.SaveLinkRequest) (*gen.SaveLinkResponse, error) {
	link, err := s.linkService.ProcessLink(ctx, req.Url)
	if err != nil {
		return &gen.SaveLinkResponse{}, status.Error(codes.Internal, "failed to save")
	}
	return &gen.SaveLinkResponse{Short: link}, nil
}
