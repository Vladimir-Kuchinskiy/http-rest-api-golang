package sqlstore

import (
	"database/sql"

	"github.com/Vladimir-Kuchinskiy/http-rest-api-golang/internal/app/entity"
	"github.com/Vladimir-Kuchinskiy/http-rest-api-golang/internal/app/store"
)

// UserRepository ...
type UserRepository struct {
	store *Store
}

// Create ...
func (r *UserRepository) Create(u *entity.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	return r.store.db.QueryRow(
		"insert into users (email, encrypted_password) values ($1, $2) returning id",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID)
}

// FindByEmail ...
func (r *UserRepository) FindByEmail(email string) (*entity.User, error) {
	u := &entity.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE email = $1",
		email,
	).Scan(&u.ID, &u.Email, &u.EncryptedPassword); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return u, nil
}

// Find ...
func (r *UserRepository) Find(id int) (*entity.User, error) {
	u := &entity.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE id = $1",
		id,
	).Scan(&u.ID, &u.Email, &u.EncryptedPassword); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, err
	}

	return u, nil
}
