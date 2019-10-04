package repos

import "dimo-backend/models"

type StoreRepo interface {
	GetAll() ([]*models.Store, error)
	GetByID(id int64) (*models.Store, error)
	GetByBrandName(brandName string) ([]*models.Store, error)
	GetByCategory(category string) ([]*models.Store, error)
	GetByCity(city string) ([]*models.Store, error)
	GetByDistrict(district string) ([]*models.Store, error)

	UpdateByID(store *models.Store) error
	CountByBrand(brandId int64) (int, error)
}
