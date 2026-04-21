package interfaces

type AuthorRepository interface {
	GetAuthors()
	GetAuthorByID()
	CreateAuthor()
	UpdateAuthor()
	DeleteAuthor()
}
