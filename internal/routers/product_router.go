package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/phuongtran6575/bookstore-be.git/internal/handler"
)

func RegisterProductRouter(r *gin.RouterGroup, handler *handler.ProductHandler) {
	productRouter := r.Group("/products")
	productRouter.GET("/", handler.GetProducts)
	productRouter.GET("/:id", handler.GetProductByID)
	productRouter.POST("/", handler.CreateProduct)
	productRouter.PUT("/:id", handler.UpdateProduct)
	productRouter.DELETE("/:id", handler.DeleteProduct)
}
