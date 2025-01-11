package login

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	GranType string `json:"grant_type"`
	LastLogin string `json:"last_login"`
}
