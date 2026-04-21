package repositories

import (
	"github.com/phuongtran6575/bookstore-be.git/internal/models"
	"github.com/phuongtran6575/bookstore-be.git/internal/repositories/interfaces"
	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) interfaces.CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) GetCategories() []models.Category {
	var categories []models.Category
	r.db.Find(&categories)
	return categories
}

func (r *categoryRepository) GetCategoryByID(id uint) models.Category {
	var category models.Category
	r.db.First(&category, id)
	return category
}

func (r *categoryRepository) CreateCategory(category models.Category) {
	r.db.Create(&category)
}

func (r *categoryRepository) UpdateCategory(category models.Category) {
	r.db.Save(&category)
}

func (r *categoryRepository) DeleteCategory(id uint) {
	r.db.Delete(&models.Category{}, id)
}
