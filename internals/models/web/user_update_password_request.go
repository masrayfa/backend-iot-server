package web

type UserUpdatePasswordRequest struct {
	IdUser      int64  `json:"id_user"`
	NewPassword string `json:"new_password"`
	OldPassword string `json:"old_password"`
}
