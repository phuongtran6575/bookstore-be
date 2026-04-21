package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/phuongtran6575/bookstore-be.git/internal/handler"
)

func RegisterAuthorRouter(r *gin.RouterGroup, handler *handler.AuthorHandler) {
	authorRouter := r.Group("/authors")
	authorRouter.GET("/", handler.GetAuthors)
	authorRouter.GET("/:id", handler.GetAuthorByID)
	authorRouter.POST("/", handler.CreateAuthor)
	authorRouter.PUT("/:id", handler.UpdateAuthor)
	authorRouter.DELETE("/:id", handler.DeleteAuthor)
}
