package main

import (
	"log"

	"github.com/Missing-Link-harkat/uuid-validation-server/internal/api"
	"github.com/Missing-Link-harkat/uuid-validation-server/internal/db"
	"github.com/Missing-Link-harkat/uuid-validation-server/internal/utils"
)

func main() {
	utils.LoadEnvVars()

	dbConn, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	r := api.SetupRouter()
	r.RUn(":8080")
}
