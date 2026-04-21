package services

import (
	repoInterfaces "github.com/phuongtran6575/bookstore-be.git/internal/repositories/interfaces"
	serInterfaces "github.com/phuongtran6575/bookstore-be.git/internal/services/interfaces"
)

type categoryService struct {
	categoryRepository repoInterfaces.CategoryRepository
}

func NewCategoryService(categoryRepository repoInterfaces.CategoryRepository) serInterfaces.CategoryService {
	return &categoryService{categoryRepository: categoryRepository}
}

func (s *categoryService) GetCategories() {

}

func (s *categoryService) GetCategoryByID() {

}

func (s *categoryService) CreateCategory() {

}

func (s *categoryService) UpdateCategory() {

}

func (s *categoryService) DeleteCategory() {

}
