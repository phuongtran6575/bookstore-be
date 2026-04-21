package handler

import (
	"github.com/gin-gonic/gin"
	serInterfaces "github.com/phuongtran6575/bookstore-be.git/internal/services/interfaces"
)

type AuthorHandler struct {
	authorService serInterfaces.AuthorService
}

func NewAuthorHandler(authorService serInterfaces.AuthorService) *AuthorHandler {
	return &AuthorHandler{authorService: authorService}
}

func (h *AuthorHandler) GetAuthors(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get Authors",
	})
}

func (h *AuthorHandler) GetAuthorByID(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get Author By ID",
	})
}

func (h *AuthorHandler) CreateAuthor(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Create Author",
	})
}

func (h *AuthorHandler) UpdateAuthor(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Update Author",
	})
}

func (h *AuthorHandler) DeleteAuthor(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Delete Author",
	})
}
