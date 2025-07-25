package cognito

import (
    "errors"
    "strings"
)

var (
    ErrUserExists        = errors.New("user already exists")
    ErrInvalidPassword   = errors.New("invalid password")
    ErrUserNotConfirmed  = errors.New("user not confirmed")
    ErrUserNotFound      = errors.New("user not found")
)

func ParseError(err error) error {
    if err == nil {
        return nil
    }
    msg := err.Error()
    switch {
    case strings.Contains(msg, "UsernameExistsException"):
        return ErrUserExists
    case strings.Contains(msg, "InvalidPasswordException"):
        return ErrInvalidPassword
    case strings.Contains(msg, "UserNotConfirmedException"):
        return ErrUserNotConfirmed
    case strings.Contains(msg, "UserNotFoundException"):
        return ErrUserNotFound
    default:
        return err
    }
}
