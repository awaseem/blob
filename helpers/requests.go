package helpers

// User incoming request body for create user
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Token incoming request body for checking tokens
type Token struct {
	Token string `json:"token"`
}
