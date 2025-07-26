package main

import (
	"countreez/configs"
	"countreez/database"
)

func main() {
	configs.GetCtx()
	configs.LoadENV()
	configs.GetPsqlConnection()
	database.DBMigrate()
	configs.GetRedisConnection()
	defer configs.DB.Close()
}