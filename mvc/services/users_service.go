package services

import (
	"microservicego/mvc/domain"
	"microservicego/mvc/utils"
)

type userService struct{}

var (
	UserService userService
)

// GetUser : get user gets the user
func (u *userService) GetUser(userid int64) (*domain.User, *utils.ApplicationError) {
	return domain.UserDto.GetUser(userid)
}
