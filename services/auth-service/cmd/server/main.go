package main

import (
	"fmt"
	"log"
	"context"

	"github.com/joho/godotenv"
	"github.com/nguyenhoanglongdev/tenant-micro-boilerplate/services/auth-service/internal/config"
	"github.com/nguyenhoanglongdev/tenant-micro-boilerplate/services/auth-service/internal/logger"
	"github.com/nguyenhoanglongdev/tenant-micro-boilerplate/services/auth-service/internal/util"
	"github.com/nguyenhoanglongdev/tenant-micro-boilerplate/services/auth-service/internal/provider/cognito"
	// "github.com/nguyenhoanglongdev/tenant-micro-boilerplate/services/auth-service/internal/router"
	"github.com/nguyenhoanglongdev/tenant-micro-boilerplate/services/auth-service/internal/service"
	// "github.com/nguyenhoanglongdev/tenant-micro-boilerplate/services/auth-service/internal/handler"
)

func main() {
	// Load .env file, ignore error if not present (e.g., in prod)
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, relying on environment variables")
	}

	fmt.Println("Load env has been done!")

	// Runtime logging config
	util.EnsureFolder("./runtime/log")

	logFilePath := "./runtime/log/app.log"
	logFile := true

	// runtime logger init instance
	logger, err := logger.NewLogger(logFile, logFilePath)
	if err != nil {
		log.Fatalf("cannot initialize zap logger: %v", err)
	}
	defer logger.Sync()

	fmt.Println("Init logger has been done!")

	cfg := config.LoadConfig()

	util.PrettyPrint(cfg)

	// Init context
	ctx := context.Background()

	authProvider, err := cognito.NewAuthProvider(ctx, cfg.AwsRegion, cfg.UserPoolId, cfg.ClientId)
	if err != nil {
		log.Fatalf("failed to create Cognito auth provider: %v", err)
	}

	authService := service.NewAuthService(authProvider)


    // authHandler := handler.NewAuthHandler(authService)
    // router := router.SetupRouter(authHandler)

    // fmt.Println("Starting server at :8080")
    // if err := router.Run(":8080"); err != nil {
    //     log.Fatalf("failed to run server: %v", err)
    // }
}
