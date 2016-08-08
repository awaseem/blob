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

// SetReq request of Set http handler
type SetReq struct {
	Token string `json:"token"`
	Key   string `json:"key"`
	Value string `json:"value"`
}
