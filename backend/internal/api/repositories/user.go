package repositories

import (
	"errors"

	"github.com/Missing-Link-harkat/uuid-validation-server/internal/api/models"
	"github.com/Missing-Link-harkat/uuid-validation-server/internal/db"
	"gorm.io/gorm"
)

func FindUserByEmail(email string) (*models.User, error) {
	var user models.User

	if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}
