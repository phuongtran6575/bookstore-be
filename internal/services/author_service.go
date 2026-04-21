package services

import (
	"github.com/phuongtran6575/bookstore-be.git/internal/models"
	repoInterfaces "github.com/phuongtran6575/bookstore-be.git/internal/repositories/interfaces"
	serInterfaces "github.com/phuongtran6575/bookstore-be.git/internal/services/interfaces"
)

type authorService struct {
	authorRepository repoInterfaces.AuthorRepository
}

func NewAuthorService(authorRepository repoInterfaces.AuthorRepository) serInterfaces.AuthorService {
	return &authorService{authorRepository: authorRepository}
}

func (s *authorService) GetAuthors() []models.Author {
	return s.authorRepository.GetAuthors()
}

func (s *authorService) GetAuthorByID(id uint) models.Author {
	return s.authorRepository.GetAuthorByID(id)
}

func (s *authorService) CreateAuthor(author models.Author) {
	s.authorRepository.CreateAuthor(author)
}

func (s *authorService) UpdateAuthor(author models.Author) {
	s.authorRepository.UpdateAuthor(author)
}

func (s *authorService) DeleteAuthor(id uint) {
	s.authorRepository.DeleteAuthor(id)
}
