package interfaces

import "github.com/phuongtran6575/bookstore-be.git/internal/models"

type AuthorService interface {
	GetAuthors() []models.Author
	GetAuthorByID(id uint) models.Author
	CreateAuthor(author models.Author)
	UpdateAuthor(author models.Author)
	DeleteAuthor(id uint)
}
