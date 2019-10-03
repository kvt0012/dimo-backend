package repos

import "dimo-backend/models"

type BrandRepo interface {
	GetAll() 						([]*models.Brand, error)
	GetByCategory(category string) 	([]*models.Brand, error)
	GetByTag(tag string) 			([]*models.Brand, error)
	GetByID(id int64) 				(*models.Brand, error)
	GetByName(name string)			(*models.Brand, error)
}