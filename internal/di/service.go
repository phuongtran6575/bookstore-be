package di

import (
	"github.com/phuongtran6575/bookstore-be.git/internal/services"
	serInterfaces "github.com/phuongtran6575/bookstore-be.git/internal/services/interfaces"
)

type Services struct {
	ProductService  serInterfaces.ProductService
	CategoryService serInterfaces.CategoryService
	AuthorService   serInterfaces.AuthorService
}

func NewServices(repositories *Repositories) *Services {
	return &Services{
		ProductService:  services.NewProductService(repositories.ProductRepository),
		CategoryService: services.NewCategoryService(repositories.CategoryRepository),
		AuthorService:   services.NewAuthorService(repositories.AuthorRepository),
	}
}
