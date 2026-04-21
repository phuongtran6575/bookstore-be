package repositories

import (
	"github.com/phuongtran6575/bookstore-be.git/internal/repositories/interfaces"
	"gorm.io/gorm"
)

type authorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) interfaces.AuthorRepository {
	return &authorRepository{db: db}
}

func (r *authorRepository) GetAuthors() {

}

func (r *authorRepository) GetAuthorByID() {

}

func (r *authorRepository) CreateAuthor() {

}

func (r *authorRepository) UpdateAuthor() {

}

func (r *authorRepository) DeleteAuthor() {

}
