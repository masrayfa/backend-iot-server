package web

type UserUpdateRequest struct {
	IdUser   int64  `json:"id_user"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Status   bool   `json:"status"`
	IsAdmin  bool   `json:"is_admin"`
}