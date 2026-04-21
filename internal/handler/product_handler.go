package handler

import (
	"github.com/gin-gonic/gin"
	serInterfaces "github.com/phuongtran6575/bookstore-be.git/internal/services/interfaces"
)

type ProductHandler struct {
	productService serInterfaces.ProductService
}

func NewProductHandler(productService serInterfaces.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get Products",
	})
}

func (h *ProductHandler) GetProductByID(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get Product By ID",
	})
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Create Product",
	})
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Update Product",
	})
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Delete Product",
	})

}
