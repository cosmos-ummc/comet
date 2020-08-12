package cmd

import (
	"comet/pkg/handlers"
	"comet/pkg/logger"
	model2 "comet/pkg/model"
	"comet/pkg/protocol/grpc"
	"comet/pkg/protocol/rest"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

// Config is configuration for Server
type Config struct {
	GRPCPort      string
	HTTPPort      string
	LogLevel      int
	LogTimeFormat string
}

// RunServer runs gRPC server and HTTP gateway
func RunServer() error {
	ctx := context.Background()

	// get configuration
	cfg := &Config{GRPCPort: "10001", HTTPPort: "10002", LogLevel: -1, LogTimeFormat: "02 Jan 2006 15:04:05 MST"}

	// initialize logger
	if err := logger.Init(cfg.LogLevel, cfg.LogTimeFormat); err != nil {
		return fmt.Errorf("failed to initialize logger: %v", err)
	}

	// Load .env configuration
	err := godotenv.Load()
	if err != nil {
		logger.Log.Warn(".env file not found, using environment variables")
	}

	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URL"))
	mongoClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return fmt.Errorf("error getting connect mongo client: %v", err)
	}
	defer mongoClient.Disconnect(ctx)

	// initialize model
	model := model2.InitModel(mongoClient)

	// initialize scheduler
	//go func() {
	//	it := utility.Scheduler{Enabled: true, Job: model.GenerateReport, RevokeJob: model.RevokeUserTokens}
	//	it.Start()
	//}()

	// initialize handlers
	handler := handlers.NewHandlers(model)

	// run HTTP gateway
	go func() {
		_ = rest.RunServer(ctx, cfg.GRPCPort, cfg.HTTPPort)
	}()

	return grpc.RunServer(ctx, handler, cfg.GRPCPort)
}
