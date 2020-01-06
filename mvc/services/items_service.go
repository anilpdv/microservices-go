package services

import (
	"microservicego/mvc/domain"
	"microservicego/mvc/utils"
	"net/http"
)

type itemService struct{}

// ItemService : type of itemService
var (
	ItemService itemService
)

// GetItem : get item implements,getting items.
func (i *itemService) GetItem(itemID string) (*domain.Item, *utils.ApplicationError) {
	itemError := new(utils.ApplicationError)
	itemError.ApplicationErrorCreator("implement me", http.StatusInternalServerError, "server_request")
	return nil, itemError
}
