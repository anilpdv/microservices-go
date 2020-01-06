package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := UserDto.GetUser(0)
	assert.Nil(t, user, "we are not expecting a user with id 0")
	assert.NotNil(t, err, "we are expecting an error when user id is 0")

	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "user not found", err.Message)
	assert.EqualValues(t, "not_found", err.Code)
}

func TestGetUserNoError(t *testing.T) {
	user, err := UserDto.GetUser(123)

	assert.Nil(t, err, "we are expecting no error")
	assert.NotNil(t, user, "we are expecting data from the user")

	assert.EqualValues(t, 123, user.ID)
	assert.EqualValues(t, "anil", user.Firstname)
	assert.EqualValues(t, "pdv", user.Lastname)
	assert.EqualValues(t, "pdvanil@gmail.com", user.Email)
}
