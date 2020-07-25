package grpc

import (
	pb "comet/pkg/api"
	"comet/pkg/handlers"
	"comet/pkg/logger"
	"comet/pkg/protocol/grpc/middleware"
	"context"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
)

// RunServer ...
func RunServer(ctx context.Context, handler handlers.IHandlers, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// gRPC server statup options
	opts := []grpc.ServerOption{}

	// add middleware
	opts = middleware.AddLogging(logger.Log, opts)

	// register handlers
	server := grpc.NewServer(opts...)
	pb.RegisterCosmosServiceServer(server, handler)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			logger.Log.Warn("shutting down gRPC server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	// start gRPC server
	logger.Log.Info("starting gRPC server...")
	return server.Serve(listen)
}
