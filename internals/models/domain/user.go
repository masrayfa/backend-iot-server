package domain

type User struct {
	IdUser   int64
	Email    string
	Username string
	Password string
	Status   bool
	IsAdmin  bool
}

type UserRead struct {
	IdUser   int64  `json:"id_user"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Status   bool   `json:"status"`
	IsAdmin  bool   `json:"is_admin"`
}