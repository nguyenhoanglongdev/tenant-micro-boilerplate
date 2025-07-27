package main

import (
	"context"
	"fmt"
	"log"

	"go.uber.org/zap"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/joho/godotenv"

	"github.com/nguyenhoanglongdev/tenant-micro-boilerplate/services/auth-service/internal/cloud"
	"github.com/nguyenhoanglongdev/tenant-micro-boilerplate/services/auth-service/internal/config"
	"github.com/nguyenhoanglongdev/tenant-micro-boilerplate/services/auth-service/internal/handler"
	"github.com/nguyenhoanglongdev/tenant-micro-boilerplate/services/auth-service/internal/logger"
	"github.com/nguyenhoanglongdev/tenant-micro-boilerplate/services/auth-service/internal/provider"
	"github.com/nguyenhoanglongdev/tenant-micro-boilerplate/services/auth-service/internal/provider/cognito"
	"github.com/nguyenhoanglongdev/tenant-micro-boilerplate/services/auth-service/internal/router"
	"github.com/nguyenhoanglongdev/tenant-micro-boilerplate/services/auth-service/internal/service"
	"github.com/nguyenhoanglongdev/tenant-micro-boilerplate/services/auth-service/internal/util"
)

func main() {
	loadEnv()
	ctx := context.Background()
	appConfig := config.LoadConfig()
	util.PrettyPrint(appConfig)

	log := initLogger()
	defer log.Sync()

	awsCfg := initAWS(ctx, appConfig)

	authHandler := initAuthHandler(ctx, appConfig, awsCfg)
	r := router.SetupRouter(authHandler, log)

	fmt.Println("Starting server at :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("failed to run server:")
	}
}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env not found, using environment variables")
	}
	util.EnsureFolder("./runtime/log")
}

func initLogger() *zap.Logger {
	logFile := true
	path := "./runtime/log/app.log"

	log, err := logger.NewLogger(logFile, path)
	if err != nil {
		fmt.Println("cannot initialize logger: %v", err)
	}
	fmt.Println("Logger initialized")
	return log
}

func initAWS(ctx context.Context, cfg *config.Config) aws.Config {
	awsCfg, err := cloud.LoadAWSConfig(ctx, cfg.AwsRegion)
	if err != nil {
		log.Fatalf("Failed to load AWS config: %v", err)
	}
	fmt.Printf("AWS Region: %s\n", awsCfg.Region)
	return awsCfg
}

func initAuthHandler(ctx context.Context, cfg *config.Config, awsCfg aws.Config) *handler.AuthHandler {
	cognitoClient, err := cognito.NewClient(ctx, awsCfg, cfg.UserPoolId)
	if err != nil {
		log.Fatalf("Failed to create Cognito client: %v", err)
	}

	userPool := cognito.NewCognitoUserPool(cognitoClient, cfg.UserPoolId, cfg.ClientId)
	authProvider := provider.NewAuthProvider(userPool)
	authService := service.NewAuthService(authProvider)
	return handler.NewAuthHandler(authService)
}
