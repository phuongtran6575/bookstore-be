package handler

import (
	"github.com/gin-gonic/gin"
	serInterfaces "github.com/phuongtran6575/bookstore-be.git/internal/services/interfaces"
)

type CategoryHandler struct {
	categoryService serInterfaces.CategoryService
}

func NewCategoryHandler(categoryService serInterfaces.CategoryService) *CategoryHandler {
	return &CategoryHandler{categoryService: categoryService}
}

func (h *CategoryHandler) GetCategories(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get Categories",
	})
}

func (h *CategoryHandler) GetCategoryByID(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get Category By ID",
	})
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Create Category",
	})
}

func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Update Category",
	})
}

func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Delete Category",
	})
}
