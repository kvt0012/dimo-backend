package store_api

type ResponseData struct {
	ID            int64   `json:"id"`
	BrandName     string  `json:"brand"`
	SubName       string  `json:"subname"`
	LogoUrl       string  `json:"logo_url"`
	Category      string  `json:"category"`
	Address       string  `json:"address"`
	Distance      float64 `json:"distance"`
	RecommendRank int     `json:"recommend_rank"`
	AvgRating     float32 `json:"avg_rating"`
	NumRating     int     `json:"num_rating"`
	Latitude      float32 `json:"latitude"`
	Longitude     float32 `json:"longitude"`
}
