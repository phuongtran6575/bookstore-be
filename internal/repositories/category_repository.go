package repositories

import (
	"github.com/phuongtran6575/bookstore-be.git/internal/repositories/interfaces"
	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) interfaces.CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) GetCategories() {

}

func (r *categoryRepository) GetCategoryByID() {

}

func (r *categoryRepository) CreateCategory() {

}

func (r *categoryRepository) UpdateCategory() {

}

func (r *categoryRepository) DeleteCategory() {

}
