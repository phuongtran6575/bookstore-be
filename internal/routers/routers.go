package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/phuongtran6575/bookstore-be.git/internal/handler"
)

type Router struct {
	ProductHandler  *handler.ProductHandler
	CategoryHandler *handler.CategoryHandler
	AuthorHandler   *handler.AuthorHandler
}

func NewRouter(
	productHandler *handler.ProductHandler,
	categoryHandler *handler.CategoryHandler,
	authorHandler *handler.AuthorHandler,
) *Router {
	return &Router{
		ProductHandler:  productHandler,
		CategoryHandler: categoryHandler,
		AuthorHandler:   authorHandler,
	}
}

func (r *Router) RegisterRoutes() *gin.Engine {
	engine := gin.Default()

	api := engine.Group("/api/v1")

	RegisterProductRouter(api, r.ProductHandler)
	RegisterCategoryRouter(api, r.CategoryHandler)
	RegisterAuthorRouter(api, r.AuthorHandler)

	return engine
}
