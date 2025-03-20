package users

import (
	"fmt"

	"github.com/PabloPei/SmartSpend-backend/internal/auth"
)

type Service struct {
	repository UserRepository
}

func NewService(repository UserRepository) *Service {
	return &Service{repository: repository}
}

func (s *Service) RegisterUser(payload RegisterUserPayload) error {
	_, err := s.repository.GetUserByEmail(payload.Email)
	if err == nil {
		return fmt.Errorf("user with email %s already exists", payload.Email)
	}

	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}

	user := User{
		UserName: payload.UserName,
		Email:    payload.Email,
		Password: hashedPassword,
	}

	return s.repository.CreateUser(user)
}

func (s *Service) GetUserByEmail(email string) (*User, error) {
	return s.repository.GetUserByEmail(email)
}

func (s *Service) UploadPhoto(payload UploadPhotoPayload, email string) error {

	_, err := s.repository.GetUserByEmail(email)
	if err != nil {
		return fmt.Errorf("failed to upload foto %s", err)
	}

	return s.repository.UploadPhoto(payload.PhotoUrl, email)
}
