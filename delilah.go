package main

import (
	"github.com/greenac/delilah/logger"
	"github.com/greenac/delilah/src/db"
	"github.com/greenac/delilah/src/helpers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file")
	}

	dv := db.DatabaseVars{}
	dv.Setup(helpers.ConnEnvVars())
	logger.Log(dv.ConnectionString())
}
