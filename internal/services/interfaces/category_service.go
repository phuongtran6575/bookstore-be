package interfaces

type CategoryService interface {
	GetCategories()
	GetCategoryByID()
	CreateCategory()
	UpdateCategory()
	DeleteCategory()
}
