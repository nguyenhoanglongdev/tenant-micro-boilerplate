package service

import (
	"context"
)

type AuthService struct {
	userPool UserPool
}

func NewAuthService(userPool UserPool) *AuthService {
	return &AuthService{userPool: userPool}
}

func (s *AuthService) Register(ctx context.Context, username, password string) error {
	return s.userPool.Register(ctx, username, password)
}

func (s *AuthService) Login(ctx context.Context, username, password string) (string, error) {
	return s.userPool.Login(ctx, username, password)
}
