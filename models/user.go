package models

type UserModel struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	UToken    string `json:"utoken"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
