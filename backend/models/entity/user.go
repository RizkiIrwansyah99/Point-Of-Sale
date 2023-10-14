package entity

type User struct {
	ID        string  `json:"id"`
	FirstName string  `json:"firstname"`
	LastName  *string `json:"lastname"`
	UserName  string  `json:"username"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	Role      *string `json:"role"`
	Image     *string `json:"image"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}
