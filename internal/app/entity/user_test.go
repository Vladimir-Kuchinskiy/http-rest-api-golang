package entity_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Vladimir-Kuchinskiy/http-rest-api-golang/internal/app/entity"
)

func TestUser_BeforeCreate(t *testing.T) {
	u := entity.TestUser(t)

	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.EncryptedPassword)
}
