package models

type SingupRequest struct {
	Fullname        string `json:"fullname" binding:"required"`
	Username        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required,min=6"`
	ConfirmPassword string `json:"confirmPassword" binding:"required,min=6"`
	Gender          string `json:"gender" binding:"required,oneof=male female"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}
