package user

type Location struct {
	UserId    int64   `json:"user_id"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}
