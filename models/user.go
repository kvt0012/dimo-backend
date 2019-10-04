package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int64          `json:"id"`
	Name      string         `json:"name"`
	Phone     string         `json:"phone"`
	Password  string         `json:"password"`
	ImageUrl  sql.NullString `json:"image_url"`
	City      sql.NullString `json:"city"`
	CreatedAt time.Time      `json:"created_at"`
}
