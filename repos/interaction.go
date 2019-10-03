package repos

import "dimo-backend/models"

type InteractionRepo interface {
	Insert(interaction *models.Interaction) error
}