package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phuongtran6575/bookstore-be.git/internal/models"
	serInterfaces "github.com/phuongtran6575/bookstore-be.git/internal/services/interfaces"
)

type AuthorHandler struct {
	authorService serInterfaces.AuthorService
}

func NewAuthorHandler(authorService serInterfaces.AuthorService) *AuthorHandler {
	return &AuthorHandler{authorService: authorService}
}

func (h *AuthorHandler) GetAuthors(c *gin.Context) {
	authors := h.authorService.GetAuthors()
	c.JSON(200, gin.H{
		"message": "Get Authors",
		"authors": authors,
	})
}

func (h *AuthorHandler) GetAuthorByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid ID",
		})
		return
	}
	author := h.authorService.GetAuthorByID(uint(id))
	c.JSON(200, gin.H{
		"message": "Get Author By ID",
		"author":  author,
	})
}

func (h *AuthorHandler) CreateAuthor(c *gin.Context) {
	var author models.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid Author",
		})
		return
	}
	h.authorService.CreateAuthor(author)
	c.JSON(200, gin.H{
		"message": "Create Author",
		"author":  author,
	})
}

func (h *AuthorHandler) UpdateAuthor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid ID",
		})
		return
	}
	var author models.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid Author",
		})
		return
	}
	author.ID = uint(id)
	h.authorService.UpdateAuthor(author)
	c.JSON(200, gin.H{
		"message": "Update Author",
		"author":  author,
	})
}

func (h *AuthorHandler) DeleteAuthor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid ID",
		})
		return
	}
	h.authorService.DeleteAuthor(uint(id))
	c.JSON(200, gin.H{
		"message": "Delete Author",
	})
}
