package main

import (
	"countreez/configs"
)

func main() {
	configs.GetCtx()
	configs.GetPsqlConnection()
	defer configs.DB.Close()
}