package helpers

// CreateUser incoming request body for create user
type CreateUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
