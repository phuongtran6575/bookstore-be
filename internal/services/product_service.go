package services

import (
	repoInterfaces "github.com/phuongtran6575/bookstore-be.git/internal/repositories/interfaces"
	serInterfaces "github.com/phuongtran6575/bookstore-be.git/internal/services/interfaces"
)

type productService struct {
	productRepository repoInterfaces.ProductRepository
}

func NewProductService(productRepository repoInterfaces.ProductRepository) serInterfaces.ProductService {
	return &productService{productRepository: productRepository}
}

func (s *productService) GetProducts() {

}

func (s *productService) GetProductByID() {

}

func (s *productService) CreateProduct() {

}

func (s *productService) UpdateProduct() {

}

func (s *productService) DeleteProduct() {

}
