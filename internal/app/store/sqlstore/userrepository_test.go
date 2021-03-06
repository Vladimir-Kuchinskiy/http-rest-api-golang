package sqlstore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Vladimir-Kuchinskiy/http-rest-api-golang/internal/app/entity"
	"github.com/Vladimir-Kuchinskiy/http-rest-api-golang/internal/app/store"
	"github.com/Vladimir-Kuchinskiy/http-rest-api-golang/internal/app/store/sqlstore"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u := entity.TestUser(t)
	err := s.User().Create(u)

	assert.NoError(t, err)
	assert.NotNil(t, u.ID)
	assert.NotEmpty(t, u.EncryptedPassword)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	email := "user@example.com"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := entity.TestUser(t)
	u.Email = email
	s.User().Create(u)

	u, err = s.User().FindByEmail(email)

	assert.NoError(t, err)
	assert.NotNil(t, u)
	assert.Equal(t, u.Email, email)
}

func TestUserRepository_Find(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)

	u := entity.TestUser(t)
	s.User().Create(u)

	u2, err := s.User().Find(u.ID)
	assert.NoError(t, err)
	assert.NotNil(t, u2)
	assert.Equal(t, u.ID, u2.ID)
}
