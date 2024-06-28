package models

type User struct {
	ID       string `json:"id"`
	Password string `json:"password"`
	Login    string `json:"login"`
}
type UserLogin struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
type UserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}
type UserResponseError struct {
	Error string `json:"error"`
}
