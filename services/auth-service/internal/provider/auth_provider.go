package provider

import "github.com/nguyenhoanglongdev/tenant-micro-boilerplate/services/auth-service/internal/model"

// AuthProvider defines the interface for user authentication operations.
//
// It allows pluggable backends such as AWS Cognito, Firebase Authentication, or custom solutions.
//
// Implementations handle user registration and authentication,
// and can be swapped without modifying business logic.
type AuthProvider interface {
	RegisterUser(user *model.User) error
	Authenticate(email, password string) (*model.User, error)
}
