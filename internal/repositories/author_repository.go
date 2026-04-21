package repositories

import (
	"github.com/phuongtran6575/bookstore-be.git/internal/models"
	"github.com/phuongtran6575/bookstore-be.git/internal/repositories/interfaces"
	"gorm.io/gorm"
)

type authorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) interfaces.AuthorRepository {
	return &authorRepository{db: db}
}

func (r *authorRepository) GetAuthors() []models.Author {
	var authors []models.Author
	r.db.Find(&authors)
	return authors
}

func (r *authorRepository) GetAuthorByID(id uint) models.Author {
	var author models.Author
	r.db.First(&author, id)
	return author
}

func (r *authorRepository) CreateAuthor(author models.Author) {
	r.db.Create(&author)
}

func (r *authorRepository) UpdateAuthor(author models.Author) {
	r.db.Save(&author)
}

func (r *authorRepository) DeleteAuthor(id uint) {
	r.db.Delete(&models.Author{}, id)
}
