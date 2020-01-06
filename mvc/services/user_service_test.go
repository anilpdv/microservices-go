package services

import (
	"microservicego/mvc/domain"
	"microservicego/mvc/utils"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	userDaoMock     usersDaoMock
	getUserFunction func(userid int64) (*domain.User, *utils.ApplicationError)
)

func init() {
	domain.UserDto = &usersDaoMock{}
}

type usersDaoMock struct{}

func (m *usersDaoMock) GetUser(userid int64) (*domain.User, *utils.ApplicationError) {
	return getUserFunction(userid)
}

func TestUserNotfound(t *testing.T) {
	getUserFunction = func(userid int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			StatusCode: http.StatusNotFound,
			Message:    "user not found",
		}
	}
	user, err := UserService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)

	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "user not found", err.Message)
}

func TestUserfound(t *testing.T) {
	getUserFunction = func(userid int64) (*domain.User, *utils.ApplicationError) {
		return &domain.User{
			Email:     "pdvanil@gmail.com",
			Firstname: "anil",
			Lastname:  "pdv",
		}, nil
	}
	user, err := UserService.GetUser(123)
	assert.NotNil(t, user)
	assert.Nil(t, err)

	assert.EqualValues(t, "pdvanil@gmail.com", user.Email)

}
