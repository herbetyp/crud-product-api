package login

type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	GranType string `json:"grant_type"`
}
