package handlers

import (
	"net/http"

	"github.com/Missing-Link-harkat/uuid-validation-server/internal/api/models"
	"github.com/Missing-Link-harkat/uuid-validation-server/internal/api/services"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var loginRequest models.LoginRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	userDTO, err := services.LoginUser(loginRequest.Email, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user":    userDTO,
	})
}
