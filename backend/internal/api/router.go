package api

import (
	"github.com/Missing-Link-harkat/uuid-validation-server/internal/api/handlers"
	"github.com/Missing-Link-harkat/uuid-validation-server/internal/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	gin.SetMode(utils.GetEnv("GIN_MODE"))
	r := gin.Default()
	r.SetTrustedProxies([]string{"localhost"})

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{utils.GetEnv("ALLOWED_ORIGIN")},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	r.GET("/login", handlers.Login)
}
