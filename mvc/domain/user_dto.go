package domain

import (
	"fmt"
	"microservicego/mvc/utils"
	"net/http"
)

var users = map[int64]*User{
	123: {ID: 123, Firstname: "anil", Lastname: "pdv", Email: "pdvanil@gmail.com"},
}

// UserDto : struct of type userDto
var UserDto userDtoInterface

type userDto struct {
}

type userDtoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

// GetUser : get user gets the users
func (u *userDto) GetUser(userid int64) (*User, *utils.ApplicationError) {
	appErr := new(utils.ApplicationError)

	user := users[userid]
	fmt.Println(user)
	if user != nil {
		return user, nil
	}

	appErr.ApplicationErrorCreator("user not found", http.StatusNotFound, "not_found")
	return nil, appErr

}
