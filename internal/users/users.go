package users

import (
	"time"
)

type User struct {
	UserId     []uint8   `json:"userId"`
	UserName   string    `json:"userName"`
	PhotoUrl   string    `json:"photoUrl"`
	Email      string    `json:"email"`
	Password   string    `json:"-"`
	LanguageId string    `json:"languageId"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type UserRepository interface {
	GetUserByEmail(email string) (*User, error)
	CreateUser(User) error
	UploadPhoto(photoUrl string, email string) error
}

type RegisterUserPayload struct {
	UserName string `json:"userName" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3,max=130"`
}

type UploadPhotoPayload struct {
	PhotoUrl string `json:"photoUrl" validate:"required,uri"`
}
