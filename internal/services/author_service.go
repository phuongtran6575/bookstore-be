package services

import (
	repoInterfaces "github.com/phuongtran6575/bookstore-be.git/internal/repositories/interfaces"
	serInterfaces "github.com/phuongtran6575/bookstore-be.git/internal/services/interfaces"
)

type authorService struct {
	authorRepository repoInterfaces.AuthorRepository
}

func NewAuthorService(authorRepository repoInterfaces.AuthorRepository) serInterfaces.AuthorService {
	return &authorService{authorRepository: authorRepository}
}

func (s *authorService) GetAuthors() {

}

func (s *authorService) GetAuthorByID() {

}

func (s *authorService) CreateAuthor() {

}

func (s *authorService) UpdateAuthor() {

}

func (s *authorService) DeleteAuthor() {

}
