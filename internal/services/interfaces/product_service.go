package interfaces

type ProductService interface {
	GetProducts()
	GetProductByID()
	CreateProduct()
	UpdateProduct()
	DeleteProduct()
}
