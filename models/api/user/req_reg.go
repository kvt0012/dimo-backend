package user

type RegisterRequest struct {
	Phone    string `json:"phone"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
