package interfaces

type AuthorService interface {
	GetAuthors()
	GetAuthorByID()
	CreateAuthor()
	UpdateAuthor()
	DeleteAuthor()
}
