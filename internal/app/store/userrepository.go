package store

import "github.com/Vladimir-Kuchinskiy/http-rest-api-golang/internal/app/entity"

// UserRepository ...
type UserRepository struct {
	store *Store
}

// Create ...
func (r *UserRepository) Create(u *entity.User) (*entity.User, error) {
	if err := u.Validate(); err != nil {
		return nil, err
	}

	if err := u.BeforeCreate(); err != nil {
		return nil, err
	}

	if err := r.store.db.QueryRow(
		"insert into users (email, encrypted_password) values ($1, $2) returning id",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID); err != nil {
		return nil, err
	}
	return u, nil
}

// FindByEmail ...
func (r *UserRepository) FindByEmail(email string) (*entity.User, error) {
	u := &entity.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE email = $1",
		email,
	).Scan(&u.ID, &u.Email, &u.EncryptedPassword); err != nil {
		return nil, err
	}

	return u, nil
}
