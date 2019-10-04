package review

type CreateRequest struct {
	UserID  int64   `json:"user_id"`
	StoreID int64   `json:"store_id"`
	Rating  float32 `json:"rating"`
	Comment string  `json:"comment"`
}
type DeleteRequest struct {
	UserID  int64 `json:"user_id"`
	StoreID int64 `json:"store_id"`
}
