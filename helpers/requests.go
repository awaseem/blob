package helpers

// User incoming request body for create user
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
