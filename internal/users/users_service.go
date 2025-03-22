package users

import (
	"github.com/PabloPei/SmartSpend-backend/conf"
	"github.com/PabloPei/SmartSpend-backend/internal/auth"
	"github.com/PabloPei/SmartSpend-backend/internal/errors"
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
		return errors.ErrUserAlreadyExist(payload.Email)
	}

	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		return errors.ErrHashingPassword(err)
	}

	user := User{
		UserName: payload.UserName,
		Email:    payload.Email,
		Password: hashedPassword,
	}

	return s.repository.CreateUser(user)
}

func (s *Service) LogInUser(user LogInUserPayload) (string, error) {

	u, err := s.repository.GetUserByEmail(user.Email)
	if err != nil {
		return "", errors.ErrUserNotFound(user.Email)
	}

	if !auth.ComparePasswords(u.Password, []byte(user.Password)) {
		return "", errors.ErrInvalidCredentials
	}

	userJWT := createJWTPayload(*u)

	secret := []byte(conf.ServerConfig.JWTSecret)
	token, err := auth.CreateJWT(secret, userJWT)
	if err != nil {
		return "", errors.ErrJWTCreation
	}

	return token, nil
}

func (s *Service) UploadPhoto(payload UploadPhotoPayload, email string) error {

	_, err := s.repository.GetUserByEmail(email)
	if err != nil {
		return errors.ErrUploadPhoto
	}

	return s.repository.UploadPhoto(payload.PhotoUrl, email)
}

func (s *Service) GetUserByEmail(email string) (*User, error) {
	return s.repository.GetUserByEmail(email)
}

// Aux Functions

func createJWTPayload(user User) auth.UserJWT {

	var userJWT auth.UserJWT

	userJWT.UserId = string(user.UserId)
	userJWT.Email = user.Email
	userJWT.UserName = user.UserName

	return userJWT

}
