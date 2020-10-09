package teststore

import (
	"github.com/Vladimir-Kuchinskiy/http-rest-api-golang/internal/app/store"

	"github.com/Vladimir-Kuchinskiy/http-rest-api-golang/internal/app/entity"
)

// UserRepository ...
type UserRepository struct {
	store *Store
	users map[string]*entity.User
}

// Create ...
func (r *UserRepository) Create(u *entity.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	r.users[u.Email] = u
	u.ID = len(r.users)

	return nil
}

// FindByEmail ...
func (r *UserRepository) FindByEmail(email string) (*entity.User, error) {
	u, ok := r.users[email]

	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return u, nil
}
