package product

import (
	"log"
	"os"

	"github.com/ridhozhr10/gdsc-backend-toko/entities"
	"github.com/ridhozhr10/gdsc-backend-toko/repository"
)

type productService struct {
	productRepo repository.Product
	logger      log.Logger
}

func NewProductService(productRepo repository.Product) Service {
	logger := log.New(os.Stdout, "ProductService", log.Ldate|log.Ltime|log.Lshortfile)
	return &productService{productRepo: productRepo, logger: *logger}
}

func (s *productService) GetAllProduct() ([]entities.Product, error) {
	data, err := s.productRepo.Get()
	if err != nil {
		log.Print(err)
	}
	return data, err
}

func (s *productService) GetOneProduct(id int) (entities.Product, error) {
	data, err := s.productRepo.GetOne(id)
	if err != nil {
		log.Print(err)
	}
	return data, err
}

func (s *productService) CreateProduct(data entities.Product) (entities.Product, error) {
	data, err := s.productRepo.Create(data)
	if err != nil {
		log.Print(err)
	}
	return data, err
}

func (s *productService) UpdateProduct(data entities.Product) (entities.Product, error) {
	data, err := s.productRepo.Update(data)
	if err != nil {
		log.Print(err)
	}
	return data, err
}

func (s *productService) DeleteProduct(id int) error {
	err := s.productRepo.Delete(id)
	if err != nil {
		log.Print(err)
	}
	return err
}
