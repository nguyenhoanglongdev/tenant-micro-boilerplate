package provider

import (
	"context"

	"github.com/nguyenhoanglongdev/tenant-micro-boilerplate/services/auth-service/internal/model"
)

type AuthInterface interface {
	Register(ctx context.Context, user *model.User) error
	Login(ctx context.Context, email, password string) (*model.User, error)
}
