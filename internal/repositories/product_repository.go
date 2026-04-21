package repositories

import (
	"github.com/phuongtran6575/bookstore-be.git/internal/repositories/interfaces"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) interfaces.ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) GetProducts() {

}

func (r *productRepository) GetProductByID() {

}

func (r *productRepository) CreateProduct() {

}

func (r *productRepository) UpdateProduct() {

}

func (r *productRepository) DeleteProduct() {

}
