package repos

import "dimo-backend/models"

type UserRepo interface {
	GetAll() ([]*models.User, error)
	GetByID(id int64) (*models.User, error)
	GetByPhone(phone string) (*models.User, error)
	Insert(*models.User) error
}