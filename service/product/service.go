package product

import "github.com/ridhozhr10/gdsc-backend-toko/entities"

type Service interface {
	GetAllProduct() ([]entities.Product, error)
	GetOneProduct(id int) (entities.Product, error)
	CreateProduct(data entities.Product) (entities.Product, error)
	UpdateProduct(data entities.Product) (entities.Product, error)
	DeleteProduct(id int) error
}
