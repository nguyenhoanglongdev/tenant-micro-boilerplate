package service

import (
	"context"
	"github.com/nguyenhoanglongdev/tenant-micro-boilerplate/services/auth-service/internal/provider"
	"github.com/nguyenhoanglongdev/tenant-micro-boilerplate/services/auth-service/internal/model"
)

type AuthService struct {
	authProvider *provider.AuthProvider
}

func NewAuthService(authProvider *provider.AuthProvider) *AuthService {
	return &AuthService{authProvider: authProvider}
}

func (s *AuthService) Register(ctx context.Context, user *model.User) error {
	return s.authProvider.Register(ctx, user)
}

func (s *AuthService) Login(ctx context.Context, email, password string) (*model.User, error) {
	return s.authProvider.Login(ctx, email, password)
}
