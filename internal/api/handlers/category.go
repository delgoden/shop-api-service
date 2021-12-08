package handlers

import (
	"net/http"

	"github.com/delgoden/shop-api-service/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createCategory(c *gin.Context) {
	var category models.Category
	if err := c.BindJSON(&category); err != nil {
		c.String(http.StatusBadRequest, "bad request : %v", err)
	}

	category, err := h.client.Category.Create(category)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error: %v", err)
	}

	c.JSON(http.StatusOK, category)
}

func (h *Handler) getAllCategories(c *gin.Context) {
	categories, err := h.client.Category.GetAll()
	if err != nil {
		c.String(http.StatusInternalServerError, "Error: %v", err)
	}

	c.JSON(http.StatusOK, categories)
}

func (h *Handler) updateCategory(c *gin.Context) {
	var category models.Category
	if err := c.BindJSON(&category); err != nil {
		c.String(http.StatusBadRequest, "bad request : %v", err)
	}

	category, err := h.client.Category.Update(category)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error: %v", err)
	}

	c.JSON(http.StatusOK, category)
}

func (h *Handler) deleteCategory(c *gin.Context) {
	var category models.Category
	if err := c.BindJSON(&category); err != nil {
		c.String(http.StatusBadRequest, "bad request : %v", err)
	}

	err := h.client.Category.Delete(category.ID)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error: %v", err)
	}

	c.JSON(http.StatusOK, "OK")
}
