package interfaces

type ProductRepository interface {
	GetProducts()
	GetProductByID()
	CreateProduct()
	UpdateProduct()
	DeleteProduct()
}
