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

	pr := db.Props{}
	pr.Setup(helpers.ConnEnvVars())

	cn := db.Connection{Props: &pr}
	err = cn.Connect(); if err != nil {
		logger.Error("connection...")
		panic(err)
	}

	rows, err := cn.FindById(1, "Users"); if err != nil {
		logger.Error("pinche:", err)
		panic(err)
	}

	logger.Log("got rows:", rows)
}
