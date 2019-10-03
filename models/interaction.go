package models

import "time"

type InteractionType string
const (
	View 		InteractionType = "view"
	Route 		InteractionType = "route"
	Transaction InteractionType = "transaction"
)
type Interaction struct {
	ID 			int64				`json:"id"`
	UserID		int64 				`json:"user_id"`
	BrandID		int64				`json:"brand_id"`
	Type		InteractionType		`json:"type"`
	CreatedAt 	time.Time			`json:"created_at"`
}