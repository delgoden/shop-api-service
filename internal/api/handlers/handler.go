package handlers

import (
	"github.com/delgoden/shop-api-service/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	client *service.Client
}

func NewHandler(client *service.Client) *Handler {
	return &Handler{client: client}
}

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")

	{
		categories := api.Group("/categories")
		{
			categories.POST("/", h.createCategory)
			categories.GET("/", h.getAllCategories)
			categories.PUT("/:id", h.updateCategory)
			categories.DELETE("/:id", h.deleteCategory)
		}
	}

	return router
}
