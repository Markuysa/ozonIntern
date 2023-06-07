package main

import (
	"context"
	"flag"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"ozonIntern/internal/database"
	logger2 "ozonIntern/internal/logger"
	"ozonIntern/internal/server"
	"ozonIntern/internal/service"
	gen2 "ozonIntern/pkg/proto/gen"
)

func main() {

	repositoryMode := flag.Bool("db", true, "repository mode: if -db is specified => the program will use PostgreSQL")
	flag.Parse()
	logger, err := logger2.InitLogger()
	if err != nil {
		log.Fatal(errors.Wrap(err, "can't init logger"))
	}
	ctx := context.Background()
	// DB
	linkDB := database.CreateDatabase(ctx, *repositoryMode)
	// SERVICE
	linkService := service.NewLinksService(linkDB)
	// SERVER
	go serveGRPC(linkService, logger)
	serveHTTP(logger)
}
func serveHTTP(logger *zap.Logger) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	grpcLinksConn, err := grpc.Dial(
		"localhost:9000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	err = gen2.RegisterLinksCreatorHandler(
		ctx,
		mux,
		grpcLinksConn,
	)
	if err != nil {
		log.Fatal(err)
	}
	logger.Info("Starting HTTP server at:", zap.String("port", "8080"))
	log.Fatal(http.ListenAndServe(":8080", mux))
}
func serveGRPC(linkService service.LinksProcessor, logger *zap.Logger) {
	grpcServer := grpc.NewServer()
	linkServer := server.NewServer(linkService, logger)
	gen2.RegisterLinksCreatorServer(grpcServer, linkServer)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}
	logger.Info("Starting grpc server at:", zap.String("port", "9000"))

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
