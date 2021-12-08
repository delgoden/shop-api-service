package service

import (
	"github.com/delgoden/shop-api-service/internal/models"
	"github.com/delgoden/shop-api-service/pkg/rest"
)

type Category interface {
	Create(category models.Category) (models.Category, error)
	GetAll() ([]models.Category, error)
	GetByID(ID int) (models.Category, error)
	Update(category models.Category) (models.Category, error)
	Delete(ID int) error
}

type Client struct {
	Category
}

func NewClient(cli *rest.BaseClient) *Client {
	return &Client{
		Category: NewCategoriesClient(),
	}
}
