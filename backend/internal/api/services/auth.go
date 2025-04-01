package services

import (
	"errors"

	"github.com/Missing-Link-harkat/uuid-validation-server/internal/api/dto"
	"github.com/Missing-Link-harkat/uuid-validation-server/internal/api/repositories"
	"github.com/Missing-Link-harkat/uuid-validation-server/internal/utils"
)

func LoginUser(email, password string) (*dto.UserDTO, error) {
	user, err := repositories.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}
	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("invalid credentials")
	}
	userDTO := dto.ToUserDTO(user)
	return &userDTO, nil
}
