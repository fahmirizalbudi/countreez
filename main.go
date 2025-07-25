package main

import (
	"countreez/configs"
)

func main() {
	configs.GetCtx()
	configs.GetPsqlConnection()
	configs.GetRedisConnection()
	defer configs.DB.Close()
}