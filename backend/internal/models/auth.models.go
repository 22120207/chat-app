package models

type SingupRequest struct {
	Fullname        string `json:"fullname" binding:"required"`
	Username        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirmPassword" binding:"required"`
	Gender          string `json:"gender" binding:"required"`
}
