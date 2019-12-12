package services

import (
	"microservicego/mvc/domain"
	"microservicego/mvc/utils"
)

// GetUser : get user gets the user
func GetUser(userid int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userid)
}
