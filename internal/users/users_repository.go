package users

import (
	"database/sql"
	"fmt"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

// CreateUser inserta un nuevo usuario en la base de datos.
func (s *Store) CreateUser(user User) error {
	_, err := s.db.Exec(
		"INSERT INTO auth.\"user\" (user_name, email, password) VALUES ($1, $2, $3)",
		user.UserName, user.Email, user.Password,
	)
	if err != nil {
		return fmt.Errorf("error al crear usuario: %w", err)
	}
	return nil
}

// GetUserByEmail obtiene un usuario por su email.
func (s *Store) GetUserByEmail(email string) (*User, error) {
	row := s.db.QueryRow("SELECT * FROM auth.\"user\" WHERE email = $1", email)
	return scanRowIntoUser(row)
}

// GetUserByID obtiene un usuario por su ID.
func (s *Store) GetUserByID(id int) (*User, error) {
	row := s.db.QueryRow("SELECT * FROM auth.\"user\" WHERE user_id = $1", id)
	return scanRowIntoUser(row)
}

// scanRowIntoUser mapea los datos de una fila a una estructura User.
func scanRowIntoUser(row *sql.Row) (*User, error) {
	user := new(User)
	err := row.Scan(
		&user.UserId,
		&user.UserName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("usuario no encontrado")
		}
		return nil, fmt.Errorf("error al escanear usuario: %w", err)
	}
	return user, nil
}
