package models

import (
	"database/sql"
	"time"
)

type Review struct {
	ID 			int64				`json:"id"`
	UserID		int64 				`json:"user_id"`
	StoreID 	int64				`json:"store_id"`
	Rating 		float32				`json:"rating"`
	Comment		sql.NullString		`json:"comment"`
	ImageUrls	[]sql.NullString	`json:"image_urls"`
	CreatedAt 	time.Time			`json:"created_at"`
}