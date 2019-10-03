package repoimpl

import (
	"database/sql"
	"dimo-backend/models"
	"dimo-backend/repos"
	"fmt"
	"time"
)

type InteractionRepoImpl struct {
	db *sql.DB
}

func (i *InteractionRepoImpl) Insert(interaction *models.Interaction) error {
	insertStatement := `
	INSERT INTO interactions (user_id, brand_id, type, created_at)
	VALUES ($1, $2, $3, $4)`
	interaction.CreatedAt = time.Now()
	res, err := i.db.Exec(insertStatement, interaction.UserID,
		interaction.BrandID, interaction.Type, interaction.CreatedAt)
	if err != nil {
		return err
	}
	interaction.ID, _ = res.LastInsertId()
	fmt.Println("Record added: ", interaction)
	return nil
}

func NewInteractionRepo(db *sql.DB) repos.InteractionRepo {
	return &InteractionRepoImpl{db:db}
}
