package store

type ResponseData struct {
	ID            int64   `json:"id"`
	BrandName     string  `json:"brand"`
	SubName       string  `json:"subname"`
	ImageUrl      string  `json:"image_url"`
	Address       string  `json:"address"`
	DistanceScore float64 `json:"distance"`
	RecommendRank int     `json:"recommend_rank"`
	AvgRating     float32 `json:"avg_rating"`
	NumRating     int64   `json:"num_rating"`
}
