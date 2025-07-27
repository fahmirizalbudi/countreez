package main

import (
	"countreez/configs"
	"countreez/database"
	"countreez/router"
)

func main() {
	configs.GetCtx()
	configs.LoadENV()
	configs.GetPsqlConnection()
	database.DBMigrate()
	configs.GetRedisConnection()

	router := router.CountreezRoute()
	router.Run(":8080")

	defer configs.DB.Close()
}