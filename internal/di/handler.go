package di

import "github.com/phuongtran6575/bookstore-be.git/internal/handler"

type Handlers struct {
	ProductHandler  *handler.ProductHandler
	CategoryHandler *handler.CategoryHandler
	AuthorHandler   *handler.AuthorHandler
}

func NewHandlers(services *Services) *Handlers {
	return &Handlers{
		ProductHandler:  handler.NewProductHandler(services.ProductService),
		CategoryHandler: handler.NewCategoryHandler(services.CategoryService),
		AuthorHandler:   handler.NewAuthorHandler(services.AuthorService),
	}
}
