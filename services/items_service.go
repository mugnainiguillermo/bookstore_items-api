package services

import (
	"github.com/mugnainiguillermo/bookstore_items-api/items"
	"github.com/mugnainiguillermo/bookstore_utils-go/rest_errors"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, *rest_errors.RestErr)
	Get(string) (*items.Item, *rest_errors.RestErr)
}

type itemsService struct {
}

func (s *itemsService) Create(item items.Item) (*items.Item, *rest_errors.RestErr) {
	return nil, rest_errors.NewInternalServerError("implement me!", nil)
}

func (s *itemsService) Get(id string) (*items.Item, *rest_errors.RestErr) {
	return nil, rest_errors.NewInternalServerError("implement me!", nil)
}
