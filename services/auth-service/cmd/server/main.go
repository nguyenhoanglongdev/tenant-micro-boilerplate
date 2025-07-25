package main

import (
    "context"
    "log"
    "github.com/nguyenhoanglongdev/tenant-micro-boilerplate/services/auth-service/internal/config"
    "github.com/nguyenhoanglongdev/tenant-micro-boilerplate/services/auth-service/internal/provider/cognito"
    "github.com/nguyenhoanglongdev/tenant-micro-boilerplate/services/auth-service/internal/router"
    "github.com/nguyenhoanglongdev/tenant-micro-boilerplate/services/auth-service/internal/service"
    "github.com/nguyenhoanglongdev/tenant-micro-boilerplate/services/auth-service/internal/handler"
)

func main() {
    cfg := config.LoadConfig()
    ctx := context.Background()

    client, err := cognito.NewClient(ctx, cfg.AwsRegion)
    if err != nil {
        log.Fatalf("failed to create Cognito client: %v", err)
    }

    userPool := cognito.NewUserPool(client, cfg.UserPoolId, cfg.ClientId)
    authService := service.NewAuthService(userPool)
    authHandler := handler.NewAuthHandler(authService)

    r := router.SetupRouter(authHandler)

    if cfg.Port == "" {
        cfg.Port = "8080"
    }
    log.Printf("starting server on :%s", cfg.Port)
    if err := r.Run(":" + cfg.Port); err != nil {
        log.Fatalf("server exited with error: %v", err)
    }
}