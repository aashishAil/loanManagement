package model

type UserLoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginOutput struct {
	Token string `json:"token"`
}
