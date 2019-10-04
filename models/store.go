package models

import "database/sql"

type Store struct {
	ID        int64          `json:"id"`
	BrandName string         `json:"brand"`
	SubName   string         `json:"subname"`
	Category  string         `json:"category"`
	LogoUrl   sql.NullString `json:"logo_url"`
	AvgRating float32        `json:"avg_rating"`
	NumRating int            `json:"num_rating"`
	Address   string         `json:"address"`
	Latitude  float32        `json:"latitude"`
	Longitude float32        `json:"longitude"`
	District  string         `json:"district"`
	City      string         `json:"city"`
}
