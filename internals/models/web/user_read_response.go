package web

type UserRead struct {
	IdUser   int64  `json:"id_user"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Status   bool   `json:"status"`
	Token    string `json:"token"`
	IsAdmin  bool   `json:"is_admin"`
}
