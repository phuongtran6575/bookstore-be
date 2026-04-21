package services

import (
	"github.com/phuongtran6575/bookstore-be.git/internal/models"
	repoInterfaces "github.com/phuongtran6575/bookstore-be.git/internal/repositories/interfaces"
	serInterfaces "github.com/phuongtran6575/bookstore-be.git/internal/services/interfaces"
)

type categoryService struct {
	categoryRepository repoInterfaces.CategoryRepository
}

func NewCategoryService(categoryRepository repoInterfaces.CategoryRepository) serInterfaces.CategoryService {
	return &categoryService{categoryRepository: categoryRepository}
}

func (s *categoryService) GetCategories() []models.Category {
	return s.categoryRepository.GetCategories()
}

func (s *categoryService) GetCategoryByID(id uint) models.Category {
	return s.categoryRepository.GetCategoryByID(id)
}

func (s *categoryService) CreateCategory(category models.Category) {
	s.categoryRepository.CreateCategory(category)
}

func (s *categoryService) UpdateCategory(category models.Category) {
	s.categoryRepository.UpdateCategory(category)
}

func (s *categoryService) DeleteCategory(id uint) {
	s.categoryRepository.DeleteCategory(id)
}
