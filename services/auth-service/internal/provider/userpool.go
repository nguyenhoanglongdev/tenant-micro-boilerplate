package provider

import "github.com/nguyenhoanglongdev/tenant-micro-boilerplate/services/auth-service/internal/model"

// UserPool defines the interface for user authentication operations.
//
// This interface is designed to be pluggable, allowing multiple implementations
// such as AWS Cognito, Firebase Authentication, or custom in-house solutions.
//
// Implementations of this interface are responsible for user registration and login flows,
// and can be injected at runtime to swap authentication backends without changing business logic.
type UserPool interface {
	Register(user *model.User) error
	Login(email, password string) (*model.User, error)
}
