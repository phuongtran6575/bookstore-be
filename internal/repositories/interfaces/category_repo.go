package interfaces

type CategoryRepository interface {
	GetCategories()
	GetCategoryByID()
	CreateCategory()
	UpdateCategory()
	DeleteCategory()
}
