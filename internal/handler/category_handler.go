package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phuongtran6575/bookstore-be.git/internal/models"
	serInterfaces "github.com/phuongtran6575/bookstore-be.git/internal/services/interfaces"
)

type CategoryHandler struct {
	categoryService serInterfaces.CategoryService
}

func NewCategoryHandler(categoryService serInterfaces.CategoryService) *CategoryHandler {
	return &CategoryHandler{categoryService: categoryService}
}

func (h *CategoryHandler) GetCategories(c *gin.Context) {
	categories := h.categoryService.GetCategories()
	c.JSON(200, gin.H{
		"message":    "Get Categories",
		"categories": categories,
	})
}

func (h *CategoryHandler) GetCategoryByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid ID",
		})
		return
	}
	category := h.categoryService.GetCategoryByID(uint(id))
	c.JSON(200, gin.H{
		"message":  "Get Category By ID",
		"category": category,
	})
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid Category",
		})
		return
	}
	h.categoryService.CreateCategory(category)
	c.JSON(200, gin.H{
		"message":  "Create Category",
		"category": category,
	})
}

func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid ID",
		})
		return
	}
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid Category",
		})
		return
	}
	category.ID = uint(id)
	h.categoryService.UpdateCategory(category)
	c.JSON(200, gin.H{
		"message":  "Update Category",
		"category": category,
	})
}

func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid ID",
		})
		return
	}
	h.categoryService.DeleteCategory(uint(id))
	c.JSON(200, gin.H{
		"message": "Delete Category",
	})
}
