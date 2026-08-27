package domain


type RegisterUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleID   int    `json:"role_id"`
}