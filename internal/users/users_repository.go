package users

import (
	"database/sql"
	"fmt"
)

type SQLRepository struct {
	db *sql.DB
}

func NewSQLRepository(db *sql.DB) *SQLRepository {
	return &SQLRepository{db: db}
}

func (s *SQLRepository) CreateUser(user User) error {
	_, err := s.db.Exec(
		"INSERT INTO auth.\"user\" (user_name, email, password) VALUES ($1, $2, $3)",
		user.UserName, user.Email, user.Password,
	)
	if err != nil {
		return fmt.Errorf("error al crear usuario: %w", err)
	}

	return nil
}

func (s *SQLRepository) UploadPhoto(photoUrl string, email string) error {

	_, err := s.db.Exec(
		"UPDATE auth.\"user\" SET photo_url = $1 WHERE email = $2",
		photoUrl, email,
	)

	if err != nil {
		return fmt.Errorf("error al actualizar la foto: %w", err)
	}

	return nil
}

func (s *SQLRepository) GetUserByEmail(email string) (*User, error) {
	row := s.db.QueryRow("SELECT user_id, user_name, email, password FROM auth.\"user\" WHERE email = $1", email)
	return scanRowIntoUser(row)
}

// scanRowIntoUser mapea los datos de una fila a una estructura User. Solucionar que devuelva todos los campos y no se rompa si esta null
func scanRowIntoUser(row *sql.Row) (*User, error) {
	user := new(User)
	err := row.Scan(
		&user.UserId,
		&user.UserName,
		&user.Email,
		&user.Password,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("usuario no encontrado")
		}
		return nil, fmt.Errorf("error al escanear usuario: %w", err)
	}
	return user, nil
}
