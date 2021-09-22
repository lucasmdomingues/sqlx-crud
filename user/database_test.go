package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	user := &User{
		Username: "lucasmdomingues",
		Password: "123",
	}

	err := Create(user)
	assert.Nil(t, err)
}

func TestFindUsers(t *testing.T) {
	users, err := FindUsers()
	assert.Nil(t, err)

	assert.Greater(t, len(users), 0)

}

func TestUpdate(t *testing.T) {
	user := &User{
		ID:       1,
		Username: "lucasmdomingues",
		Password: "123",
	}

	err := Update(user)
	assert.Nil(t, err)
}

func TestGetByID(t *testing.T) {
	user, err := GetByID(1)
	assert.Nil(t, err)

	assert.NotNil(t, user)
}

func TestDelete(t *testing.T) {
	err := Delete(1)
	assert.Nil(t, err)
}
