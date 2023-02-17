package repository

import (
	"github.com/ridhozhr10/gdsc-backend-toko/entities"
)

type Product interface {
	Get() ([]entities.Product, error)
	GetOne(id int) (entities.Product, error)
	Create(data entities.Product) (entities.Product, error)
	Update(data entities.Product) (entities.Product, error)
	Delete(id int) error
}
