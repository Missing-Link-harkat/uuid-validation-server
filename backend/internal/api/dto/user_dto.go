package dto

import "github.com/Missing-Link-harkat/uuid-validation-server/internal/api/models"

type UserDTO struct {
	Email string `json:"email"`
}

func ToUserDTO(user *models.User) UserDTO {
	return UserDTO{
		Email: user.Email,
	}
}
