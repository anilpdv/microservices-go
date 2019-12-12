package domain

// User : struct
type User struct {
	ID        int64  `json:"id"`
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Email     string `json:"email"`
}
