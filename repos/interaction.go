package repos

import "dimo-backend/models"

type InteractionRepo interface {
	Insert(interaction *models.Interaction) error
	GetByUserID(userId int64) ([]*models.Interaction, error)
}
