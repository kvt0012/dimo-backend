package repos

import "dimo-backend/models"

type ReviewRepo interface {
	Insert(review *models.Review) error
	GetByID(id int64) (*models.Review, error)
	GetAll() ([]*models.Review, error)
	GetByUserID(userId int64) ([]*models.Review, error)
	GetByStoreID(storeId int64) ([]*models.Review, error)
	DeleteByID(id int64) error
}