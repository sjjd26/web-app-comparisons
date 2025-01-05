package models

type User struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Salt      string `json:"salt"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (user User) GetPasswordHash() string {
	// TODO: Implement this function
	return ""
}
