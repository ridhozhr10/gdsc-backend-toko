package psql

import (
	"github.com/ridhozhr10/gdsc-backend-toko/entities"
	"github.com/ridhozhr10/gdsc-backend-toko/repository"
	"gorm.io/gorm"
)

type productPsqlRepo struct {
	db *gorm.DB
}

func NewProductPSQLRepo(db *gorm.DB) repository.Product {
	return &productPsqlRepo{db}
}

func (r *productPsqlRepo) Get() ([]entities.Product, error) {
	result := []entities.Product{}
	query := r.db.Find(&result)
	if query.Error != nil {
		return []entities.Product{}, query.Error
	}
	return result, nil
}

func (r *productPsqlRepo) GetOne(id int) (entities.Product, error) {
	result := entities.Product{}
	query := r.db.First(&result)
	if query.Error != nil {
		return entities.Product{}, query.Error
	}
	return result, nil
}

func (r *productPsqlRepo) Create(data entities.Product) (entities.Product, error) {
	q := r.db.Create(&data)
	if q.Error != nil {
		return entities.Product{}, q.Error
	}
	return data, nil
}

func (r *productPsqlRepo) Update(data entities.Product) (entities.Product, error) {
	q := r.db.Save(&data)
	if q.Error != nil {
		return entities.Product{}, q.Error
	}
	return data, nil
}

func (r *productPsqlRepo) Delete(id int) error {
	data, err := r.GetOne(id)
	if err != nil {
		return err
	}
	q := r.db.Delete(&data)
	return q.Error
}
