package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/phuongtran6575/bookstore-be.git/internal/handler"
)

func RegisterCategoryRouter(r *gin.RouterGroup, handler *handler.CategoryHandler) {
	categoryRouter := r.Group("/categories")
	categoryRouter.GET("/", handler.GetCategories)
	categoryRouter.GET("/:id", handler.GetCategoryByID)
	categoryRouter.POST("/", handler.CreateCategory)
	categoryRouter.PUT("/:id", handler.UpdateCategory)
	categoryRouter.DELETE("/:id", handler.DeleteCategory)
}
