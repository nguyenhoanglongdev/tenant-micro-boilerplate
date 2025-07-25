package service

import (
    "context"
    "github.com/nguyenhoanglongdev/tenant-micro-boilerplate/services/auth-service/internal/provider/cognito"
)

type AuthService struct {
    userPool cognito.UserPool
}

func NewAuthService(userPool cognito.UserPool) *AuthService {
    return &AuthService{userPool: userPool}
}

func (s *AuthService) Register(ctx context.Context, username, password string) error {
    return s.userPool.Register(ctx, username, password)
}

func (s *AuthService) Login(ctx context.Context, username, password string) (string, error) {
    return s.userPool.Login(ctx, username, password)
}
