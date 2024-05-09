package web

type UserUpdatePasswordRequest struct {
	NewPassword string `json:"new_password"`
	OldPassword string `json:"old_password"`
}
