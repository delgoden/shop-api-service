package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/delgoden/shop-api-service/internal/models"
	"github.com/delgoden/shop-api-service/pkg/rest"
)

type CategoriesClient struct {
	base     rest.BaseClient
	Resource int
}

const (
	baseURL   = "http://127.0.0.1:8002"
	urlCreate = "/api/categories"
)

func NewCategoriesClient() *CategoriesClient {
	return &CategoriesClient{base: rest.BaseClient{
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}}
}

func (c *CategoriesClient) Create(category models.Category) (models.Category, error) {
	data, err := json.Marshal(&category)
	if err != nil {
		return category, err
	}

	req, err := http.NewRequest("POST", baseURL+urlCreate, bytes.NewBuffer(data))
	if err != nil {
		return category, err
	}

	response, err := c.base.SendRequest(req)
	if err != nil {
		return category, err
	}

	if response.StatusCode != 200 {
		return category, err
	}

	err = json.NewDecoder(response.Body).Decode(&category)
	if err != nil {
		return category, err
	}

	return category, err

}

func (c *CategoriesClient) GetAll() ([]models.Category, error) {
	var categories []models.Category
	req, err := http.NewRequest("POST", baseURL+urlCreate, bytes.NewBuffer([]byte{}))
	if err != nil {
		return categories, err
	}

	response, err := c.base.SendRequest(req)
	if err != nil {
		return categories, err
	}

	if response.StatusCode != 200 {
		return categories, err
	}

	err = json.NewDecoder(response.Body).Decode(&categories)
	if err != nil {
		return categories, err
	}

	return categories, err
}

func (c *CategoriesClient) GetByID(ID int) (models.Category, error) {
	category := models.Category{
		ID: ID,
	}
	data, err := json.Marshal(&category)
	if err != nil {
		return category, err
	}

	req, err := http.NewRequest("POST", baseURL+urlCreate, bytes.NewBuffer(data))
	if err != nil {
		return category, err
	}

	response, err := c.base.SendRequest(req)
	if err != nil {
		return category, err
	}

	if response.StatusCode != 200 {
		return category, err
	}

	err = json.NewDecoder(response.Body).Decode(&category)
	if err != nil {
		return category, err
	}

	return category, err
}

func (c *CategoriesClient) Update(category models.Category) (models.Category, error) {
	data, err := json.Marshal(&category)
	if err != nil {
		return category, err
	}

	req, err := http.NewRequest("POST", baseURL+urlCreate, bytes.NewBuffer(data))
	if err != nil {
		return category, err
	}

	response, err := c.base.SendRequest(req)
	if err != nil {
		return category, err
	}

	if response.StatusCode != 200 {
		return category, err
	}

	err = json.NewDecoder(response.Body).Decode(&category)
	if err != nil {
		return category, err
	}

	return category, err
}

func (c *CategoriesClient) Delete(ID int) error {
	category := models.Category{
		ID: ID,
	}
	data, err := json.Marshal(&category)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", baseURL+urlCreate, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	response, err := c.base.SendRequest(req)
	if err != nil {
		return err
	}

	if response.StatusCode != 200 {
		return err
	}

	return err
}
