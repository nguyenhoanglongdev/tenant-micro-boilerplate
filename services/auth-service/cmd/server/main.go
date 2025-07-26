package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/nguyenhoanglongdev/tenant-micro-boilerplate/services/auth-service/internal/logger"
	"github.com/nguyenhoanglongdev/tenant-micro-boilerplate/services/auth-service/internal/util"
	// "github.com/nguyenhoanglongdev/tenant-micro-boilerplate/services/auth-service/internal/config"
	// "github.com/nguyenhoanglongdev/tenant-micro-boilerplate/services/auth-service/internal/provider/cognito"
	// "github.com/nguyenhoanglongdev/tenant-micro-boilerplate/services/auth-service/internal/router"
	// "github.com/nguyenhoanglongdev/tenant-micro-boilerplate/services/auth-service/internal/service"
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

	logger.Info("Init logger has been done")

	// cfg := config.LoadConfig()
	// ctx := context.Background()

	// client, err := cognito.NewClient(ctx, cfg.AwsRegion)
	// if err != nil {
	//     log.Fatalf("failed to create Cognito client: %v", err)
	// }

	// userPool := cognito.NewUserPool(client, cfg.UserPoolId, cfg.ClientId)
	// authService := service.NewAuthService(userPool)
	// authHandler := handler.NewAuthHandler(authService)

	// r := router.SetupRouter(authHandler)

	// if cfg.Port == "" {
	//     cfg.Port = "8080"
	// }
	// log.Printf("starting server on :%s", cfg.Port)
	// if err := r.Run(":" + cfg.Port); err != nil {
	//     log.Fatalf("server exited with error: %v", err)
	// }
}
