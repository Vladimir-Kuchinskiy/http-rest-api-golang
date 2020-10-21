package users

import (
	"database/sql"

	_ "github.com/lib/pq" // ...

	"github.com/Vladimir-Kuchinskiy/http-rest-api-golang/internal/database"
)

// RepositoryI ...
type RepositoryI interface {
	Create(*User) error
	FindByEmail(string) (*User, error)
	Find(int) (*User, error)
}

// Repository ...
type Repository struct {
	db *sql.DB
}

// NewRepository ...
func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// Create ...
func (r *Repository) Create(u *User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	return r.db.QueryRow(
		"insert into users (email, encrypted_password) values ($1, $2) returning id",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID)
}

// FindByEmail ...
func (r *Repository) FindByEmail(email string) (*User, error) {
	u := &User{}
	if err := r.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE email = $1",
		email,
	).Scan(&u.ID, &u.Email, &u.EncryptedPassword); err != nil {
		if err == sql.ErrNoRows {
			return nil, database.ErrRecordNotFound
		}

		return nil, err
	}

	return u, nil
}

// Find ...
func (r *Repository) Find(id int) (*User, error) {
	u := &User{}
	if err := r.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE id = $1",
		id,
	).Scan(&u.ID, &u.Email, &u.EncryptedPassword); err != nil {
		if err == sql.ErrNoRows {
			return nil, database.ErrRecordNotFound
		}

		return nil, err
	}

	return u, nil
}
