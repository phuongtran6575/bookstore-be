package di

import (
	"github.com/phuongtran6575/bookstore-be.git/internal/repositories"
	repoInterfaces "github.com/phuongtran6575/bookstore-be.git/internal/repositories/interfaces"
	"gorm.io/gorm"
)

type Repositories struct {
	ProductRepository  repoInterfaces.ProductRepository
	CategoryRepository repoInterfaces.CategoryRepository
	AuthorRepository   repoInterfaces.AuthorRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		ProductRepository:  repositories.NewProductRepository(db),
		CategoryRepository: repositories.NewCategoryRepository(db),
		AuthorRepository:   repositories.NewAuthorRepository(db),
	}
}
