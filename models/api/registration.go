package api

type Registration struct {
	Phone    string `json:"phone"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
