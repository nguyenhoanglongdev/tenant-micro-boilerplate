package provider

import (
	"context"

	"github.com/nguyenhoanglongdev/tenant-micro-boilerplate/services/auth-service/internal/model"
)

type AuthProvider struct {
	UserPool AuthInterface
}

func NewAuthProvider(up AuthInterface) *AuthProvider {
	return &AuthProvider{UserPool: up}
}

func (a *AuthProvider) Register(ctx context.Context, user *model.User) error {
	return a.UserPool.Register(ctx, user)
}

func (a *AuthProvider) Login(ctx context.Context, email, password string) (*model.User, error) {
	return a.UserPool.Login(ctx, email, password)
}
