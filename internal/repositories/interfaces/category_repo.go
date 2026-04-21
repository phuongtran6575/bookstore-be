package interfaces

import "github.com/phuongtran6575/bookstore-be.git/internal/models"

type CategoryRepository interface {
	GetCategories() []models.Category
	GetCategoryByID(id uint) models.Category
	CreateCategory(category models.Category)
	UpdateCategory(category models.Category)
	DeleteCategory(id uint)
}
