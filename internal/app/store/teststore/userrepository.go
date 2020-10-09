package teststore

import (
	"github.com/Vladimir-Kuchinskiy/http-rest-api-golang/internal/app/store"

	"github.com/Vladimir-Kuchinskiy/http-rest-api-golang/internal/app/entity"
)

// UserRepository ...
type UserRepository struct {
	store *Store
	users map[int]*entity.User
}

// Create ...
func (r *UserRepository) Create(u *entity.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	u.ID = len(r.users) + 1
	r.users[u.ID] = u

	return nil
}

// FindByEmail ...
func (r *UserRepository) FindByEmail(email string) (*entity.User, error) {
	for _, u := range r.users {
		if u.Email == email {
			return u, nil
		}
	}

	return nil, store.ErrRecordNotFound
}

// Find ...
func (r *UserRepository) Find(id int) (*entity.User, error) {
	u, ok := r.users[id]
	if !ok {
		return nil, store.ErrRecordNotFound
	}

	return u, nil
}
