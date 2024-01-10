package web

type UserForgotPasswordRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}