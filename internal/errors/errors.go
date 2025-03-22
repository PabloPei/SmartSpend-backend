package errors

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrJWTCreation        = errors.New("unable to create JWT token")
	ErrUploadPhoto        = errors.New("unable to upload photo")
	ErrUserNotFound       = func(email string) error {
		return fmt.Errorf("user with email %s not found", email)
	}
	ErrUserAlreadyExist = func(email string) error {
		return fmt.Errorf("user with email %s already exists", email)
	}
	ErrHashingPassword = func(hashError error) error {
		return fmt.Errorf("failed to hash password: %v", hashError)
	}
)
